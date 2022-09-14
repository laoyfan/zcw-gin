package main

import (
	"github.com/gin-gonic/gin"
	_ "zcw-admin-server/boot"
	"zcw-admin-server/router"
)

func main() {
	r := gin.New()
	router.Load(r)
	r.Run()
}
