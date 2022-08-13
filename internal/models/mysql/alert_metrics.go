package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"gorm.io/gorm"
)

type alterMetricsRepo struct{}

func NewAlterMetricsRepo() models.AlertMetricsRepo {
	return &alterMetricsRepo{}
}

func (r *alterMetricsRepo) TableName() string {
	return "alert_metrics"
}

func (r *alterMetricsRepo) GetAll(db *gorm.DB) ([]*models.AlertMetrics, error) {
	entity := make([]*models.AlertMetrics, 0)
	err := db.Table(r.TableName()).
		Find(&entity).
		Error
	if err != nil {
		return nil, err
	}

	return entity, nil
}
