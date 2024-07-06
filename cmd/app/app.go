package app

import (
	"plabfootball/config"
	"plabfootball/repository"
	"plabfootball/router"
	"plabfootball/service"
)

type App struct {
	config *config.Config
	router *router.Router

	repository *repository.Repository
	service    *service.Service
}

// NewApp 은 Rpository, Service, Router들을 초기화 시킵니다.
func NewApp(config *config.Config) *App {
	a := &App{
		config: config,
	}
	var err error

	if a.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	}

	if a.service, err = service.NewService(config, a.repository); err != nil {
		panic(err)
	}
	if a.router, err = router.NewRouter(config, a.service, a.repository); err != nil {
		panic(err)
	}

	return a
}
