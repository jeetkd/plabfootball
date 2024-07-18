package mongo

import (
	"fmt"
	"net/http"
	"plabfootball/types"
	"testing"
	"time"
)

func Test_getGirlStadium(t *testing.T) {
	today := time.Now() //현재시간.
	currentDate := today.Format("2006-01-02")
	testUrl := fmt.Sprintf("https://www.plabfootball.com/api/v2/integrated-matches/?page_size=700&ordering=schedule&sch=%s&sex=0&hide_soldout=&region=2", currentDate) //실시간 url 요청.
	testData := types.GirlUrlRes{Url: []string{"https://www.plabfootball.com/api/v2/matches/534/"}}

	tests := []struct {
		name     string // 테스트 이름
		url      string // 테스트 할 값
		err      error  // 예상되는 에러
		expected *types.GirlUrlRes
	}{
		{name: "unmarshalError", url: "https://www.plabfootball.com/api/v2/matches/534/", err: fmt.Errorf("JSON 데이터 디코딩 오류: json: cannot unmarshal object into Go value of type []types.StadiumReq")},
		{name: "StatusNotFound", url: "https://www.plabfootball.com/api/v2/matches/40362732s/", err: fmt.Errorf("HTTP 상태 코드 오류: %d", http.StatusNotFound)},
		{name: "GETError", url: "https://www.plabfootballsad.com/api/v2/matches/534/", err: fmt.Errorf("GET 요청 에러")},
		{name: "NoExistStadium_Girl", url: testUrl, err: nil, expected: nil},     //url이 존재하지 않으면 성공.
		{name: "ExistStadium_Girl", url: testUrl, err: nil, expected: &testData}, //url이 존재하면 성공.
	}

	for _, e := range tests {
		expected, err := getGirlStadium(e.url)
		if e.err != nil && err == nil {
			t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
		} else if e.err == nil && err != nil {
			t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
		} else {
			// 여자가 존재하는지 여부 확인.
			if e.expected == nil && expected != nil {
				t.Errorf("%s: expected %v, but got %v", e.name, e.expected, expected)
			}
			if e.expected != nil && expected == nil {
				t.Errorf("%s: expected %v, but got %v", e.name, e.expected, expected)
			}
		}
	}
}

func Test_checkSex(t *testing.T) {
	tests := []struct {
		name     string // 테스트 이름
		url      string // 테스트 할 값
		expected bool   // 예상되는 bool 값
		err      error  // 예상되는 에러
	}{
		{name: "StatusOK", url: "https://www.plabfootball.com/api/v2/matches/534/", expected: false, err: nil},
		{name: "StatusNotFound", url: "https://www.plabfootball.com/api/v2/matches/40362732s/", expected: false, err: fmt.Errorf("HTTP 상태 코드 오류: %d", http.StatusNotFound)},
		{name: "GETError", url: "https://www.plabfootballsad.com/api/v2/matches/534/", expected: false, err: fmt.Errorf("GET 요청 에러")},
		{name: "unmarshalError", url: "https://www.plabfootball.com/api/v2/integrated-matches/?page_size=700&ordering=schedule&sch=2024-07-17&region=me", expected: false, err: fmt.Errorf("JSON 데이터 디코딩 오류: json: cannot unmarshal array into Go value of type types.UsersReq")},
		{name: "ExistGirl", url: "https://www.plabfootball.com/api/v2/matches/395320/", expected: true, err: nil},
		{name: "NoExistGirl", url: "https://www.plabfootball.com/api/v2/matches/399007/", expected: false, err: nil},
	}

	for _, e := range tests {
		result, err := checkSex(e.url)
		if e.err != nil && err == nil {
			t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
		} else if e.err == nil && err != nil {
			t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
		} else {
			// 여자가 있는지 체크
			if e.expected && !result {
				t.Errorf("%s: expected true but got false", e.name)
			}

			if !e.expected && result {
				t.Errorf("%s: expected false but got true", e.name)
			}
		}
	}

}

func Test_getStadiums(t *testing.T) {

	today := time.Now() //현재시간.
	currentDate := today.Format("2006-01-02")
	testUrl := fmt.Sprintf("https://www.plabfootball.com/api/v2/integrated-matches/?page_size=700&ordering=schedule&sch=%s&sex=0&hide_soldout=&region=1", currentDate) //실시간 url 요청.

	tests := []struct {
		name string // 테스트 이름
		url  string // 테스트 할 값
		err  error  // 예상되는 에러
	}{
		{name: "unmarshalError", url: "https://www.plabfootball.com/api/v2/matches/534/", err: fmt.Errorf("JSON 데이터 디코딩 오류: json: cannot unmarshal object into Go value of type []types.StadiumReq")},
		{name: "StatusNotFound", url: "https://www.plabfootball.com/api/v2/matches/40362732s/", err: fmt.Errorf("HTTP 상태 코드 오류: %d", http.StatusNotFound)},
		{name: "StatusOK", url: testUrl, err: nil},
		{name: "GETError", url: "https://www.plabfootballsad.com/api/v2/matches/534/", err: fmt.Errorf("GET 요청 에러")},
	}
	for _, e := range tests {
		_, err := getStadiums(e.url)
		if e.err != nil && err == nil {
			t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
		} else if e.err == nil && err != nil {
			t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
		}
	}
}

func Test_getBody(t *testing.T) {
	tests := []struct {
		name string //테스트 이름
		url  string //테스트 할 값
		err  error  //예상되는 에러
	}{
		{name: "StatusOK", url: "https://www.plabfootball.com/api/v2/matches/534/", err: nil},
		{name: "StatusNotFound", url: "https://www.plabfootball.com/api/v2/matches/40362732s/", err: fmt.Errorf("HTTP 상태 코드 오류: %d", http.StatusNotFound)},
		{name: "GETError", url: "https://www.plabfootballsad.com/api/v2/matches/534/", err: fmt.Errorf("GET 요청 에러")},
	}
	for _, e := range tests {
		_, err := getBody(e.url)
		if e.err != nil && err == nil {
			t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
		} else if e.err == nil && err != nil {
			t.Errorf("%s: expected %v, but got %v", e.name, e.err, err)
		}
	}
}
