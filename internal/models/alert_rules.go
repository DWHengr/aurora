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

func (a *AlertRules) BeforeSave(tx *gorm.DB) error {
	if a.RulesArr != nil {
		rulesResult, err := json.Marshal(a.RulesArr)
		if err != nil {
			return err
		}
		a.Rules = string(rulesResult)
	}
	if a.AlertObjectArr != nil {
		alertObjectResult, err := json.Marshal(a.AlertObjectArr)
		if err != nil {
			return err
		}
		a.AlertObject = string(alertObjectResult)
	}
	return nil
}

type AlertRulesRepo interface {
	GetAll(db *gorm.DB) ([]*AlertRules, error)
	FindById(db *gorm.DB, id string) (*AlertRules, error)
	Create(db *gorm.DB, alertRule *AlertRules) error
	Delete(db *gorm.DB, alertRuleId string) error
	Update(db *gorm.DB, alertRule *AlertRules) error
}
