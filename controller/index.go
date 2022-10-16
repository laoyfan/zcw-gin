package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"zcw-admin-server/global"
)

var Index = new(IndexController)

type IndexController struct {
	Controller
}

func (c *IndexController) Test(r *gin.Context) {
	var user map[string]interface{}
	db := global.DB["default"]
	redis := global.REDIS["default"]
	err := redis.Set(context.Background(), "test", 444, 10*10*time.Second).Err()
	if err != nil {
		fmt.Println(err, 111)
	}

	i, _ := redis.Get(context.Background(), "test").Int()

	db.Table("user").Find(&user)
	fmt.Println(user, i)

	c.Success(r, map[string]interface{}{
		"user": user,
		"test": i,
	})
}
