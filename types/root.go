package types

// -> Request
type AddReq struct {
	Sex    int    `json:"sex" binding:"eq=0|eq=1"`         // 성별
	Region int    `json:"region" binding:"required,gte=1"` // 지역
	Sch    string `json:"sch" binding:"required"`          // 날짜
}

type ViewReq struct {
	Sex    int    `json:"sex" binding:"eq=0|eq=1"`         // 성별
	Region int    `json:"region" binding:"required,gte=1"` // 지역
	Sch    string `json:"sch" binding:"required"`          // 날짜
}

type UpdateReq struct {
	Sex    int    `json:"sex" binding:"eq=0|eq=1"`         // 업데이트 하고 싶은 데이터 조건 : 성별
	Region int    `json:"region" binding:"required,gte=1"` // 업데이트 하고 싶은 데이터 조건 : 지역
	Sch    string `json:"sch" binding:"required"`          // 업데이트 하고 싶은 데이터 조건 : 날짜
	Upsert AddReq `json:"upsert"`                          // 업데이트할 데이터
}

type DeleteReq struct {
	Sex    int    `json:"sex" binding:"eq=0|eq=1"`         // 성별
	Region int    `json:"region" binding:"required,gte=1"` // 지역
	Sch    string `json:"sch" binding:"required"`          // 날짜
}

type PlaceReq struct {
	Sex    int    `json:"sex" binding:"eq=0|eq=1"`         // 성별
	Region int    `json:"region" binding:"required,gte=1"` // 지역
	Sch    string `json:"sch" binding:"required"`          // 날짜
}

type StadiumReq struct {
	Id       int    `json:"id"`
	Schedule string `json:"schedule"`
	Finish   bool   `json:"is_finish"`
}

type User struct {
	Id      int `json:"id"`
	UserSex int `json:"user_sex"`
}

type UsersReq struct {
	Applys []User `json:"applys"`
}

// -> Response
type GirlUrlRes struct {
	Url []string
}
