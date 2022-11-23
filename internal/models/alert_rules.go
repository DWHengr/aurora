package models

import (
	"encoding/json"
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
	"time"
)

type AlertObjectArr struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AlertRules struct {
	BaseModel

	Name        string `json:"name"`
	AlertObject string `json:"alertObject"`
	//AlertObjectArr  map[string]string     `json:"alertObjectArr" gorm:"-"`
	AlertObjectArr  []*AlertObjectArr     `json:"alertObjectArr" gorm:"-"`
	RulesArr        []*RuleMetricRelation `json:"rulesArr" gorm:"-"`
	RulesStatus     int                   `json:"rulesStatus"`
	Severity        string                `json:"severity"`
	Webhook         string                `json:"webhook"`
	AlertSilencesId string                `json:"alertSilencesId"`
	Persistent      string                `json:"persistent"`
	AlertInterval   string                `json:"alertInterval"`
	StoreInterval   string                `json:"storeInterval"`
	UserGroupIds    string                `json:"userGroupIds"`
	UserGroupIdsArr []string              `json:"userGroupIdsArr" gorm:"-"`
	Description     string                `json:"description"`
	CreateTime      int64                 `json:"createTime"`
	UpdateTime      int64                 `json:"updateTime"`
}

func (a *AlertRules) AfterFind(tx *gorm.DB) (err error) {
	var alertObjectResult []*AlertObjectArr
	err = json.Unmarshal([]byte(a.AlertObject), &alertObjectResult)
	a.AlertObjectArr = alertObjectResult
	var userGroupIdsArr []string
	err = json.Unmarshal([]byte(a.UserGroupIds), &userGroupIdsArr)
	a.UserGroupIdsArr = userGroupIdsArr
	return
}

func (a *AlertRules) BeforeSave(tx *gorm.DB) error {
	a.UpdateTime = time.Now().Unix()
	if a.AlertObjectArr != nil {
		alertObjectResult, err := json.Marshal(a.AlertObjectArr)
		if err != nil {
			return err
		}
		a.AlertObject = string(alertObjectResult)
	}
	if a.UserGroupIdsArr != nil {
		userGroupIdsResult, err := json.Marshal(a.UserGroupIdsArr)
		if err != nil {
			return err
		}
		a.UserGroupIds = string(userGroupIdsResult)
	}
	return nil
}

func (a *AlertRules) BeforeCreate(tx *gorm.DB) error {
	a.CreateTime = time.Now().Unix()
	return nil
}

type AlertRulesRepo interface {
	GetAll(db *gorm.DB) ([]*AlertRules, error)
	FindById(db *gorm.DB, id string) (*AlertRules, error)
	FindEnableByIds(db *gorm.DB, ids []string) ([]*AlertRules, error)
	Create(db *gorm.DB, alertRule *AlertRules) error
	Delete(db *gorm.DB, alertRuleId string) error
	Deletes(db *gorm.DB, ids []string) error
	Update(db *gorm.DB, alertRule *AlertRules) error
	UpdateStatus(db *gorm.DB, alertRule *AlertRules) error
	Page(db *gorm.DB, page *page.ReqPage) (*page.RespPage, error)
}
