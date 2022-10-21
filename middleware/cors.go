package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zcw-admin-server/global"
)

func Cors(r *gin.Context) {
	// debug 模式 放行全部
	if global.CONFIG.App.Mode == "debug" {
		setHeader(r)
		r.Next()
		return
	}
	// 开启校验
	if checkCors(r) {
		// 设置请求头
		setHeader(r)
	} else {
		// 拒绝请求
		r.AbortWithStatus(http.StatusForbidden)
	}
	// 放行所有OPTIONS方法
	if r.Request.Method == "OPTIONS" {
		r.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	r.Next()
}

// 允许跨域设置
func setHeader(r *gin.Context) {
	r.Header("Access-Control-Allow-Origin", r.GetHeader("origin"))
	r.Header("Access-Control-Allow-Headers", global.CONFIG.App.Cors.AllowHeaders)
	r.Header("Access-Control-Allow-Methods", global.CONFIG.App.Cors.AllowMethods)
	r.Header("Access-Control-Expose-Headers", global.CONFIG.App.Cors.ExposeHeaders)
	r.Header("Access-Control-Allow-Credentials", global.CONFIG.App.Cors.AllowCredentials)
	r.Header("Access-Control-Max-Age", global.CONFIG.App.Cors.MaxAge)
}

// 校验跨域
func checkCors(r *gin.Context) bool {
	for _, o := range global.CONFIG.App.Cors.AllowOrigins {
		if r.GetHeader("origin") == o {
			return true
		}
	}
	return false
}
