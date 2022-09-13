package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Controller 基础控制器
// 此处封装请求响应
type Controller struct{}

// Result 基础封装
func (c *Controller) Result(r *gin.Context, code int, msg string, data interface{}) {
	r.JSON(http.StatusOK, Response{
		code, msg, data,
	})
}

// Success 成功响应
func (c *Controller) Success(r *gin.Context, data interface{}) {
	c.Result(r, 200, "success", data)
}

// Error 失败响应
func (c *Controller) Error(r *gin.Context, data interface{}) {
	c.Result(r, -1, "error", data)
}
