package mongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"plabfootball/config"
	"plabfootball/types"
	"plabfootball/types/schema"
	"time"
)

type Mongo struct {
	config *config.Config

	client *mongo.Client
	db     *mongo.Database

	// 사용할 컬렉션
	places *mongo.Collection
}

// NewMongo 는 DB와 연결을 하기 위한 초기화 과정을 거친 후 Mongo 객체를 반환합니다.
func NewMongo(config *config.Config) (*Mongo, error) {
	m := &Mongo{
		config: config,
	}

	ctx := context.Background()
	var err error

	// mongo 접속
	if m.client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo.Uri)); err != nil {
		panic(err)
	} else if err = m.client.Ping(ctx, nil); err != nil {
		panic(err)
	} else {
		// db 연결
		m.db = m.client.Database(config.Mongo.Db)
		// 컬렉션 연결
		m.places = m.db.Collection("Places")
	}
	return m, nil
}

// Add 는 service단에서 받은 데이터를 DB에 추가합니다.
func (m *Mongo) Add(sch string, sex, region int) error {
	currentTime := time.Now()
	url := fmt.Sprintf("https://www.plabfootball.com/api/v2/integrated-matches/?page_size=700&ordering=schedule&sch=%s&sex=%d&hide_soldout=&region=%d", sch, sex, region)

	document := bson.M{
		"url":       url,
		"sch":       sch,
		"sex":       sex,
		"region":    region,
		"updatedAt": currentTime,
		"createdAt": currentTime,
		"expireAt":  currentTime.Add(time.Hour * 24 * 7), // 7일후 document 자동 삭제
	}

	_, err := m.places.InsertOne(context.Background(), document)
	return err
}

// View 는 service단에서 받은 데이터를 DB에서 가져옵니다.
func (m *Mongo) View(sch string, region, sex int) (*schema.Stadium, error) {
	var filter bson.M
	var s schema.Stadium

	filter = bson.M{"sch": sch, "region": region, "sex": sex}

	if err := m.places.FindOne(context.Background(), filter).Decode(&s); err != nil {
		return nil, err
	} else {
		return &s, nil
	}
}

// ViewAll 는 DB에 특정 collection안의 데이터를 모두 가져옵니다.
func (m *Mongo) ViewAll() ([]*schema.Stadium, error) {
	filter := bson.M{}

	ctx := context.Background()

	if cursor, err := m.places.Find(ctx, filter); err != nil {
		return nil, err
	} else {
		defer cursor.Close(ctx)

		var s []*schema.Stadium
		// 모든 document들을 순회하면서 s에 디코딩함.
		if err := cursor.All(ctx, &s); err != nil {
			return nil, err
		} else {
			return s, nil
		}
	}
}

// Upsert 는 service단에서 받은 데이터를 DB에 추가 또는 업데이트합니다.
func (m *Mongo) Upsert(sch string, sex, region int, upsert types.AddReq) (*schema.Stadium, error) {
	var s schema.Stadium

	// document가 없으면 새로 생성 가능 옵션.
	opts := options.FindOneAndUpdate().SetUpsert(true)
	//업데이트할 필터 조건
	filter := bson.M{
		"sch":    sch,
		"sex":    sex,
		"region": region,
	}
	// 현재시간
	currentTime := time.Now()
	url := fmt.Sprintf("https://www.plabfootball.com/api/v2/integrated-matches/?page_size=700&ordering=schedule&sch=%s&sex=%d&hide_soldout=&region=%d", upsert.Sch, upsert.Sex, upsert.Region)

	// 업데이트할 내용
	update := bson.M{
		"$set": bson.M{
			"url":       url,
			"sch":       upsert.Sch,
			"sex":       upsert.Sex,
			"region":    upsert.Region,
			"updatedAt": currentTime,
			"expireAt":  currentTime.Add(time.Hour * 24 * 7), // 7일후 document 자동 삭제
		},
	}

	err := m.places.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&s)
	if err != nil {
		return nil, err
	} else {
		return &s, nil
	}
}

// Delete 는 service단에서 받은 데이터를 DB에서 삭제합니다.
func (m *Mongo) Delete(sch string, region, sex int) error {
	url := fmt.Sprintf("https://www.plabfootball.com/api/v2/integrated-matches/?page_size=700&ordering=schedule&sch=%s&sex=%d&hide_soldout=&region=%d", sch, sex, region)
	//삭제할 필터 조건
	filter := bson.M{
		"url": url,
	}

	result, err := m.places.DeleteOne(context.Background(), filter)
	if result.DeletedCount == 0 {
		return errors.New("delete할 document가 없습니다.")
	}
	return err
}
