GIT_COMMIT=`git rev-parse --short HEAD`
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X go-app-core/demo/version.gitCommit=$(GIT_COMMIT) -X go-app-core/demo/version.buildTime=$(BUILD_TIME)"
all:
	go build $(LDFLAGS)
