/*
 *初始化配置
 */
package config

import (
	"github.com/spf13/viper"
)

// load config file
func InitConfig() {
	viper.SetConfigName("config.json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic("[Load Config]failed:%v" + err.Error())
	}
}
