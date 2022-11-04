package config

import (
	"errors"
	"github.com/DWHengr/aurora/internal"
	"github.com/DWHengr/aurora/internal/alertcore"
	"github.com/DWHengr/aurora/pkg/httpclient"
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/DWHengr/aurora/pkg/misc/email"
	"github.com/DWHengr/aurora/pkg/misc/kafka"
	"github.com/DWHengr/aurora/pkg/misc/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Conf 配置文件
var Conf *Config

// DefaultPath 默认配置路径
var DefaultPath = "../../configs/config.yml"

var allConfig *Config

// Config 配置文件
type Config struct {
	Port       string                `yaml:"port"`
	Log        logger.LogConfig      `yaml:"log"`
	Mysql      mysql.MysqlConfig     `yaml:"mysql"`
	Kafka      kafka.KafkaConfig     `yaml:"kafka"`
	Alert      alertcore.AlertConfig `yaml:"alert"`
	Email      email.EmailConfig     `yaml:"email"`
	HttpClient httpclient.HttpConfig `yaml:"httpclient"`
	Aurora     internal.AuroraConfig `yaml:"aurora"`
}

//GetAllConfig Get all configurations
func GetAllConfig() (*Config, error) {
	if allConfig == nil {
		return nil, errors.New("config is nil")
	}
	return allConfig, nil
}

// NewConfig
func NewConfig(path string) (*Config, error) {
	if path == "" {
		path = DefaultPath
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &Conf)
	if err != nil {
		return nil, err
	}
	allConfig = Conf
	return Conf, nil
}
