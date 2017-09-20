<h1>Go Application Core</h1>

本项目为 Go 程序提供基本的配置文件读取管理和命令行参数管理功能，并提供了 go 程序最佳实践

### Screenshots

#### terminal
![demo](images/terminal.png)

#### chrome

![chrome](images/chrome.png)
### usage

```bash
cd demo/
make
./demo -h 127.0.0.1 -p 8888 -c ./config.json
```

### check compile version
```bash
➜ ./demo -v
Git Commit: e6a6ba1
Build Time: 2017-09-20T19:23:19+0800
```

### run
```bash
➜ go run demo.go
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

2017/09/20 19:22:21 Running in: localhost:8910
```

### read from config.json

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
