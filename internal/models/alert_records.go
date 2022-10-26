package models

import (
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
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
}

type AlertRecordsRepo interface {
	Create(db *gorm.DB, record *AlertRecords) error
	Delete(db *gorm.DB, alertRecordId string) error
	Deletes(db *gorm.DB, ids []string) error
	Page(db *gorm.DB, page *page.ReqPage) (*page.RespPage, error)
}
