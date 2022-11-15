package utils

import (
	"github.com/DWHengr/aurora/pkg/config"
	"github.com/DWHengr/aurora/pkg/httpclient"
)

func GetPrometheusMetricValue(metric string) (interface{}, error) {
	allConfig, _ := config.GetAllConfig()
	var result interface{}
	err := httpclient.GET(allConfig.Aurora.PrometheusUrl+"/api/v1/query?query="+metric, &result)
	return result, err
}

func PostPrometheusReload() error {
	allConfig, _ := config.GetAllConfig()
	return httpclient.Request(allConfig.Aurora.PrometheusUrl+"/-/reload", "POST", nil, nil, nil)
}
