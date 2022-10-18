package router

import (
	"github.com/gin-gonic/gin"
	"zcw-admin-server/app/controller"
)

func Api(r *gin.Engine) {
	// api
	api := r.Group("/api")
	api.GET("/test", controller.Index.Test)
}
