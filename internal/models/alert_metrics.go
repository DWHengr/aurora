package models

import (
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
)

type AlertMetrics struct {
	BaseModel

	Name        string `json:"name"`
	Type        string `json:"type"`
	Expression  string `json:"expression"`
	Unit        string `json:"unit"`
	Operator    string `json:"operator"`
	Description string `json:"description"`
}

type AlertMetricsRepo interface {
	GetAll(db *gorm.DB) ([]*AlertMetrics, error)
	FindById(db *gorm.DB, id string) (*AlertMetrics, error)
	Create(db *gorm.DB, alertMetric *AlertMetrics) error
	Delete(db *gorm.DB, alertMetricId string) error
	Page(db *gorm.DB, page *page.ReqPage) (*page.RespPage, error)
	Update(db *gorm.DB, alertMetric *AlertMetrics) error
	Deletes(db *gorm.DB, ids []string) error
}
