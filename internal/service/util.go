package service

import (
	"github.com/DWHengr/aurora/pkg/config"
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/DWHengr/aurora/pkg/misc/mysql"
	"gorm.io/gorm"
)

var (
	mysqlDBInst *gorm.DB
)

func CreateMysqlConn(conf *config.Config) (*gorm.DB, error) {
	if mysqlDBInst == nil {
		db, err := mysql.New(conf.Mysql, logger.Logger)
		if err != nil {
			return nil, err
		}
		mysqlDBInst = db
	}
	return mysqlDBInst, nil
}
