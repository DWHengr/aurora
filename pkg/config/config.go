package config

import (
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/DWHengr/aurora/pkg/misc/kafka"
	"github.com/DWHengr/aurora/pkg/misc/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Conf 配置文件
var Conf *Config

// DefaultPath 默认配置路径
var DefaultPath = "./configs/config.yml"

// Config 配置文件
type Config struct {
	Port  string            `yaml:"port"`
	Log   logger.LogConfig  `yaml:"log"`
	Mysql mysql.MysqlConfig `yaml:"mysql"`
	Kafka kafka.KafkaConfig `yaml:"kafka"`
}

// NewConfig 获取配置配置
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

	return Conf, nil
}
