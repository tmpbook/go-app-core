package controllers

import (
	"fmt"
	"net/http"

	"github.com/tmpbook/go-app-core/core/common/jsonConfig"
)

// ReloadConfigController R.T.
func ReloadConfigController(w http.ResponseWriter, r *http.Request) error {
	content, err := jsonConfig.GetConfigByKey("content.say-hello")
	if err != nil {
		return fmt.Errorf("when get config by key: %v", err)
	}
	fmt.Fprintln(w, "old:", content.(string))

	err = jsonConfig.ReloadConfig()
	if err != nil {
		return fmt.Errorf("when reload config: %v", err)
	}

	content, err = jsonConfig.GetConfigByKey("content.say-hello")
	if err != nil {
		return fmt.Errorf("when get config by key: %v", err)
	}
	fmt.Fprintln(w, "new:", content.(string))

	return nil
}
