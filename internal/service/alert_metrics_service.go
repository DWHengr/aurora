package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/pkg/id"
	"gorm.io/gorm"
)

type AlertMetricsService interface {
	GetAllAlertMetrics() ([]*models.AlertMetrics, error)
	Create(rule *models.AlertMetrics) (*CreateAlertMetricResp, error)
}

type alertMetricsService struct {
	db               *gorm.DB
	alertMetricsRepo models.AlertMetricsRepo
}

func NewAlertMetricsService() (AlertMetricsService, error) {
	db := GetMysqlInstance()

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

type CreateAlertMetricResp struct {
	ID string `json:"id"`
}

func (s *alertMetricsService) Create(rule *models.AlertMetrics) (*CreateAlertMetricResp, error) {
	rule.ID = "mtc-" + id.ShortID(8)
	err := s.alertMetricsRepo.Create(s.db, rule)
	if err != nil {
		return nil, err
	}
	return &CreateAlertMetricResp{
		ID: rule.ID,
	}, nil
}
