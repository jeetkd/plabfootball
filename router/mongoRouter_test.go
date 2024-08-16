package router

import (
	"github.com/gin-gonic/gin"
	"strings"
	"testing"
)

func Test_MongoRouter_handler(t *testing.T) {

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

	// route 테스트.
	for _, route := range registered {
		if !routeExists(route.route, route.method, router.Engin) {
			t.Errorf("route %s is not registered", route.route)
		}
	}

}

func Test_MongoRouter_view(t *testing.T) {

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
