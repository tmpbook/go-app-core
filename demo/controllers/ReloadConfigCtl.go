package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tmpbook/go-app-core/utils/common"
)

// ReloadConfigController R.T.
type ReloadConfigController struct{}

func (*ReloadConfigController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	content, _ := common.GetConfigByKey("content.say-hello")
	fmt.Fprintln(w, "old:", content.(string))

	_ = common.ReloadConfig()
	content, _ = common.GetConfigByKey("content.say-hello")
	fmt.Fprintln(w, "new:", content.(string))
}
