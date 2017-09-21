package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tmpbook/go-app-core/utils/common"
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
