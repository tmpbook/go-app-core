<h1>Go Application Core</h1>

[中文 README.md](README-zh.md)

This project provides the basic configuration file management and command line parameter management for Go program, and continuing to pursue Go programming best practices.

### Screenshots

#### Terminal
![demo](images/terminal.png)

#### Chrome

![chrome](images/chrome.png)
### usage

```bash
cd demo/
make
./demo -h 127.0.0.1 -p 8888 -c ./config.json
```

### Check compile version
```bash
➜ ./demo -v
Git Commit: e6a6ba1
Build Time: 2017-09-20T19:23:19+0800
```

### Run
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
