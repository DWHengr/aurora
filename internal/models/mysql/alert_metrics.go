package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
)

type alertMetricsRepo struct{}

func NewAlertMetricsRepo() models.AlertMetricsRepo {
	return &alertMetricsRepo{}
}

func (r *alertMetricsRepo) TableName() string {
	return AlertMetrics
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

func (r *alertMetricsRepo) Create(db *gorm.DB, alertMetric *models.AlertMetrics) error {
	err := db.Table(r.TableName()).Create(alertMetric).Error
	return err
}

func (r *alertMetricsRepo) Delete(db *gorm.DB, alertMetricId string) error {
	entity := &models.AlertMetrics{
		BaseModel: models.BaseModel{
			ID: alertMetricId,
		},
	}
	err := db.Table(r.TableName()).Delete(entity).Error
	return err
}

func (r *alertMetricsRepo) Page(db *gorm.DB, page *page.ReqPage) (*page.RespPage, error) {
	metrics := make([]*models.AlertMetrics, 0)
	var count int64
	db = db.Table(r.TableName())
	for _, filter := range page.Filters {
		db = db.Where(filter.Column, filter.Value)
	}
	for _, order := range page.Orders {
		db = db.Order(order.Column + " " + order.Direction)
	}
	if page.Page > 0 && page.Size > 0 {
		db = db.Limit(page.Size).Offset((page.Page - 1) * page.Size)
	}
	err := db.Find(&metrics).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}
	return &page.RespPage{
		Page:     page.Page,
		Size:     page.Size,
		Total:    count,
		DataList: metrics,
	}, nil
}

func (r *alertMetricsRepo) Update(db *gorm.DB, alertMetric *models.AlertMetrics) error {
	err := db.Table(r.TableName()).Updates(alertMetric).Error
	return err
}
