package main

import (
	"flag"
	"fmt"
	"log"
	"config"

	"github.com/spf13/viper"
)

const (
	appName = "go-study"
	Version = "0.1.0"
	buildTime
)

func main() {

	// load common line args
	flag.Parse()

	// recover
	defer func() {
		if r := recover(); r != nil {
			log.Println(appName, "Process Stop ", r)
		} else {
			log.Println(appName, "Process Stop")
		}
	}()

	//
	config.InitConfig()

	fmt.Println(viper.Get("app.name"))
}
