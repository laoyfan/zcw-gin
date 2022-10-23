package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"time"
	"zcw-gin/app/entity"
	"zcw-gin/app/model/basic"
	"zcw-gin/app/service"
	"zcw-gin/global"
	"zcw-gin/pkg/jwt"
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

func (c *IndexController) Login(r *gin.Context) {
	var loginReq entity.LoginReq
	if err := c.Valid(r, &loginReq); err != nil {
		return
	}
	access, refresh, err := jwt.CreateToken(global.UserInfo{
		Username:    loginReq.Username,
		UID:         1,
		AuthorityId: 0,
	})
	if err != nil {
		fmt.Println(err)
		c.Error(r, "登录失败")
		return
	}
	c.Success(r, entity.LoginResp{
		Username: loginReq.Username,
		Access:   access,
		Refresh:  refresh,
	})
}
