package service

import (
	"fmt"
	"github.com/DWHengr/aurora/pkg/config"
	"github.com/DWHengr/aurora/pkg/misc/mysql"
	"testing"
)

func Test_alertSilencesService_GetAllAlertSilences(t *testing.T) {
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
	service, err := NewAlertSilencesService(config)
	if err != nil {
		t.Fatal(err)
	}
	alertSilences, err := service.GetAllAlertSilences()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*alertSilences[0])
}
