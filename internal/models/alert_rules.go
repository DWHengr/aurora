package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Rule struct {
	Metric     string `json:"metric"`
	Statistics string `json:"statistics"`
	Operator   string `json:"operator"`
	AlertValue string `json:"alertValue"`
}

type AlertRules struct {
	BaseModel

	Name            string            `json:"name"`
	AlertObject     string            `json:"alertObject"`
	AlertObjectArr  map[string]string `json:"alertObjectArr" gorm:"-"`
	Rules           string            `json:"rules"`
	RulesArr        []Rule            `json:"rulesArr" gorm:"-"`
	RulesStatus     string            `json:"rulesStatus"`
	Severity        string            `json:"severity"`
	Webhook         string            `json:"webhook"`
	AlertSilencesId string            `json:"alertSilencesId"`
	Persistent      string            `json:"persistent"`
	AlertInterval   string            `json:"alertInterval"`
	StoreInterval   string            `json:"storeInterval"`
	Description     string            `json:"description"`
}

func (a *AlertRules) AfterFind(tx *gorm.DB) (err error) {
	rulesResult := make([]Rule, 0)
	err = json.Unmarshal([]byte(a.Rules), &rulesResult)
	a.RulesArr = rulesResult
	alertObjectResult := map[string]string{}
	err = json.Unmarshal([]byte(a.AlertObject), &alertObjectResult)
	a.AlertObjectArr = alertObjectResult
	return
}

type AlertRulesRepo interface {
	GetAll(db *gorm.DB) ([]*AlertRules, error)
	FindById(db *gorm.DB, id string) (*AlertRules, error)
}
