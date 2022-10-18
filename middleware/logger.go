package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
	"zcw-admin-server/global"
)

// LogLayout 日志layout
type LogLayout struct {
	Time      time.Time
	Metadata  map[string]interface{} // 存储自定义原数据
	Path      string                 // 访问路径
	Query     string                 // 携带query
	Body      string                 // 携带body数据
	IP        string                 // ip地址
	UserAgent string                 // 代理
	Error     string                 // 错误
	Cost      time.Duration          // 花费时间
	Source    string                 // 来源
}

type Log struct {
	// Filter 用户自定义过滤
	Filter func(c *gin.Context) bool
	// FilterKeyword 关键字过滤(key)
	FilterKeyword func(layout *LogLayout) bool
	// AuthProcess 鉴权处理
	AuthProcess func(c *gin.Context, layout *LogLayout)
	// 日志处理
	Print func(LogLayout)
	// Source 服务唯一标识
	Source string
}

func Logger(r *gin.Context) {
	start := time.Now()
	path := r.Request.URL.Path
	query := r.Request.URL.RawQuery
	r.Next()

	cost := time.Since(start)
	global.LOG.Info(path,
		zap.Int("status", r.Writer.Status()),
		zap.String("method", r.Request.Method),
		zap.String("path", path),
		zap.String("query", query),
		zap.String("ip", r.ClientIP()),
		zap.String("user-agent", r.Request.UserAgent()),
		zap.String("errors", r.Errors.ByType(gin.ErrorTypePrivate).String()),
		zap.Duration("cost", cost),
	)
}
