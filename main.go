package main

import (
	"github.com/gin-gonic/gin"
	"github.com/overcls/blog/api"
)

func main() {
	r:=gin.Default()


	r.GET("/",api.IndexController)
	_ = r.Run(":3000")
}
