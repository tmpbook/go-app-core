package gossh

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 请修改为自己的机器
const remoteUser = "root"
const remoteHost = "10.211.55.17"

func TestLocalExec(t *testing.T) {
	cmd := "hostname"

	// expected
	exec := exec.Command("/bin/bash", "-c", cmd)
	var stdoutByte bytes.Buffer
	exec.Stdout = &stdoutByte
	err := exec.Run()
	assert.Equal(t, nil, err)

	// actual
	rst, err := LocalExec(cmd, 60)
	assert.Equal(t, nil, err)

	// compare
	assert.Equal(t, stdoutByte.String(), rst.Result)
}

func TestRemoteExec(t *testing.T) {
	cmd := "whoami"
	user := remoteUser
	host := remoteHost

	rst, err := RemoteExec(cmd, user, 22, []string{host}, 10)
	assert.Equal(t, nil, err)

	// compare
	assert.Equal(t, fmt.Sprintf("%s\n", user), rst[0].Result)
}
