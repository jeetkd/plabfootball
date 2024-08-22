package router

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"plabfootball/types"
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

// view 핸들러 테스트
func Test_MongoRouter_view(t *testing.T) {

	var tests = []struct {
		name               string
		postedData         types.ViewReq
		expectedStatusCode int
	}{
		{ //ResponseOK : document가 있는 경우를 예상하는 테스트 케이스1.
			name: "success",
			postedData: types.ViewReq{
				Region: 2,
				Sch:    "2024-08-17",
			},
			expectedStatusCode: http.StatusOK,
		},
		{ // Failed To Call view Data err : document가 없는 경우를 예상하는 테스트 케이스2.
			name: "no document in result",
			postedData: types.ViewReq{
				Region: 1,
				Sch:    "9999-99-99",
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{ //bind 실패 : validation required tag 에러가 나는 경우를 예상하는 테스트 케이스3"
			name: "Field validation for '' on the 'required' tag",
			postedData: types.ViewReq{
				Sch: "9999-99-99",
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, e := range tests {
		jsonValue, _ := json.Marshal(e.postedData)
		req, _ := http.NewRequest("POST", baseUri+"/view", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.Engin.ServeHTTP(w, req)

		if w.Code != e.expectedStatusCode {
			t.Errorf("%s: returned wrong status code; expected %d, but got %d", e.name, e.expectedStatusCode, w.Code)
		}
	}
}

func Test_MongoRouter_viewAll(t *testing.T) {
	var tests = []struct {
		name string
		body string
	}{
		{ // ResponseOK : document가 있는 경우를 예상하는 테스트 케이스1.
			name: "success",
			body: "url",
		},
		{ // Failed To Call view Data err : document가 없는 경우를 예상하는 테스트 케이스2.
			name: "no document in result",
			body: "null",
		},
	}

	for _, e := range tests {
		req, _ := http.NewRequest("POST", baseUri+"/viewAll", nil)
		//req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.Engin.ServeHTTP(w, req)

		// body에 데이터가 비었는지 비교하고 테스트 테이스 실패 결정.
		if !strings.Contains(w.Body.String(), e.body) {
			t.Errorf("%s: returned wrong body; expected %s, but got %s", e.name, e.body, w.Body.String())
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
