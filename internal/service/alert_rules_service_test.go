package service

import (
	"fmt"
	"github.com/DWHengr/aurora/pkg/misc/mysql"
	"testing"
)

func Test_alertRulesService_GetAllAlertMetrics(t *testing.T) {
	mysqlConfig := &mysql.MysqlConfig{
		Host:     "127.0.0.1",
		DB:       "aurora",
		User:     "root",
		Password: "123456",
		Log:      true,
	}
	service, err := NewAlertRulesService(mysqlConfig)
	if err != nil {
		t.Fatal(err)
	}
	alertMetrics, err := service.GetAllAlertRules()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*alertMetrics[0])
}
