package service

import (
	"fmt"
	"github.com/DWHengr/aurora/pkg/config"
	"github.com/DWHengr/aurora/pkg/misc/mysql"
	"testing"
)

func Test_alertRulesService_GetAllAlertMetrics(t *testing.T) {
	mysqlConfig := mysql.MysqlConfig{
		Host:     "127.0.0.1",
		DB:       "aurora",
		User:     "root",
		Password: "123456",
		Log:      true,
	}
	config := &config.Config{
		Mysql: mysqlConfig,
	}
	service, err := NewAlertRulesService(config)
	if err != nil {
		t.Fatal(err)
	}
	alertMetrics, err := service.GetAllAlertRules()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*alertMetrics[0])
}
