package mongo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"plabfootball/types"
)

// 여자가 포함된 구장들을 가져옵니다.
func getGirlStadium(url string) (*types.GirlUrlRes, error) {
	var girl_URLs types.GirlUrlRes

	if stadiums, err := getStadiums(url); err != nil {
		return nil, err
	} else {
		for _, stadium := range stadiums {
			urlMatch := fmt.Sprintf("https://www.plabfootball.com/api/v2/matches/%d/", stadium.Id)
			urlStadium := fmt.Sprintf("https://www.plabfootball.com/match/%d/", stadium.Id)
			// 경기가 끝난지 체크
			if !stadium.Finish {
				// 경기장에 성별(여) 체크
				if existGirl, err := checkSex(urlMatch); err == nil {
					if existGirl {
						girl_URLs.Url = append(girl_URLs.Url, urlStadium)
					}
				}
			}
		}
	}

	// url이 존재하지 않을 때
	if girl_URLs.Url == nil {
		return nil, nil
	}

	return &girl_URLs, nil
}

// 성별의 체크합니다.
func checkSex(url string) (bool, error) {
	var users types.UsersReq

	data, err := getBody(url)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return false, fmt.Errorf("JSON 데이터 디코딩 오류: %v", err)
	}

	for _, user := range users.Applys {
		// 여자인지 체크(-1 여자, 1남자)
		if user.UserSex == -1 {
			return true, nil
		}
	}
	return false, nil
}

// 경기장들의 정보를 가져옵니다.
func getStadiums(url string) ([]types.StadiumReq, error) {
	var stadiums []types.StadiumReq

	data, err := getBody(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &stadiums)
	if err != nil {
		return nil, fmt.Errorf("JSON 데이터 디코딩 오류: %v", err)
	}
	return stadiums, nil
}

// html body를 가져옵니다.
func getBody(url string) ([]byte, error) {
	// GET 요청 생성
	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.New("GET 요청 에러")
	}
	defer resp.Body.Close()

	// 응답 상태 코드 확인
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP 상태 코드 오류: %d", resp.StatusCode)
	}

	// 응답 본문 읽기
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("응답 본문 읽기 오류: %v", err)
	}
	return data, nil
}
