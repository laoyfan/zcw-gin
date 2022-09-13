package router

import (
	"github.com/gin-gonic/gin"
	"zcw-admin-server/controller"
	"zcw-admin-server/middleware"
)

func Load(r *gin.Engine) {
	// 请求预处理
	r.Use(middleware.Cors)

	// 回调
	r.GET("/common/notify", controller.Common.Notify)
	r.GET("/test", middleware.Test, controller.Index.Test)
	// api
	api := r.Group("/api").Use(middleware.Test)
	api.GET("/test", controller.Index.Test)
}
