### 🔖프로젝트 개요
***
- golang 크롤링 강의를 듣고 난후 배운 개념들을 활용해보기 위하여 어떤 프로젝트를 만들까 고민하던 중 평소에 자주 이용하던 플랩풋볼 사이트를 기반으로 golang과 mongodb를 사용하여 프로젝트를 만들어 보기로 했다.
- 플랩풋볼 사이트는 개인이 풋살을 즐길 수 있는 플랫폼이다. (https://www.plabfootball.com/)
- 프로젝트 설명 : 지역별 남녀혼성에서 경기장마다 여자분들이 신청한 구장의 사이트를 가져옵니다.

### 🚩 프로젝트 기간
- 2024년 6월 24일 ~ 2024년 7월 06일(약 2주)

### ⚙개발 환경
***
- `go 1.22.4`
- **IDE** : Goland 2023.2.1
- **Database** : MongoDB Atlas 7.0.11, MongoDB Compass
- **OS** : Windows10
- **API 테스트 툴** : POSTMAN

### 🗂외부 패키지 다운
***

    go get github.com/gin-gonic/gin
    go get github.com/naoina/toml
    go get go.mongodb.org/mongo-driver

### 📃참고 사이트
***
- https://www.inflearn.com/course/관리형-api-자동화-만들며-학습하기
- https://www.inflearn.com/course/서버개발-백엔드-mysql-mongodb


### ▶실행 방법
***
1. cmd/config.toml 파일 작성. ([network], [mongo])
2. cmd/main.go 실행.