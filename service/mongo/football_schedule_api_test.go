package mongo

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
)

//
//import "testing"
//
//func Test_getGirlStadium(t testing.T) {
//	tests := []struct {
//		name     string
//		testUrl  string //테스트 할 값
//		expected bool   //예상되는 bool값
//		msg      string //예상되는 메시지
//	}{
//		{},
//	}
//
//}
//
//func Test_checkSex(t testing.T) {
//
//}

func Test_getStadiums(t *testing.T) {

	today := time.Now() //현재시간.
	currentDate := today.Format("2006-01-02")
	testUrl := fmt.Sprintf("https://www.plabfootball.com/api/v2/integrated-matches/?page_size=700&ordering=schedule&sch=%s&sex=0&hide_soldout=&region=1", currentDate)

	tests := []struct {
		name string
		url  string //테스트 할 값
		err  error  //예상되는 에러
	}{
		{name: "unmarshalError", url: "https://www.plabfootball.com/api/v2/matches/534/", err: fmt.Errorf("JSON 데이터 디코딩 오류: json: cannot unmarshal object into Go value of type []types.StadiumReq")},
		{name: "StatusNotFound", url: "https://www.plabfootball.com/api/v2/matches/40362732s/", err: fmt.Errorf("HTTP 상태 코드 오류: %d", http.StatusNotFound)},
		{name: "StatusOK", url: testUrl, err: nil},
		{name: "GETError", url: "https://www.plabfootballsad.com/api/v2/matches/534/", err: fmt.Errorf("GET 요청 에러")},
	}
	for _, e := range tests {
		_, err := getStadiums(e.url)
		if e.err == nil && err != nil {
			t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
		} else if err == nil && e.err != nil {
			t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
		} else if err != nil && e.err != nil {
			// err 메시지가 둘다 nil이 아닐때 문자열 비교
			if !strings.EqualFold(e.err.Error(), err.Error()) {
				t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
			}
		}
	}
}

func Test_getBody(t *testing.T) {
	tests := []struct {
		name string
		url  string //테스트 할 값
		err  error  //예상되는 에러
	}{
		{name: "StatusOK", url: "https://www.plabfootball.com/api/v2/matches/534/", err: nil},
		{name: "StatusNotFound", url: "https://www.plabfootball.com/api/v2/matches/40362732s/", err: fmt.Errorf("HTTP 상태 코드 오류: %d", http.StatusNotFound)},
		{name: "GETError", url: "https://www.plabfootballsad.com/api/v2/matches/534/", err: fmt.Errorf("GET 요청 에러")},
	}
	for _, e := range tests {
		_, err := getBody(e.url)
		if e.err == nil && err != nil {
			t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
		} else if err == nil && e.err != nil {
			t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
		} else if err != nil && e.err != nil {
			// err 메시지가 둘다 nil이 아닐때 문자열 비교
			if !strings.EqualFold(e.err.Error(), err.Error()) {
				t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
			}
		}
	}
}
