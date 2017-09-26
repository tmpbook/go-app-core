package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	// flag params
	confFile *string
	watch    *bool

	// they has setter & getter
	configByte []byte
	configMap  map[string]interface{}
	configFile string
)

func init() {
	confFile = flag.String("c", "config.json", "configuration file, json format")
	watch = flag.Bool("w", false, "reload config file by signal (kill -s SIGHUP [pid])")
}

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

// LoadConfigFromData read config from []byte, and store it into configMap
func LoadConfigFromData(data []byte) (err error) {
	err = json.Unmarshal(data, &configMap)
	return
}

// LoadConfigFromFile read config from json file
func LoadConfigFromFile(filename string) (err error) {
	configFile = filename
	configByte, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	LoadConfigFromData(configByte)
	return
}

// LoadConfigFromFileAndWatch read configuration from json file, and through the -w flag to decide whether to watch the signal
func LoadConfigFromFileAndWatch() (err error) {

	// prevent users from forgetting
	if !flag.Parsed() {
		flag.Parse()
	}

	err = LoadConfigFromFile(*confFile)
	if err != nil {
		log.Fatal(err)
	}

	if *watch {
		go watchReload()
		log.Println("Starting watching signal...")
	}
	return
}

// PrintConfig display all configs
func PrintConfig() {
	var prejson bytes.Buffer
	json.Indent(&prejson, configByte, "", "  ")
	fmt.Println("---------")
	fmt.Println(" configs")
	fmt.Println("---------")
	fmt.Println(string(prejson.Bytes()))
}

// ReloadConfig 重新读取新的配置
func ReloadConfig() (err error) {
	err = LoadConfigFromFile(configFile)
	return
}

// WatchReload watch signal to reload config file
func watchReload() {
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
