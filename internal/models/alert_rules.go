package models

import (
	"encoding/json"
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
)

type AlertRules struct {
	BaseModel

	Name            string                `json:"name"`
	AlertObject     string                `json:"alertObject"`
	AlertObjectArr  map[string]string     `json:"alertObjectArr" gorm:"-"`
	RulesArr        []*RuleMetricRelation `json:"rulesArr" gorm:"-"`
	RulesStatus     string                `json:"rulesStatus"`
	Severity        string                `json:"severity"`
	Webhook         string                `json:"webhook"`
	AlertSilencesId string                `json:"alertSilencesId"`
	Persistent      string                `json:"persistent"`
	AlertInterval   string                `json:"alertInterval"`
	StoreInterval   string                `json:"storeInterval"`
	Description     string                `json:"description"`
}

func (a *AlertRules) AfterFind(tx *gorm.DB) (err error) {
	alertObjectResult := map[string]string{}
	err = json.Unmarshal([]byte(a.AlertObject), &alertObjectResult)
	a.AlertObjectArr = alertObjectResult
	return
}

func (a *AlertRules) BeforeSave(tx *gorm.DB) error {
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
	FindByIds(db *gorm.DB, ids []string) ([]*AlertRules, error)
	Create(db *gorm.DB, alertRule *AlertRules) error
	Delete(db *gorm.DB, alertRuleId string) error
	Update(db *gorm.DB, alertRule *AlertRules) error
	Page(db *gorm.DB, page *page.ReqPage) (*page.RespPage, error)
}
