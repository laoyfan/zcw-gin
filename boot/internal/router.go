package internal

import (
	"github.com/gin-gonic/gin"
	"zcw-admin-server/middleware"
	"zcw-admin-server/router"
)

func Route() *gin.Engine {
	// 开启gin实例
	r := gin.New()
	// 中间件
	r.Use(
		middleware.Logger, //日志处理
		middleware.Cors,   //跨域处理
		middleware.Error,  //异常处理
	)
	// 转载路由
	router.Api(r)

	return r
}
