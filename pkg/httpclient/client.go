package httpclient

import (
	"context"
	"net"
	"net/http"
	"time"
)

type HttpConfig struct {
	Timeout      time.Duration `yaml:"timeout"`
	MaxIdleConns int           `yaml:"maxIdleConns"`
}

var client *http.Client

func GetHttpClient() *http.Client {
	if client == nil {
		panic("http client is nil,first call NewClient() to initialize instance")
	}
	return client
}

// NewClient new a http client
func NewClient(conf *HttpConfig) {
	client = &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				deadline := time.Now().Add(conf.Timeout * time.Second)
				c, err := net.DialTimeout(network, addr, time.Second*conf.Timeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
			MaxIdleConns: conf.MaxIdleConns,
		},
	}
}
