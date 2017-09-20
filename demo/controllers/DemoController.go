package controllers

import (
	"fmt"
	"go-app-core/utils/common"
	"log"
	"net/http"
)

// DemoController R.T.
type DemoController struct{}

func (*DemoController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	content, _ := common.GetConfigByKey("content.say-hello")
	fmt.Fprintln(w, content.(string))
}
