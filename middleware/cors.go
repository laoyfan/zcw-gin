package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors(r *gin.Context) {
	method := r.Request.Method
	origin := r.Request.Header.Get("Origin")
	r.Header("Access-Control-Allow-Origin", origin)
	r.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
	r.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
	r.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
	r.Header("Access-Control-Allow-Credentials", "true")

	// 放行所有OPTIONS方法
	if method == "OPTIONS" {
		r.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	r.Next()
}
