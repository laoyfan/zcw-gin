package internal

import (
	"github.com/gin-gonic/gin"
	"zcw-admin-server/router"
)

func Route(r *gin.Engine) *gin.Engine {
	// 装载路由
	router.Api(r)
	return r
}
