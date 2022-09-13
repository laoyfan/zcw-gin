package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(r *gin.Context) {
	fmt.Println("1111111111111111111")
	r.Next()
}
