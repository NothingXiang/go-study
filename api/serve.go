/*
 *  提供api接口
 */
package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NothingXiang/go-study/api/demo"
	"github.com/NothingXiang/go-study/common/pkg"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Serve(info pkg.Info) {

	// 初始化
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	// 注册路由
	engine.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, info)
	})
	demo.RegisterRoutes(engine)

	port := viper.GetString("http.port")
	if err := engine.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Printf("GIN Server Fail:%v", err)
	}
}
