package controller

import (
	"github.com/gin-gonic/gin"
	"zcw-admin-server/boot"
)

var Common = new(CommonController)

type CommonController struct {
	Controller
}

func (c *CommonController) Notify(r *gin.Context) {
	data := boot.ConfigMap
	c.Success(r, data)
}
