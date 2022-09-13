package router

import (
	"github.com/gin-gonic/gin"
	"zcw-admin-server/controller"
	"zcw-admin-server/middleware"
)

func Load(r *gin.Engine) {
	// 请求预处理
	r.Use(
		middleware.Cors,  //跨域处理
		middleware.Error, //异常处理
	)
	// 回调
	r.GET("/common/notify", controller.Common.Notify)
	// api
	api := r.Group("/api")
	api.GET("/test", controller.Index.Test)
}
