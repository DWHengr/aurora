package models

import "gorm.io/gorm"

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
}
