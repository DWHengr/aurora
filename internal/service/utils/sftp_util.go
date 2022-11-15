package utils

import (
	"fmt"
	"github.com/DWHengr/aurora/pkg/config"
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/pkg/sftp"
	"io/ioutil"
	"net"
	"path"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func ReadRule() ([]byte, error) {
	//write local
	allConfig, _ := config.GetAllConfig()
	auroraConfig := allConfig.Aurora
	if auroraConfig.Remote == nil || len(strings.Trim(auroraConfig.Remote.PrometheusHostIp, " ")) <= 0 {
		ruleBytes, err := ioutil.ReadFile(auroraConfig.PrometheusRulePath)
		if err != nil {
			return ruleBytes, nil
		}
		return nil, err
	}

	//write remote
	var (
		err        error
		sftpClient *sftp.Client
	)
	sftpClient, err = connect(auroraConfig.Remote.PrometheusHostSshUsername, auroraConfig.Remote.PrometheusHostSshPassword,
		auroraConfig.Remote.PrometheusHostIp, auroraConfig.Remote.PrometheusHostSshPort)
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	defer sftpClient.Close()
	var remoteDir = auroraConfig.PrometheusRulePath
	var remoteFileName = path.Base("aurora_rule.yml")

	dstFile, err := sftpClient.Open(path.Join(remoteDir, remoteFileName))
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	defer dstFile.Close()
	var ruleBytes []byte
	dstFile.Read(ruleBytes)
	fmt.Println("read remote server finished!")
	return ruleBytes, nil
}

func WriteRule(ruleBytes []byte) error {
	//write local
	allConfig, _ := config.GetAllConfig()
	auroraConfig := allConfig.Aurora
	if len(strings.Trim(auroraConfig.Remote.PrometheusHostIp, " ")) <= 0 {
		if err := ioutil.WriteFile(auroraConfig.PrometheusRulePath, ruleBytes, 0666); err != nil {
			logger.Logger.Error(err)
			return err
		}
		return nil
	}

	//write remote
	var (
		err        error
		sftpClient *sftp.Client
	)
	sftpClient, err = connect(auroraConfig.Remote.PrometheusHostSshUsername, auroraConfig.Remote.PrometheusHostSshPassword,
		auroraConfig.Remote.PrometheusHostIp, auroraConfig.Remote.PrometheusHostSshPort)
	if err != nil {
		logger.Logger.Error(err)
		return err
	}
	defer sftpClient.Close()
	var remoteDir = auroraConfig.PrometheusRulePath
	var remoteFileName = path.Base("aurora_rule.yml")

	sftpClient.Remove(path.Join(remoteDir, remoteFileName))
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		logger.Logger.Error(err)
		return err
	}
	defer dstFile.Close()

	dstFile.Write(ruleBytes)

	fmt.Println("write remote server finished!")
	return nil
}

func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil
}
