package router

import (
	"github.com/gin-gonic/gin"
	"zcw-admin-server/app/controller"
	"zcw-admin-server/middleware"
)

func Load(r *gin.Engine) {
	// 请求预处理
	r.Use(
		middleware.Cors,  //跨域处理
		middleware.Error, //异常处理
	)
	// api
	api := r.Group("/api")
	api.GET("/test", controller.Index.Test)
}
