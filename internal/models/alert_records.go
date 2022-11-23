package models

import (
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
	"time"
)

type AlertRecords struct {
	BaseModel

	AlertName  string `json:"alertName"`
	RuleName   string `json:"ruleName"`
	RuleId     string `json:"ruleId"`
	Severity   string `json:"severity"`
	Summary    string `json:"summary"`
	Value      string `json:"value"`
	Attribute  string `json:"attribute"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

func (a *AlertRecords) BeforeSave(tx *gorm.DB) error {
	a.UpdateTime = time.Now().Unix()
	return nil
}

func (a *AlertRecords) BeforeCreate(tx *gorm.DB) error {
	a.CreateTime = time.Now().Unix()
	return nil
}

type AlertRecordsRepo interface {
	Create(db *gorm.DB, record *AlertRecords) error
	Delete(db *gorm.DB, alertRecordId string) error
	Deletes(db *gorm.DB, ids []string) error
	Page(db *gorm.DB, page *page.ReqPage) (*page.RespPage, error)
}
