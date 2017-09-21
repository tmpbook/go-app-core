package common

import (
	"flag"
	"fmt"
	"os"
)

var (
	gitCommit = "gitCommit"
	buildTime = "2000-01-01T00:00:00+0800"
	version   *bool
)

// Init 添加 flag
func init() {
	version = flag.Bool("v", false, "version")
}

// PrintVersion 如果调用了 init，那么用户可以控制是否输出当前版本
func PrintVersion() {
	if *version {
		fmt.Println("Git Commit: " + gitCommit)
		fmt.Println("Build Time: " + buildTime)
		os.Exit(0)
	}
}
