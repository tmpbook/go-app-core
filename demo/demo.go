package main

import (
	"flag"
	"go-app-core/demo/config"
	"go-app-core/demo/controllers"
	"go-app-core/demo/version" // flag v
	"go-app-core/utils/common"
	"log"
	"net/http"
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
	version.Print()
	// 打印解析后的 flags
	common.PrintFlags()
	// 打印读取的配置文件
	common.PrintConfig()

	mux := http.NewServeMux()

	// 开始你的表演
	mux.Handle("/", &controllers.DemoController{})

	hostPort := *host + ":" + *port

	log.Print("Running in: ", hostPort)
	log.Fatal(http.ListenAndServe(hostPort, mux))
}
