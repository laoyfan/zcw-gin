package router

import (
	"github.com/gin-gonic/gin"
	"zcw-admin-server/app/controller"
	"zcw-admin-server/middleware"
)

// Api api
func Api(r *gin.Engine) {

	// 公开接口
	api := r.Group("/api")
	api.GET("/test", controller.Index.Test)

	// 权限接口
	api.Use(middleware.JWT)
	api.GET("/login", controller.Index.Test)
}
