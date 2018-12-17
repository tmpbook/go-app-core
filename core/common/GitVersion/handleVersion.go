package version

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

func init() {
	version = flag.Bool("v", false, "version")
}

// PrintVersion if import this(version) package
// you can run the function below to print version
// which injected by `-ldflags` when compile
func PrintVersion() {

	// prevent users from forgetting
	if !flag.Parsed() {
		flag.Parse()
	}

	if *version {
		fmt.Println("Git Commit: " + gitCommit)
		fmt.Println("Build Time: " + buildTime)
		os.Exit(0)
	}
}
