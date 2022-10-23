package router

import (
	"github.com/gin-gonic/gin"
	"zcw-gin/app/controller"
	"zcw-gin/middleware"
)

// Api api
func Api(r *gin.Engine) {

	// 公开接口
	api := r.Group("/api")
	api.POST("/login", controller.Index.Login)
	api.GET("/test", controller.Index.Test)

	// 权限接口
	api.Use(middleware.JWT)
	api.POST("/test", controller.Index.Test)
}
