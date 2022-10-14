package controller

import (
	"github.com/gin-gonic/gin"
)

var Common = new(CommonController)

type CommonController struct {
	Controller
}

func (c *CommonController) Notify(r *gin.Context) {
	c.Success(r, "data")
}
