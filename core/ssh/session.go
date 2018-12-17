package ssh

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// HostSession session 信息结构体
type HostSession struct {
	Username string
	Password string
	Hostname string
	Signers  []ssh.Signer
	Port     int
	Auths    []ssh.AuthMethod
}

// ExecResult 命令执行结果详情结构体
type ExecResult struct {
	ID             int
	Host           string
	Command        string
	LocalFilePath  string
	RemoteFilePath string
	Result         string
	StartTime      time.Time
	EndTime        time.Time
	Error          error
}

// Exec execute the command and return a result structure
func (exec *HostSession) Exec(id int, command string, config ssh.ClientConfig) *ExecResult {

	result := &ExecResult{
		ID:      id,
		Host:    exec.Hostname,
		Command: command,
	}

	client, err := ssh.Dial("tcp", exec.Hostname+":"+strconv.Itoa(exec.Port), &config)

	if err != nil {
		result.Error = err
		return result
	}

	session, err := client.NewSession()

	if err != nil {
		result.Error = err
		return result
	}

	defer session.Close()

	var b bytes.Buffer

	session.Stdout = &b
	var b1 bytes.Buffer
	session.Stderr = &b1
	start := time.Now()
	if err := session.Run(command); err != nil {
		result.Error = err
		result.Result = b1.String()
		return result
	}
	end := time.Now()
	result.Result = b.String()
	result.StartTime = start
	result.EndTime = end
	return result
}

// Transfer 使用 sftp 传输文件
func (exec *HostSession) Transfer(id int, localFilePath string, remoteFilePath string, config ssh.ClientConfig) *ExecResult {

	result := &ExecResult{
		ID:             id,
		Host:           exec.Hostname,
		LocalFilePath:  localFilePath,
		RemoteFilePath: remoteFilePath,
	}
	start := time.Now()
	result.StartTime = start
	client, err := ssh.Dial("tcp", exec.Hostname+":"+strconv.Itoa(exec.Port), &config)

	if err != nil {
		result.Error = err
		return result
	}

	session, err := client.NewSession()
	if err != nil {
		result.Error = err
		return result
	}
	defer session.Close()

	var fileSize int64
	if s, err := os.Stat(localFilePath); err != nil {
		result.Error = err
		return result
	} else {
		fileSize = s.Size()
	}

	srcFile, err := os.Open(localFilePath)
	if err != nil {
		result.Error = err
		return result
	}
	defer srcFile.Close()

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		result.Error = err
		return result
	}
	defer sftpClient.Close()

	dstFile, err := sftpClient.Create(remoteFilePath)
	if err != nil {
		result.Error = err
		return result
	}
	defer dstFile.Close()

	n, err := io.Copy(dstFile, io.LimitReader(srcFile, fileSize))
	if err != nil {
		result.Error = err
		return result
	}
	if n != fileSize {
		result.Error = fmt.Errorf("copy: expected %v bytes, got %d", fileSize, n)
		return result
	}
	end := time.Now()
	result.EndTime = end
	return result
}

// GenerateConfig 生成配置
func (exec *HostSession) GenerateConfig() ssh.ClientConfig {
	var auths []ssh.AuthMethod

	if len(exec.Password) != 0 {
		auths = append(auths, ssh.Password(exec.Password))
	} else {
		if len(exec.Auths) > 0 {
			auths = exec.Auths
		} else {
			auths = append(auths, ssh.PublicKeys(exec.Signers...))
		}
	}

	config := ssh.ClientConfig{
		User: exec.Username,
		Auth: auths,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// config.Ciphers = []string{"aes128-cbc", "3des-cbc"}

	return config
}

func readKey(filename string) (ssh.Signer, error) {
	f, _ := os.Open(filename)

	bytes, _ := ioutil.ReadAll(f)
	return generateKey(bytes)
}

func generateKey(keyData []byte) (ssh.Signer, error) {
	return ssh.ParsePrivateKey(keyData)
}
