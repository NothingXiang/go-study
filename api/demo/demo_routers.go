package demo

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	e.Group("/demo")
	{
		e.GET("/get", GetDemoHandler)
		e.POST("/set", SetDemoHandler)
	}
}
