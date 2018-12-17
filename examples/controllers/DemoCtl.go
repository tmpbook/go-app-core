package controllers

import (
	"fmt"
	"net/http"

	"github.com/tmpbook/go-app-core/core/common/jsonConfig"
)

// DemoController R.T.
func DemoController(w http.ResponseWriter, r *http.Request) error {
	content, err := jsonConfig.GetConfigByKey("content.say-hello")
	if err != nil {
		return fmt.Errorf("when get config key: %v", err)
	}
	fmt.Fprintln(w, content.(string))
	return nil
}
