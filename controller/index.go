package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"zcw-admin-server/utils"
)

var Index = new(IndexController)

type IndexController struct {
	Controller
}

func (c *IndexController) Test(r *gin.Context) {
	utils.WriteLog("ceshi", errors.New("1111111111"))
	c.Success(r, true)
}
