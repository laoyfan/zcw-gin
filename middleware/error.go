package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"zcw-gin/global"
)

// Error 处理异常捕获
func Error(r *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			var brokenPipe bool
			if ne, ok := err.(*net.OpError); ok {
				if se, ok := ne.Err.(*os.SyscallError); ok {
					if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
						brokenPipe = true
					}
				}
			}

			httpRequest, _ := httputil.DumpRequest(r.Request, false)
			if brokenPipe {
				global.LOG.Error(r.Request.URL.Path,
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
				)
				r.JSON(http.StatusOK, gin.H{
					"code": global.ERROR,
					"msg":  "异常,请稍后重试",
					"data": nil,
				})
				r.Abort()
				return
			}

			global.LOG.Error("[Recovery from panic]",
				zap.Any("error", err),
				zap.String("request", string(httpRequest)),
				zap.String("stack", string(debug.Stack())),
			)
			r.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	r.Next()
}
