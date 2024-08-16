package router

import (
	"os"
	"plabfootball/config"
	"plabfootball/repository"
	"plabfootball/service"
	"testing"
)

var baseUri = "/mongo"
var abPath = "/Users/lovet/GolandProjects/plabfootball/cmd/config.toml"

var router *Router
var s *service.Service
var r *repository.Repository

func TestMain(m *testing.M) {
	var err error
	config := config.NewConfig(abPath)

	if r, err = repository.NewRepository(config); err != nil {
		panic(err)
	}

	if s, err = service.NewService(config, r); err != nil {
		panic(err)
	}
	router, err = NewRouter(config, s, r)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
