package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"zcw-gin/global"
	"zcw-gin/pkg/jwt"
)

func JWT(r *gin.Context) {
	aToken := r.Request.Header.Get("access_token")
	rToken := r.Request.Header.Get("refresh_token")
	if aToken == "" || rToken == "" {
		r.JSON(http.StatusOK, gin.H{
			"code": global.FORBIDDEN,
			"msg":  "未登陆或非法访问",
			"data": nil,
		})
		r.Abort()
		return
	}
	userInfo, err := jwt.ParseToken(aToken)
	fmt.Println(userInfo, "5555555555555")
	if err != nil {
		aToken, rToken, err = jwt.RefreshToken(aToken, rToken)
		if err != nil {
			r.JSON(http.StatusOK, gin.H{
				"code": global.FORBIDDEN,
				"msg":  err.Error(),
				"data": nil,
			})
			r.Abort()
			return
		}
		r.Header("access_token", aToken)
		r.Header("refresh_token", rToken)
	}
	r.Set("userInfo", userInfo)
	r.Next()
}
