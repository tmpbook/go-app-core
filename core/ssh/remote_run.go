package ssh

import "time"

// SSHWorker 可以通过 gossh.SSHWorker 修改
const SSHWorker = 10

// RemoteExec 执行远程命令，需要提供 hosts 列表
func RemoteExec(command string, runUser string, port int, hosts []string, timeOutSecond int64) ([]ExecResult, error) {
	sshExecAgent := SSHExecAgent{}
	sshExecAgent.Worker = SSHWorker
	sshExecAgent.TimeOut = time.Duration(timeOutSecond) * time.Second
	s, err := sshExecAgent.SSHHostByKey(hosts, port, runUser, command)
	return s, err
}
