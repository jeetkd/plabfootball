package router

import (
	"github.com/gin-gonic/gin"
	"plabfootball/service/mongo"
	. "plabfootball/types"
	. "plabfootball/types/err"
)

type MongoRouter struct {
	router   *Router
	mService *mongo.MService
}

// NewMongoRouter 는 mongo와 관련된 핸들러를 등록합니다.
func NewMongoRouter(router *Router, mService *mongo.MService) {
	m := MongoRouter{
		router:   router,
		mService: mService,
	}

	baseUri := "/mongo"

	//POST 등록
	m.router.POST(baseUri+"/view", m.view)
	m.router.POST(baseUri+"/viewAll", m.viewAll)
	m.router.POST(baseUri+"/add", m.add)

	//PUT 등록
	m.router.PUT(baseUri+"/upsert", m.upsert)

	//DELETE 등록
	m.router.DELETE(baseUri+"/delete", m.delete)

	//Girl 플래버(Plaber)
	m.router.POST("/plaber-girl", m.girlUser)
}

// add 는 클라이언트에서 보낸 데이터를 바인딩 하고, 서비스단 Add를 호출합니다.
func (m *MongoRouter) add(c *gin.Context) {
	var req AddReq

	if err := c.ShouldBindJSON(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if err := m.mService.Add(req.Sch, req.Sex, req.Region); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, "Success")
	}
}

// upsert 는 클라이언트에서 보낸 데이터를 바인딩 하고, 서비스단 Upsert를 호출합니다.
func (m *MongoRouter) upsert(c *gin.Context) {
	var req UpdateReq

	if err := c.ShouldBindJSON(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if response, err := m.mService.Upsert(req.Sch, req.Sex, req.Region, req.Upsert); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, response)
	}
}

// view 는 클라이언트에서 보낸 데이터를 바인딩 하고, 서비스단 View를 호출합니다.
func (m *MongoRouter) view(c *gin.Context) {
	var req ViewReq

	if err := c.ShouldBindJSON(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if response, err := m.mService.View(req.Sch, req.Region, req.Sex); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, response)
	}
}

// viewAll 는 클라이언트에서 보낸 데이터를 바인딩 하고, 서비스단 ViewAll를 호출합니다.
func (m *MongoRouter) viewAll(c *gin.Context) {
	if response, err := m.mService.ViewAll(); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, response)
	}
}

// delete 는 클라이언트에서 보낸 데이터를 바인딩 하고, 서비스단 Delete를 호출합니다.
func (m *MongoRouter) delete(c *gin.Context) {
	var req DeleteReq

	if err := c.ShouldBindJSON(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if err = m.mService.Delete(req.Sch, req.Sex, req.Region); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, "Success")
	}
}

// girlUser 는 클라이언트에서 보낸 데이터를 바인딩 하고, 서비스단 GetGirlUser를 호출합니다.
func (m *MongoRouter) girlUser(c *gin.Context) {
	var req PlaceReq

	if err := c.ShouldBindJSON(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if response, err := m.mService.GetGirlUser(req.Sch, req.Region, req.Sex); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, response)
	}
}
