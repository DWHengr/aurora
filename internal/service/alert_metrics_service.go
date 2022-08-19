package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	mysql2 "github.com/DWHengr/aurora/pkg/misc/mysql"
	"gorm.io/gorm"
)

type AlertMetricsService interface {
	GetAllAlertMetrics() ([]*models.AlertMetrics, error)
}

type alertMetricsService struct {
	db               *gorm.DB
	alertMetricsRepo models.AlertMetricsRepo
}

func NewAlertMetricsService(conf *mysql2.MysqlConfig) (AlertMetricsService, error) {
	db, err := CreateMysqlConn(conf)
	if err != nil {
		return nil, err
	}

	return &alertMetricsService{
		db:               db,
		alertMetricsRepo: mysql.NewAlterMetricsRepo(),
	}, nil
}
func (s *alertMetricsService) GetAllAlertMetrics() ([]*models.AlertMetrics, error) {
	tables, err := s.alertMetricsRepo.GetAll(s.db)
	if err != nil {
		return nil, err
	}
	// TODO
	return tables, err
}
