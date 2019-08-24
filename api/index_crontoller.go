package api

import (
	"github.com/gin-gonic/gin"
	"github.com/overcls/blog/service"
	"net/http"
)

func IndexController(c *gin.Context)  {


	var indexService service.IndexService
	if err:=c.ShouldBindQuery(&indexService);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"data":indexService.Name,
	})

}