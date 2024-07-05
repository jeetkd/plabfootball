package schema

import "time"

// -> schema(DB내에 어떤 구조로 데이터가 저장되는가를 나타내는 데이터베이스 구조)

type Stadium struct {
	URL         string    `json:"url" bson:"url"`                   // 구장들 정보가 있는 url
	Sex         int64     `json:"sex" bson:"sex"`                   // 성별
	HideSoldout string    `json:"hide_soldout" bson:"hide_soldout"` // 마감된 구장 구별 여부
	Region      int64     `json:"region" bson:"region"`             // 지역
	Sch         string    `json:"sch" bson:"sch"`                   // 날짜
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`       // 생성시간
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`       // 업데이트된 시간
	expireAt    time.Time `json:"expireAt" bson:"expireAt"`         // document를 삭제할 만료 날짜
}
