package service

import (
	"errors"
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/pkg/id"
	"gorm.io/gorm"
)

type AlertMetricsService interface {
	GetAllAlertMetrics() ([]*models.AlertMetrics, error)
	Create(rule *models.AlertMetrics) (*CreateAlertMetricResp, error)
	Delete(metricId string) error
}

type alertMetricsService struct {
	db                     *gorm.DB
	alertMetricsRepo       models.AlertMetricsRepo
	ruleMetricRelationRepo models.RuleMetricRelationRepo
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

func (s *alertMetricsService) Create(metric *models.AlertMetrics) (*CreateAlertMetricResp, error) {
	metric.ID = "mtc-" + id.ShortID(8)
	err := s.alertMetricsRepo.Create(s.db, metric)
	if err != nil {
		return nil, err
	}
	return &CreateAlertMetricResp{
		ID: metric.ID,
	}, nil
}

func (s *alertMetricsService) Delete(metricId string) error {
	count, err := s.ruleMetricRelationRepo.GetCountByMetricID(s.db, metricId)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("this metric has alert in use")
	}
	err = s.alertMetricsRepo.Delete(s.db, metricId)
	if err != nil {
		return err
	}
	return err

}
