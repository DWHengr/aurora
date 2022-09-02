package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"gorm.io/gorm"
)

type alertMetricsRepo struct{}

func NewAlterMetricsRepo() models.AlertMetricsRepo {
	return &alertMetricsRepo{}
}

func (r *alertMetricsRepo) TableName() string {
	return "alert_metrics"
}

func (r *alertMetricsRepo) GetAll(db *gorm.DB) ([]*models.AlertMetrics, error) {
	entity := make([]*models.AlertMetrics, 0)
	err := db.Table(r.TableName()).
		Find(&entity).
		Error
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *alertMetricsRepo) FindById(db *gorm.DB, id string) (*models.AlertMetrics, error) {
	metric := &models.AlertMetrics{}
	err := db.Table(r.TableName()).Where("id = ?", id).Find(metric).Error
	return metric, err
}
