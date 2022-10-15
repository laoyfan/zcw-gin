package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zcw-admin-server/core"
)

var Index = new(IndexController)

type IndexController struct {
	Controller
}

func (c *IndexController) Test(r *gin.Context) {
	var user map[string]interface{}
	db := core.DB["default"]
	db.Table("user").Find(&user)
	fmt.Println(user)

	c.Success(r, user)
}
