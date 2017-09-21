<h1>Go Application Core</h1>

[![Build Status](https://travis-ci.org/tmpbook/go-app-core.svg?branch=master)](https://travis-ci.org/tmpbook/go-app-core) [![Go Report Card](https://goreportcard.com/badge/github.com/tmpbook/go-app-core)](https://goreportcard.com/report/github.com/tmpbook/go-app-core)

This project provides the basic configuration file management and command line parameter management for Go program, and continuing to pursue Go programming best practices.

[中文 README.md](README-zh.md)

### Usage

#### Step 1: Compile

```bash
cd demo/
make
./demo -h 127.0.0.1 -p 8888 -c ./config.json
```
Check compile version
```bash
➜ ./demo -v
Git Commit: e6a6ba1
Build Time: 2017-09-20T19:23:19+0800
```

#### Step 2: Run
```bash
➜ ./demo.go
=-------=
| flags |
=-------=
flag = c          value = config.json     help = configuration file, json format
flag = h          value = localhost       help = host
flag = p          value = 8910    help = port
flag = v          value = false           help = version
=-------=
|configs|
=-------=
{
  "version": "1.0",
  "content": {
    "say-hello": "Hello from config.json."
  }
}

2017/09/20 19:22:21 Listening on: localhost:8910
```

### Read from config.json

If we have a `config.json` file like below:
```json
{
    "version": "1.0",
    "content": {
        "say-hello": "Hello from config.json."
    }
}
```
Read it like this:
```go
content, _ := common.GetConfigByKey("content.say-hello")
```
As you see, it support get config with dot.

### Screenshots

#### Terminal
![demo](images/terminal.png)


#### Http test

![chrome](images/chrome.png)

#### Contribution Welcomed!
> [Report issue](https://github.com/tmpbook/go-app-core/issues/new) or pull request, or email nauy2011@126.com.