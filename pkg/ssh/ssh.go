package ssh

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/tmpbook/go-app-core/pkg/grpool"
	"golang.org/x/crypto/ssh"
)

// SSHExecAgent worker
type SSHExecAgent struct {
	Worker  int
	TimeOut time.Duration
}

// PublicKeyFile 获取
func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

// GetAuthPassword 密码认证
func GetAuthPassword(password string) []ssh.AuthMethod {
	return []ssh.AuthMethod{ssh.Password(password)}
}

// GetAuthKeys key 认证
func GetAuthKeys(keys []string) []ssh.AuthMethod {
	methods := []ssh.AuthMethod{}
	for _, keyname := range keys {
		pkey := PublicKeyFile(keyname)
		if pkey != nil {
			methods = append(methods, pkey)
		}
	}
	return methods
}

// SSHHostByKey 通过 key 认证
func (s *SSHExecAgent) SSHHostByKey(hosts []string, port int, user string, cmd string) ([]ExecResult, error) {
	if len(hosts) == 0 {
		return nil, errors.New("no hosts given")
	}
	if s.Worker == 0 {
		s.Worker = 40
	}
	if s.TimeOut == 0 {
		s.TimeOut = 3600 * time.Second
	}
	keys := []string{
		os.Getenv("HOME") + "/.ssh/id_dsa",
		os.Getenv("HOME") + "/.ssh/id_rsa",
	}
	authKeys := GetAuthKeys(keys)
	if len(authKeys) < 1 {
		return nil, errors.New("the user has no key")
	}
	pool := grpool.NewPool(s.Worker, len(hosts), s.TimeOut)
	defer pool.Release()
	pool.WaitCount(len(hosts))
	for i := range hosts {
		count := i
		pool.JobQueue <- grpool.Job{
			JobID: count,
			JobFunc: func() (interface{}, error) {
				session := &HostSession{
					Username: user,
					Password: "",
					Hostname: hosts[count],
					Port:     port,
					Auths:    authKeys,
				}
				r := session.Exec(count, cmd, session.GenerateConfig())
				return *r, nil
			},
		}
	}

	pool.WaitAll()
	returnResult := make([]ExecResult, len(hosts))
	errorText := ""
	for res := range pool.Jobresult {
		JobID, _ := res.JobID.(int)
		if res.Timedout {
			returnResult[JobID].ID = JobID
			returnResult[JobID].Host = hosts[JobID]
			returnResult[JobID].Command = cmd
			returnResult[JobID].Error = errors.New("ssh time out")
			errorText += "the host " + hosts[JobID] + " command exec time out."
		} else {
			execResult, _ := res.Result.(ExecResult)
			returnResult[JobID] = execResult
			if execResult.Error != nil {
				errorText += "Host " + execResult.Host + " command exec error,result: " + execResult.Result + " Error info : " + execResult.Error.Error()
			}
		}
	}
	if errorText != "" {
		return returnResult, errors.New(errorText)

	}
	return returnResult, nil
}

// SFTPHostByKey 通过 SFTP 协议传文件
func (s *SSHExecAgent) SFTPHostByKey(hosts []string, port int, user string, localFilePath string, remoteFilePath string) ([]ExecResult, error) {
	if len(hosts) == 0 {
		return nil, errors.New("no hosts given")
	}
	if s.Worker == 0 {
		s.Worker = 40
	}
	if s.TimeOut == 0 {
		s.TimeOut = 3600 * time.Second
	}
	keys := []string{
		os.Getenv("HOME") + "/.ssh/id_dsa",
		os.Getenv("HOME") + "/.ssh/id_rsa",
	}
	authKeys := GetAuthKeys(keys)
	if len(authKeys) < 1 {
		log.Println("the user no key")
		return nil, errors.New("the user no key")
	}
	pool := grpool.NewPool(s.Worker, len(hosts), s.TimeOut)
	defer pool.Release()
	pool.WaitCount(len(hosts))
	for i, _ := range hosts {
		count := i
		pool.JobQueue <- grpool.Job{
			JobID: count,
			JobFunc: func() (interface{}, error) {
				session := &HostSession{
					Username: user,
					Password: "",
					Hostname: hosts[count],
					Port:     port,
					Auths:    authKeys,
				}
				r := session.Transfer(count, localFilePath, remoteFilePath, session.GenerateConfig())
				return *r, nil
			},
		}
	}

	pool.WaitAll()
	returnResult := make([]ExecResult, len(hosts))
	errorText := ""
	for res := range pool.Jobresult {
		JobID, _ := res.JobID.(int)
		if res.Timedout {
			returnResult[JobID].ID = JobID
			returnResult[JobID].Host = hosts[JobID]
			returnResult[JobID].LocalFilePath = localFilePath
			returnResult[JobID].RemoteFilePath = remoteFilePath
			returnResult[JobID].Error = errors.New("ssh time out")
			errorText += "the host " + hosts[JobID] + " command  exec time out."
		} else {
			execResult, _ := res.Result.(ExecResult)
			returnResult[JobID] = execResult
			if execResult.Error != nil {
				errorText += "the host " + execResult.Host + " command  exec error.\n" + "result info :" + execResult.Result + ".\nerror info :" + execResult.Error.Error()
			}
		}
	}
	if errorText != "" {
		return returnResult, errors.New(errorText)

	}
	return returnResult, nil
}
