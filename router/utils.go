package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 핸들러를 등록합니다.

func (r *Router) GET(path string, handlers gin.HandlerFunc) gin.IRoutes {
	return r.Engin.GET(path, handlers)
}

func (r *Router) POST(path string, handlers gin.HandlerFunc) gin.IRoutes {
	return r.Engin.POST(path, handlers)
}

func (r *Router) PUT(path string, handlers gin.HandlerFunc) gin.IRoutes {
	return r.Engin.PUT(path, handlers)
}

func (r *Router) DELETE(path string, handlers gin.HandlerFunc) gin.IRoutes {
	return r.Engin.DELETE(path, handlers)
}

// -> Response Util

func (r *Router) ResponseOK(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}

func (r *Router) ResponseErr(c *gin.Context, err ...interface{}) {
	c.JSON(http.StatusInternalServerError, err)
}
