package utils

import (
	"github.com/DWHengr/aurora/pkg/config"
	"github.com/DWHengr/aurora/pkg/httpclient"
)

func PostPrometheusReload() error {
	allConfig, _ := config.GetAllConfig()
	return httpclient.Request(allConfig.Aurora.PrometheusUrl+"/-/reload", "POST", nil, nil, nil)
}
