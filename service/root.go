package service

import (
	"plabfootball/config"
	"plabfootball/repository"
	"plabfootball/service/mongo"
)

type Service struct {
	config     *config.Config
	repository *repository.Repository
	MService   *mongo.MService
}

// NewService 는 repository와 연결해줄 다리인 새로운 서비스를 생성.
func NewService(config *config.Config, repository *repository.Repository) (*Service, error) {
	r := &Service{
		config:     config,
		repository: repository,
		MService:   mongo.NewMService(repository),
	}

	return r, nil
}
