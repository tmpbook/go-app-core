package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/tmpbook/go-app-core/demo/config"
	"github.com/tmpbook/go-app-core/demo/controllers"

	"github.com/tmpbook/go-app-core/net/decorator"
	"github.com/tmpbook/go-app-core/utils/common"
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.String("p", "8910", "port")
)

func init() {
	// 解析 flag，包含import package 中的 init
	flag.Parse()
	// flag 解析完成后，读取配置文件（因为配置文件是通过 flag 来指定的）
	config.Load()
}

func main() {
	// -v 打印 version
	common.PrintVersion()
	// 打印解析后的 flags
	common.PrintFlags()
	// 打印读取的配置文件
	common.PrintConfig()

	// 开始你的表演
	http.HandleFunc("/", decorator.ErrorCatcher(controllers.DemoController))
	http.HandleFunc("/reload-config", decorator.ErrorCatcher(controllers.ReloadConfigController))

	hostPort := *host + ":" + *port

	log.Print("Listinning on: ", hostPort)
	log.Fatal(http.ListenAndServe(hostPort, nil))
}
