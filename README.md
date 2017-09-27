<h1>Go Application Core</h1>

[![Build Status](https://travis-ci.org/tmpbook/go-app-core.svg?branch=master)](https://travis-ci.org/tmpbook/go-app-core) [![Go Report Card](https://goreportcard.com/badge/github.com/tmpbook/go-app-core)](https://goreportcard.com/report/github.com/tmpbook/go-app-core)

This project provides the basic configuration file management and command line parameter management for Go program, and continuing to pursue Go programming best practices.

[中文 README.md](README-zh.md)

### Usage

#### Step 1: Go get

```
go get github.com/tmpbook/go-app-core/utils/common
```

#### Step 2: Write your app
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

#### Step 3: Compile it

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

#### Step 4: Run it
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

### Read config which default is ./config.json or specified by flag `-c`

If you have a [config.json](/demo/config.json) file like below:
```json
{
    "version": "1.0",
    "content": {
        "say-hello": "Hello from config.json."
    }
}
```

Get values:
```go
content, _ := common.GetConfigByKey("content.say-hello")
```
As you see, it support dot.

### Demo Screenshots

#### Terminal
![demo](/images/terminal.png)


#### Testing read config file by HTTP

![chrome](/images/chrome.png)

#### Contribution Welcomed!
> [Report issue](https://github.com/tmpbook/go-app-core/issues/new) or pull request, or email nauy2011@126.com