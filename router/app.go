package router

import (
	"github.com/gin-gonic/gin"
	"github.com/overcls/blog/api"
)

func InitRouter() *gin.Engine {

	router :=gin.Default()

	router.GET("/",api.IndexController)

	return  router

}