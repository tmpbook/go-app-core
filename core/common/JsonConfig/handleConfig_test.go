package configV

import (
	"testing"
)

func TestGet(t *testing.T) {
	_ = LoadConfigFromFile("test.json")
	t.Run("Get", func(t *testing.T) {
		t.Run("should successfully retrieve string", func(t *testing.T) {
			s, err := GetConfigByKey("string")
			if err != nil {
				t.Error(err)
			}

			if s.(string) != "asdf" {
				t.Errorf("Expected 'asdf', got '%v'", s)
			}

		})

		t.Run("should got a nil value and a error", func(t *testing.T) {
			s, err := GetConfigByKey("nonexistent")

			if err == nil {
				t.Error(err)
			}

			if s != nil {
				t.Errorf("Expected 'nil', got '%v'", s)
			}
		})

		t.Run("should got a error for nonexistent files", func(t *testing.T) {
			err := LoadConfigFromFile("nonexistant.json")
			if err == nil {
				t.Error("file shouldn't found")
			}
			err = ReloadConfig()
			if err == nil {
				t.Error("shouldn't reload success")
			}
		})

		t.Run("should successfully load and reload for existent files", func(t *testing.T) {
			err := LoadConfigFromFile("test.json")
			if err != nil {
				t.Error("file should successfully loaded")
			}
			err = ReloadConfig()
			if err != nil {
				t.Error("file should successfully reload")
			}
		})
	})
}
