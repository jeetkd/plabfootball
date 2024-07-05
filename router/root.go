package router

import (
	"github.com/gin-gonic/gin"
	"plabfootball/config"
	"plabfootball/repository"
	"plabfootball/service"
)

type Router struct {
	config *config.Config

	engin      *gin.Engine
	service    *service.Service
	repository *repository.Repository
}

func NewRouter(config *config.Config, service *service.Service, repository *repository.Repository) (*Router, error) {
	r := &Router{
		config:     config,
		engin:      gin.New(),
		service:    service,
		repository: repository,
	}

	NewMongoRouter(r, r.service.MService) //mongo 라우터

	return r, r.engin.Run(config.Network.Port) //서버 실행.
}
