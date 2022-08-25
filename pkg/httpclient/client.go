package httpclient

import (
	"context"
	"net"
	"net/http"
	"time"
)

type HttpConfig struct {
	Timeout      time.Duration
	MaxIdleConns int
}

// NewClient new a http client
func NewClient(conf HttpConfig) http.Client {
	return http.Client{
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
