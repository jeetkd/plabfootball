package router

import (
	"github.com/gin-gonic/gin"
	"plabfootball/config"
	"plabfootball/repository"
	"plabfootball/service"
	"strings"
	"testing"
)

func Test_MongoRouter_handler(t *testing.T) {
	baseUri := "/mongo"
	var abPath = "/Users/lovet/GolandProjects/plabfootball/cmd/config.toml"

	var r *repository.Repository
	var s *service.Service
	var err error

	var registered = []struct {
		route  string // 요청 경로
		method string // 메소드(get, post, put, delete...)
	}{
		{baseUri + "/view", "POST"},
		{baseUri + "/viewAll", "POST"},
		{baseUri + "/add", "POST"},
		{baseUri + "/upsert", "PUT"},
		{baseUri + "/delete", "DELETE"},
		{"/plaber-girl", "POST"},
	}

	// 핸들러 등록을 위해서 필요한 객체들.(초기화)
	config := config.NewConfig(abPath)

	if r, err = repository.NewRepository(config); err != nil {
		panic(err)
	}

	if s, err = service.NewService(config, r); err != nil {
		panic(err)
	}

	router := &Router{
		engin:   gin.New(),
		service: s,
	}

	NewMongoRouter(router, router.service.MService) //mongo 라우터(핸들러 등록)

	// route 테스트.
	for _, route := range registered {
		if !routeExists(route.route, route.method, router.engin) {
			t.Errorf("route %s is not registered", route.route)
		}
	}

}

// 핸들러 경로 존재 확인.
func routeExists(testRoute, testMethod string, engine *gin.Engine) bool {
	for _, route := range engine.Routes() {
		if strings.EqualFold(route.Method, testMethod) && strings.EqualFold(route.Path, testRoute) {
			return true
		}
	}
	return false
}
