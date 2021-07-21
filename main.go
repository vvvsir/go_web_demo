// @title go_web_demo
package main

import (
	"flag"
	"fmt"
	"go_web_demo/models"
	"go_web_demo/pkg/logger"
	"go_web_demo/pkg/setting"
	"go_web_demo/routers"
)

func main() {
	flag.Parse()

	// 初始化操作 (因为 init 方法无法保证我们想要的顺序)
	setting.Setup()
	logger.Setup()
	models.Setup()

	router := routers.InitRouter()
	panic(router.Run(fmt.Sprintf("%s:%d", setting.AppConf.Host, setting.AppConf.Port)))
}
