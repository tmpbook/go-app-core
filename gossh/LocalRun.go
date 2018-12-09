package gossh

import (
	"bytes"
	"errors"
	"os/exec"
	"time"
)

// LocalExec 在本地执行命令，无需目标 host
func LocalExec(cmd string, timeOutSecond int64) (ExecResult, error) {
	timeout := time.After(time.Duration(timeOutSecond) * time.Second)
	execResultCh := make(chan *ExecResult, 1)
	go func() {
		execResult := localExec(cmd)
		execResultCh <- &execResult
	}()
	select {
	case res := <-execResultCh:
		execRst := *res
		errorText := ""
		if execRst.Error != nil {
			errorText +=
				"Local exec error." + "\nResult info:" + execRst.Result + "\nError info:" + execRst.Error.Error()
		}
		if errorText != "" {
			return execRst, errors.New(errorText)
		}
		return execRst, nil
	case <-timeout:
		return ExecResult{Command: cmd, Error: errors.New("Local Exec time out")},
			errors.New("Local Exec time out")
	}
}

func localExec(cmd string) ExecResult {
	execResult := ExecResult{}
	execResult.StartTime = time.Now()
	execResult.Command = cmd
	execCommand := exec.Command("/bin/bash", "-c", cmd)
	var stdoutByte bytes.Buffer
	execCommand.Stdout = &stdoutByte
	var stderrByte bytes.Buffer
	execCommand.Stderr = &stderrByte
	err := execCommand.Run()
	if err != nil {
		execResult.Error = err
		execResult.Result = stderrByte.String()
	} else {
		execResult.EndTime = time.Now()
		execResult.Result = stdoutByte.String()
	}
	return execResult
}
