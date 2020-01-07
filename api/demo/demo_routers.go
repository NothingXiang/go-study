package demo

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	demo := e.Group("/demo")
	{
		demo.GET("/get", GetDemoHandler)
		demo.POST("/set", SetDemoHandler)
	}
}
