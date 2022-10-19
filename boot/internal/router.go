package internal

import (
	"github.com/gin-gonic/gin"
	"zcw-admin-server/router"
)

func Route() *gin.Engine {
	// 开启gin实例
	r := gin.New()
	// 转载路由
	router.Api(r)

	return r
}
