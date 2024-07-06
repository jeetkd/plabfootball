package repository

import (
	"plabfootball/config"
	"plabfootball/repository/mongo"
)

type Repository struct {
	config *config.Config

	Mongo *mongo.Mongo
}

// NewRepository 는 새로운 DB를 연결하고 Repository 객체를 반환합니다.
func NewRepository(config *config.Config) (*Repository, error) {
	r := &Repository{
		config: config,
	}

	var err error
	if r.Mongo, err = mongo.NewMongo(config); err != nil {
		panic(err)
	} else {
		return r, nil
	}
}
