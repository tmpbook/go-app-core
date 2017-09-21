package config

import (
	"flag"
	"log"

	"github.com/tmpbook/go-app-core/utils/common"
)

var (
	confFile *string
)

// Init 初始化读取配置文件
func init() {
	confFile = flag.String("c", "config.json", "configuration file, json format")
}

// Load 读取 flag 指定的文件
func Load() {
	err := common.LoadConfigFromFile(*confFile)
	if err != nil {
		log.Fatal(err)
	}
}
