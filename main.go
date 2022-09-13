package main

import (
	"github.com/gin-gonic/gin"
	"zcw-admin-server/router"
)

func main() {
	r := gin.New()
	router.Load(r)
	r.Run()
}
