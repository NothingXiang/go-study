package demo

import (
	"net/http"

	"github.com/NothingXiang/go-study/services"
	"github.com/gin-gonic/gin"
)

// todo: Add Implements
var service services.DemoService

func GetDemoHandler(c *gin.Context) {
	id := c.Query("id")

	m, err := service.GetDemo(id)
	if err != nil {
		c.JSON(http.StatusOK, m)
	} else {
		c.JSON(http.StatusNotFound, err)
	}
}
