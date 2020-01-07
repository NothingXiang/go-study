package demo

import (
	"net/http"

	"github.com/NothingXiang/go-study/common/errors"
	"github.com/NothingXiang/go-study/models"
	"github.com/NothingXiang/go-study/services"
	"github.com/gin-gonic/gin"
)

var service services.DemoService = services.NewDemoServiceIml()

// todo : format return type
func GetDemoHandler(c *gin.Context) {
	id := c.Query("id")

	m, err := service.GetDemo(id)
	if err == nil {
		c.JSON(http.StatusOK, m)
	} else {
		c.JSON(http.StatusNotFound, errors.NotExist.SetMsg(err).Error())
	}
}

func SetDemoHandler(c *gin.Context) {
	// todo: can update to dto
	var dto *models.DemoModel

	if err := c.Bind(&dto); err != nil {
		c.JSON(http.StatusOK, errors.ParamFmt.SetMsg(err).Error())
		return
	}

	err := service.SetDemo(*dto)
	if err == nil {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusNotFound, errors.Unknown.SetMsg(err).Error())
	}
}
