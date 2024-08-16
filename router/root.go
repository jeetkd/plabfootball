package router

import (
	"github.com/gin-gonic/gin"
	"plabfootball/config"
	"plabfootball/repository"
	"plabfootball/service"
)

type Router struct {
	config *config.Config

	Engin      *gin.Engine
	service    *service.Service
	repository *repository.Repository
}

// NewRouter 는 라우터를 초기화와 등록을 함.
func NewRouter(config *config.Config, service *service.Service, repository *repository.Repository) (*Router, error) {
	r := &Router{
		config:     config,
		Engin:      gin.New(),
		service:    service,
		repository: repository,
	}

	NewMongoRouter(r, r.service.MService) //mongo 라우터

	return r, nil
}
