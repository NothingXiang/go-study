package main

import (
	"flag"
	"log"
	"time"

	"github.com/NothingXiang/go-study/api"
	"github.com/NothingXiang/go-study/common/config"
	"github.com/NothingXiang/go-study/common/pkg"
)

var (
	PkgInfo = pkg.Info{
		AppName:   "go-study",
		Version:   "0.1.0",
		StartTime: time.Now(),
	}
)

func main() {

	// 1. load common line args
	// todo:may be can update to cobra
	flag.Parse()

	//2. recover
	defer func() {
		if r := recover(); r != nil {
			log.Println(PkgInfo.AppName, "Process Stop ", r)
		} else {
			log.Println(PkgInfo.AppName, "Process Stop")
		}
	}()

	// 3. load config
	config.InitConfig()

	// 4. set routers
	api.Serve(PkgInfo)

}
