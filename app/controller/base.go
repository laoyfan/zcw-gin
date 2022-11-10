package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"zcw-gin/global"
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
	c.Result(r, global.SUCCESS, "请求成功", data)
}

// Error 失败响应
func (c *Controller) Error(r *gin.Context, data interface{}) {
	c.Result(r, global.ERROR, "请求失败", data)
}

// Valid 参数校验
func (c *Controller) Valid(r *gin.Context, valid interface{}) error {
	if err := r.ShouldBind(valid); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.Result(r, global.VALID, "请求参数校验失败", c.removeTopStruct(errs.Translate(global.Trans)))
		} else {
			c.Result(r, global.VALID, "请求参数校验失败", err.Error())
		}
		return err
	}
	return nil
}

func (c *Controller) removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
