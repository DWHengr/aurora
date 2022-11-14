package utils

import (
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/DWHengr/aurora/pkg/misc/mysql"
	"gorm.io/gorm"
)

var (
	mysqlDBInst *gorm.DB
)

func GetMysqlInstance() *gorm.DB {
	if mysqlDBInst == nil {
		panic("mysql instance is nil")
	}
	return mysqlDBInst
}

func NewMysqlInstanceByConn(conf *mysql.MysqlConfig) (*gorm.DB, error) {
	if mysqlDBInst == nil {
		db, err := mysql.New(conf, logger.Logger)
		if err != nil {
			return nil, err
		}
		mysqlDBInst = db
	}
	return mysqlDBInst, nil
}
