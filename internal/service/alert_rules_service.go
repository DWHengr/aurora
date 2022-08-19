package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	mysql2 "github.com/DWHengr/aurora/pkg/misc/mysql"
	"gorm.io/gorm"
)

type AlertRulesService interface {
	GetAllAlertRules() ([]*models.AlertRules, error)
}

type alertRulesService struct {
	db             *gorm.DB
	alertRulesRepo models.AlertRulesRepo
}

func NewAlertRulesService(conf *mysql2.MysqlConfig) (AlertRulesService, error) {
	db, err := CreateMysqlConn(conf)
	if err != nil {
		return nil, err
	}

	return &alertRulesService{
		db:             db,
		alertRulesRepo: mysql.NewAlterRulesRepo(),
	}, nil
}
func (s *alertRulesService) GetAllAlertRules() ([]*models.AlertRules, error) {
	tables, err := s.alertRulesRepo.GetAll(s.db)
	if err != nil {
		return nil, err
	}
	// TODO
	return tables, err
}
