

<h1>Go Application Core</h1>

[![Build Status](https://travis-ci.org/tmpbook/go-app-core.svg?branch=master)](https://travis-ci.org/tmpbook/go-app-core) [![Go Report Card](https://goreportcard.com/badge/github.com/tmpbook/go-app-core)](https://goreportcard.com/report/github.com/tmpbook/go-app-core)

go-app-core 只需两行代码，即可为 Go 程序提供必不可少的配置文件管理、命令行参数管理、编译版本管理等功能，追求 Go 编程最「简」实践。

[中文 README.md](README-zh.md)

### 如何使用

#### 第一步

```
go get github.com/tmpbook/go-app-core/utils/common
```

#### 第二步
```go
package main

import (
    "github.com/tmpbook/go-app-core/utils/common"
    ...
)

func init() {
	// parse all flag that include common package
	// this statement must run first
	flag.Parse()

	// default is ./config.json, you can specified it by flag -c, watch signal to reload config file(CMD:kill -s SIGHUP [pid]) by add -w when start 
	common.LoadConfigFromFileAndWatch()

	// print versions if -v = true
	common.PrintVersion()

	// print all flags we used
	common.PrintFlags()

	// print config file content we loaded
	common.PrintConfig()
}

func main() {
  ...
}
```

#### 第三步：编译

[Makefile](/demo/Makefile)
```
GIT_COMMIT=`git rev-parse --short HEAD`
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/tmpbook/go-app-core/utils/common.gitCommit=$(GIT_COMMIT) -X github.com/tmpbook/go-app-core/utils/common.buildTime=$(BUILD_TIME)"
all:
	go build $(LDFLAGS)
```
and then
```bash
cd demo/
make
```

Check compile version
```bash
➜ ./demo -v
Git Commit: e6a6ba1
Build Time: 2017-09-20T19:23:19+0800
```

#### 第四步：执行
```bash
➜ ./demo.go
-------
 flags
-------
flag = c          value = config.json      configuration file, json format
flag = h          value = localhost        host
flag = p          value = 8910             port
flag = v          value = false            version
flag = w          value = false            reload config file by signal (kill -s SIGHUP [pid])
---------
 configs
---------
{
  "version": "1.0",
  "content": {
    "say-hello": "Hello from config.json."
  }
}

2017/09/20 19:22:21 Listening on: localhost:8910
```

### 读取通过 flag `-c` 指定的配置（如果不指定则为 `./config.json`）

假设我们有一个 [config.json](/demo/config.json) 如下：
```json
{
    "version": "1.0",
    "content": {
        "say-hello": "Hello from config.json."
    }
}
```

这样取值即可
```go
content, _ := common.GetConfigByKey("content.say-hello")
```
如你所见，它支持点号

### 示例程序截图

#### 终端
![demo](/images/terminal.png)


#### 测试 HTTP 请求访问配置文件

![chrome](/images/chrome.png)

#### 欢迎贡献你的代码
> [Report issue](https://github.com/tmpbook/go-app-core/issues/new) or pull request, or email nauy2011@126.com.