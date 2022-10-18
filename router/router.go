package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zcw-admin-server/global"
	"zcw-admin-server/middleware"
)

func Server() *http.Server {
	r := gin.New()
	// 请求预处理
	r.Use(
		middleware.Logger, //日志处理
		middleware.Cors,   //跨域处理
		middleware.Error,  //异常处理
	)

	Api(r)

	return &http.Server{
		Addr:    global.CONFIG.App.Port,
		Handler: r,
	}
}
