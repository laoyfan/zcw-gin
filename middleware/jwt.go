package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zcw-admin-server/global"
	"zcw-admin-server/pkg/jwt"
)

func JWT(r *gin.Context) {
	token := r.Request.Header.Get("x-token")
	if token == "" {
		r.JSON(http.StatusOK, gin.H{
			"code": global.FORBIDDEN,
			"msg":  "未登陆或非法访问",
			"data": nil,
		})
		r.Abort()
		return
	}
	userInfo, err := jwt.ParseToken(token)
	if err != nil {
		r.JSON(http.StatusOK, gin.H{
			"code": global.FORBIDDEN,
			"msg":  err.Error(),
			"data": nil,
		})
		r.Abort()
		return
	}
	createToken, _ := jwt.CreateToken(*userInfo)
	oldToken, _ := jwt.CreateTokenByOldToken(createToken, *userInfo)
	r.Set("claims", userInfo)
	r.Set("token", oldToken)
	r.Next()
}
