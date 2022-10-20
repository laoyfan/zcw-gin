package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"time"
	"zcw-admin-server/app/entity"
	"zcw-admin-server/app/model/basic"
	"zcw-admin-server/app/service"
	"zcw-admin-server/global"
)

var Index = new(IndexController)

type IndexController struct {
	Controller
	indexService service.IndexService
}

func (c *IndexController) Test(r *gin.Context) {
	var params entity.Index
	if err := c.Valid(r, &params); err != nil {
		return
	}
	redis := global.REDIS["default"]
	redis.Set(context.Background(), "test", 444, 10*10*time.Second)
	i, _ := redis.Get(context.Background(), "test").Int()
	var userModel basic.UserModel
	user := userModel.GetByCondition()
	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	last := gjson.Get(json, "name.last")
	c.Success(r, map[string]interface{}{
		"user": user,
		"test": i,
		"last": last.String(),
	})
}
