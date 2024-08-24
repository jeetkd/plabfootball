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
		name       string
		postedData types.ViewReq
		body       string
	}{
		{ //ResponseOK : document가 있는 경우를 예상하는 테스트 케이스1.
			name: "success",
			postedData: types.ViewReq{
				Region: 2,
				Sch:    "2024-08-23",
			},
			body: "url",
		},
		{ // Failed To Call view Data err : document가 없는 경우를 예상하는 테스트 케이스2.
			name: "no document in result",
			postedData: types.ViewReq{
				Region: 1,
				Sch:    "9999-99-99",
			},
			body: "server 에러 : mongo: no documents in result",
		},
		{ //bind 실패 : validation required tag 에러가 나는 경우를 예상하는 테스트 케이스3"
			name: "Field validation for '' on the 'required' tag",
			postedData: types.ViewReq{
				Sch: "9999-99-99",
			},
			body: "bind 실패",
		},
	}

	for _, e := range tests {
		jsonValue, _ := json.Marshal(e.postedData)
		req, _ := http.NewRequest("POST", baseUri+"/view", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.Engin.ServeHTTP(w, req)

		// body에 데이터가 비었는지 비교하고 테스트 테이스 실패 결정.
		if !strings.Contains(w.Body.String(), e.body) {
			t.Errorf("%s: returned wrong body; expected %s, but got %s", e.name, e.body, w.Body.String())
		}
	}
}

// viewAll 핸들러 테스트
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

// add 핸들러 테스트
func Test_MongoRouter_add(t *testing.T) {
	var tests = []struct {
		name       string
		postedData types.AddReq
		body       string
	}{
		{ //ResponseOK : document가 성공적으로 생생된 경우를 예상하는 테스트 케이스1.
			name: "add successes",
			postedData: types.AddReq{
				Sex:    0,
				Region: 2,
				Sch:    "9999-99-99",
			},
			body: "Success",
		},
		{ // "server 에러 : Document가 이미 존재합니다" : 같은 document가 이미 존재하는 경우를 예상하는 테스트 케이스2.
			name: "document already exists",
			postedData: types.AddReq{
				Sex:    0,
				Region: 2,
				Sch:    "9999-99-99",
			},
			body: "server 에러 : Document가 이미 존재합니다.",
		},
		{ //bind 실패 : validation required tag 에러가 나는 경우를 예상하는 테스트 케이스3"
			name: "Field validation for '' on the 'required' tag",
			postedData: types.AddReq{
				Sch: "9999-99-99",
			},
			body: "bind 실패 :",
		},
	}

	for _, e := range tests {
		jsonValue, _ := json.Marshal(e.postedData)
		req, _ := http.NewRequest("POST", baseUri+"/add", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.Engin.ServeHTTP(w, req)

		// body에 데이터가 비었는지 비교하고 테스트 테이스 실패 결정.
		if !strings.Contains(w.Body.String(), e.body) {
			t.Errorf("%s: returned wrong body; expected %s, but got %s", e.name, e.body, w.Body.String())
		}
	}

}

// upsert 핸들러 테스트
func Test_MongoRouter_upsert(t *testing.T) {
	var tests = []struct {
		name       string
		postedData types.UpdateReq
		body       string
	}{
		{ //server 에러 : 새로운 document가  생생된 경우를 예상하는 테스트 케이스1.
			name: "New document added",
			postedData: types.UpdateReq{
				Sex:    0,
				Region: 2,
				Sch:    "9999-99-99",
				Upsert: types.AddReq{
					Sex:    0,
					Region: 1,
					Sch:    "9999-99-00",
				},
			},
			body: "server 에러 : Created new document",
		},
		{ // server 에러 : Data that you want to insert is already exist : 내가 원하는 값으로 바꾸기 위한 데이터(AddReq)가 이미 존재하는 경우를 예상하는 테스트 케이스2.
			name: "Data(AddReq) that you want to insert is already exist",
			postedData: types.UpdateReq{
				Sex:    0,
				Region: 1,
				Sch:    "9999-99-99",
				Upsert: types.AddReq{
					Sex:    0,
					Region: 1,
					Sch:    "9999-99-00",
				},
			},
			body: "server 에러 : Data that you want to insert is already exist",
		},
		{ //ResponseOK : document가 성공적으로 업데이트된 경우를 예상하는 테스트 케이스3.
			name: "New document added",
			postedData: types.UpdateReq{
				Sex:    0,
				Region: 1,
				Sch:    "9999-99-00",
				Upsert: types.AddReq{
					Sex:    0,
					Region: 2,
					Sch:    "9999-99-11",
				},
			},
			body: `"url"`,
		},
		{ //bind 실패 : validation required tag 에러가 나는 경우를 예상하는 테스트 케이스4"
			name: "Field validation for '' on the 'required' tag",
			postedData: types.UpdateReq{
				Sch: "9999-99-99",
			},
			body: "bind 실패 :",
		},
	}

	for _, e := range tests {
		jsonValue, _ := json.Marshal(e.postedData)
		req, _ := http.NewRequest("PUT", baseUri+"/upsert", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.Engin.ServeHTTP(w, req)

		// body에 데이터가 비었는지 비교하고 테스트 테이스 실패 결정.
		if !strings.Contains(w.Body.String(), e.body) {
			t.Errorf("%s: returned wrong body; expected %s, but got %s", e.name, e.body, w.Body.String())
		}
	}

}

// delete 핸들러 테스트
func Test_MongoRouter_delete(t *testing.T) {
	var tests = []struct {
		name       string
		postedData types.DeleteReq
		body       string
	}{
		{ //ResponseOK : document가 성공적으로 삭제된 경우를 예상하는 테스트 케이스1.
			name: "delete successes",
			postedData: types.DeleteReq{
				Sex:    0,
				Region: 2,
				Sch:    "9999-99-99",
			},
			body: "Success",
		},
		{ // "server 에러 : delete할 document가 없습니다." : 삭제할 document가 존재하지 않는 경우를 예상하는 테스트 케이스2.
			name: "document already exists",
			postedData: types.DeleteReq{
				Sex:    0,
				Region: 2,
				Sch:    "9999-99-99",
			},
			body: " delete할 document가 없습니다.",
		},
		{ //bind 실패 : validation required tag 에러가 나는 경우를 예상하는 테스트 케이스3"
			name: "Field validation for '' on the 'required' tag",
			postedData: types.DeleteReq{
				Sch: "9999-99-99",
			},
			body: "bind 실패 :",
		},
	}

	for _, e := range tests {
		jsonValue, _ := json.Marshal(e.postedData)
		req, _ := http.NewRequest("DELETE", baseUri+"/delete", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.Engin.ServeHTTP(w, req)

		// body에 데이터가 비었는지 비교하고 테스트 테이스 실패 결정.
		if !strings.Contains(w.Body.String(), e.body) {
			t.Errorf("%s: returned wrong body; expected %s, but got %s", e.name, e.body, w.Body.String())
		}
	}

}

// girlUser 핸들러 테스트
func Test_MongoRouter_girlUser(t *testing.T) {
	var tests = []struct {
		name       string
		postedData types.PlaceReq
		body       string
	}{
		{ //ResponseOK : 성공적으로 데이터를 가져오는것을 예상하는 테스트 케이스1.
			name: "success",
			postedData: types.PlaceReq{
				Sex:    0,
				Region: 2,
				Sch:    "2024-08-25",
			},
			body: "Url",
		},

		{ // Failed To Call view Data err : document가 없는 경우를 예상하는 테스트 케이스2.
			name: "no document in result",
			postedData: types.PlaceReq{
				Sex:    0,
				Region: 2,
				Sch:    "9999-99-99",
			},
			body: "server 에러 : mongo: no documents in result",
		},
		{ //bind 실패 : validation required tag 에러가 나는 경우를 예상하는 테스트 케이스3"
			name: "Field validation for '' on the 'required' tag",
			postedData: types.PlaceReq{
				Sch: "9999-99-99",
			},
			body: "bind 실패 :",
		},
		{ //ResponseOK : 성공적으로 데이터를 가져오나 빈값을 예상하는 테스트 케이스4.
			name: "success",
			postedData: types.PlaceReq{
				Sex:    0,
				Region: 2,
				Sch:    "2021-09-01",
			},
			body: "null",
		},
	}

	for _, e := range tests {
		jsonValue, _ := json.Marshal(e.postedData)
		req, _ := http.NewRequest("POST", "/plaber-girl", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
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
