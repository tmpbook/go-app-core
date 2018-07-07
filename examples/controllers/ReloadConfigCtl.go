package controllers

import (
	"fmt"
	"net/http"

	"github.com/tmpbook/go-app-core/utils/common"
)

// ReloadConfigController R.T.
func ReloadConfigController(w http.ResponseWriter, r *http.Request) error {
	content, err := common.GetConfigByKey("content.say-hello")
	if err != nil {
		return fmt.Errorf("when get config by key: %v", err)
	}
	fmt.Fprintln(w, "old:", content.(string))

	err = common.ReloadConfig()
	if err != nil {
		return fmt.Errorf("when reload config: %v", err)
	}

	content, err = common.GetConfigByKey("content.say-hello")
	if err != nil {
		return fmt.Errorf("when get config by key: %v", err)
	}
	fmt.Fprintln(w, "new:", content.(string))

	return nil
}
