package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"net/http"
	"zcw-gin/global"
)

func Limiter(r *gin.Context) {
	err := tollbooth.LimitByRequest(global.LIMITER, r.Writer, r.Request)
	if err != nil {
		r.JSON(http.StatusOK, gin.H{
			"code": global.ERROR,
			"msg":  "服务繁忙，请稍后再试...",
			"data": nil,
		})
		r.Abort()
	}
	r.Next()
}
