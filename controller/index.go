package controller

import "github.com/gin-gonic/gin"

var Index = new(IndexController)

type IndexController struct {
	Controller
}

func (c *IndexController) Test(r *gin.Context) {
	c.Success(r, true)
}
