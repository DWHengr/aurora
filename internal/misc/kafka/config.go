package kafka

import (
	"crypto/tls"

	"github.com/Shopify/sarama"
)

// Config config
type KafkaConfig struct {
	Sarama sarama.Config

	Broker []string `yaml:"broker"`
	Topic  string   `yaml:"topic"`
	TLS    *tls.Config
}

func pre(conf KafkaConfig) *sarama.Config {
	config := sarama.NewConfig()

	// TLS
	config.Net.TLS.Enable = conf.Sarama.Net.TLS.Enable
	config.Net.TLS.Config = conf.Sarama.Net.TLS.Config

	return config
}
