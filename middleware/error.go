package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Error 处理异常捕获
func Error(r *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			r.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "异常,请稍后重试",
				"data": nil,
			})
			r.Abort()
		}
	}()
	r.Next()
}
