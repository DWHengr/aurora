package models

import "gorm.io/gorm"

type AlertRules struct {
	BaseModel

	Name          string `json:"name"`
	AlertObject   string `json:"alert_object"`
	Rules         string `json:"rules"`
	RulesStatus   string `json:"rules_status"`
	Webhook       string `json:"webhook"`
	Persistent    string `json:"persistent"`
	AlertInterval string `json:"alert_interval"`
	StoreInterval string `json:"store_interval"`
	Description   string `json:"description"`
}

type AlertRulesRepo interface {
	GetAll(db *gorm.DB) ([]*AlertRules, error)
}
