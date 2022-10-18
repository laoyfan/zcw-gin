package router

import (
	"github.com/gin-gonic/gin"
	"zcw-admin-server/global"
	"zcw-admin-server/middleware"
)

func Server() *gin.Engine {
	// 关闭控制台颜色
	gin.DisableConsoleColor()
	gin.SetMode(global.CONFIG.App.Mode)

	r := gin.New()

	r.Use(
		middleware.Logger, //日志处理
		middleware.Cors,   //跨域处理
		middleware.Error,  //异常处理
	)

	Api(r)

	return r
}
