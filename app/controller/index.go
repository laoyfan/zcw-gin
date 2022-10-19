package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
	"zcw-admin-server/app/model/basic"
	"zcw-admin-server/global"
)

var Index = new(IndexController)

type IndexController struct {
	Controller
}

func (c *IndexController) Test(r *gin.Context) {
	redis := global.REDIS["default"]
	redis.Set(context.Background(), "test", 444, 10*10*time.Second)
	i, _ := redis.Get(context.Background(), "test").Int()
	var userModel basic.UserModel
	user := userModel.GetByCondition()
	c.Success(r, map[string]interface{}{
		"user": user,
		"test": i,
	})
}
