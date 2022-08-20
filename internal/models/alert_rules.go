package models

import "gorm.io/gorm"

type AlertRules struct {
	BaseModel

	Name            string `json:"name"`
	AlertObject     string `json:"alertObject"`
	Rules           string `json:"rules"`
	RulesStatus     string `json:"rulesStatus"`
	Severity        string `json:"severity"`
	Webhook         string `json:"webhook"`
	AlertSilencesId string `json:"alertSilencesId"`
	Persistent      string `json:"persistent"`
	AlertInterval   string `json:"alertInterval"`
	StoreInterval   string `json:"storeInterval"`
	Description     string `json:"description"`
}

type AlertRulesRepo interface {
	GetAll(db *gorm.DB) ([]*AlertRules, error)
	FindById(db *gorm.DB, id string) (*AlertRules, error)
}
