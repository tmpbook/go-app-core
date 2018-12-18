package main

import (
	"fmt"
	"log"

	"github.com/tmpbook/go-app-core/pkg/common/jsonConfig"
)

func main() {
	err := jsonConfig.LoadConfigFromFile("test.json")
	if err != nil {
		log.Fatal("got error: ", err)
	}
	var boolValue interface{}
	boolValue, err = jsonConfig.GetConfigByKey("bool")

	var rst bool
	rst = boolValue.(bool)
	if rst {
		fmt.Println("got true")
	}
}
