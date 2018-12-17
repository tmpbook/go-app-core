GIT_COMMIT=`git rev-parse --short HEAD`
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/tmpbook/go-app-core/utils/common.gitCommit=$(GIT_COMMIT) -X github.com/tmpbook/go-app-core/utils/common.buildTime=$(BUILD_TIME)"
all:
	go build $(LDFLAGS)
