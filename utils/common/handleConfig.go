package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// they has setter & getter
var (
	configByte []byte
	configMap  map[string]interface{}
	configFile string
)

type block struct {
	data interface{}
}

func (b *block) getValue(key string) *block {
	m := b.getData()
	if v, ok := m[key]; ok {
		b.data = v
		return b
	}
	return nil
}

func (b *block) getData() map[string]interface{} {
	if m, ok := (b.data).(map[string]interface{}); ok {
		return m
	}
	return nil
}

func getConfig(data map[string]interface{}) (*block, error) {
	if data == nil {
		return nil, errors.New("get config fail, config content is nil")
	}
	return &block{
		data: data,
	}, nil
}

// GetConfigByKey 从 JSONObject 中读取 key，支持点号
func GetConfigByKey(keys string) (value interface{}, err error) {
	keyList := strings.Split(keys, ".")
	block, err := getConfig(configMap)
	if err != nil {
		return nil, err
	}
	for _, key := range keyList {
		block = block.getValue(key)
		if block == nil {
			return nil, fmt.Errorf(`can not get [%s]'s value`, string(keys))
		}
	}
	return block.data, nil
}

// LoadConfigFromData 从 Byte 切片中读取配置，存入 configMap 变量
func LoadConfigFromData(data []byte) (err error) {
	err = json.Unmarshal(data, &configMap)
	return
}

// LoadConfigFromFile 从文件中读取配置，提供 Getter: GetConfigByKey 获取
func LoadConfigFromFile(filename string) (err error) {
	go watch()
	configFile = filename
	configByte, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	LoadConfigFromData(configByte)
	return
}

// PrintConfig 输出导入的 config 文件
func PrintConfig() {
	var prejson bytes.Buffer
	json.Indent(&prejson, configByte, "", "  ")
	fmt.Println("=-------=")
	fmt.Println("|configs|")
	fmt.Println("=-------=")
	fmt.Println(string(prejson.Bytes()))
}

// ReloadConfig 重新读取新的配置
func ReloadConfig() (err error) {
	log.Println("Reloading config...")
	err = LoadConfigFromFile(configFile)
	return
}

func watch() {
	l := log.New(os.Stderr, "", 0)

	// Catch SIGHUP to automatically reload cache
	sighup := make(chan os.Signal, 1)
	signal.Notify(sighup, syscall.SIGHUP)

	for {
		<-sighup
		l.Println("Caught SIGHUP, reloading config...")
		ReloadConfig()
	}
}
