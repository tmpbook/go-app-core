package main

import (
	"flag"
	"log"
	"net/http"

	"go-app-core/examples/controllers"

	"github.com/tmpbook/go-app-core/net/decorator"
	"github.com/tmpbook/go-app-core/utils/common"
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.String("p", "8910", "port")
)

func init() {
	// parse all flag that include common package
	// this statement must run first
	flag.Parse()

	// load config file which specified by flag -c
	common.LoadConfigFromFileAndWatch()

	// print versions if -v = true
	common.PrintVersion()

	// print all flags we used
	common.PrintFlags()

	// print config file content we loaded
	common.PrintConfig()
}

func main() {

	// start to play
	http.HandleFunc("/", decorator.ErrorCatcher(controllers.DemoController))
	http.HandleFunc("/reload-config", decorator.ErrorCatcher(controllers.ReloadConfigController))

	hostPort := *host + ":" + *port

	log.Print("Listinning on: ", hostPort)
	log.Fatal(http.ListenAndServe(hostPort, nil))
}
