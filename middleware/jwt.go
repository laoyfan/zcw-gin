package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"zcw-gin/global"
	"zcw-gin/pkg/jwt"
)

func JWT(r *gin.Context) {
	Authorization := r.Request.Header.Get("Authorization")
	if Authorization == "" {
		r.JSON(http.StatusOK, gin.H{
			"code": global.FORBIDDEN,
			"msg":  "未登陆或非法访问",
			"data": nil,
		})
		r.Abort()
		return
	}

	tokens := strings.SplitN(Authorization, " ", 2)

	token := tokens[1]

	userInfo, err := jwt.ParseToken(token)
	if err != nil {
		return
	}
	if err != nil {
		r.JSON(http.StatusOK, gin.H{
			"code": global.FORBIDDEN,
			"msg":  err.Error(),
			"data": nil,
		})
		r.Abort()
		return
	}
	r.Set("userInfo", userInfo)
	r.Next()
}
