package service

import (
	"fmt"
	"github.com/DWHengr/aurora/pkg/misc/mysql"
	"testing"
)

func Test_alertMetricsService_GetAllAlertMetrics(t *testing.T) {
	mysqlConfig := &mysql.MysqlConfig{
		Host:     "127.0.0.1",
		DB:       "aurora",
		User:     "root",
		Password: "123456",
		Log:      true,
	}
	service, err := NewAlertMetricsService(mysqlConfig)
	if err != nil {
		t.Fatal(err)
	}
	alertMetrics, err := service.GetAllAlertMetrics()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*alertMetrics[0])
}
