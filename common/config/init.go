/*
 *初始化配置
 */
package config

import (
	"sync"

	"github.com/spf13/viper"
)

const (
	DefaultConfigFile = "config.json"
	DefaultConfigPath = "."
)

var once sync.Once

// load config file
func InitConfig() {
	// make sure config just init once
	once.Do(func() {
		viper.SetConfigName(DefaultConfigFile)
		viper.AddConfigPath(DefaultConfigPath)
		if err := viper.ReadInConfig(); err != nil {
			panic("[Load Config]failed:%v" + err.Error())
		}
	})
}
