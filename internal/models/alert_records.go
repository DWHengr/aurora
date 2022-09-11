package models

import (
	"gorm.io/gorm"
	"time"
)

type AlertRecords struct {
	BaseModel

	AlertName  string    `json:"alertName"`
	RuleName   string    `json:"ruleName"`
	RuleId     string    `json:"ruleId"`
	Severity   string    `json:"severity"`
	Summary    string    `json:"summary"`
	Value      string    `json:"value"`
	Attribute  string    `json:"attribute"`
	CreateTime time.Time `json:"createTime"`
}

type AlertRecordsRepo interface {
	Create(db *gorm.DB, record *AlertRecords) error
	Delete(db *gorm.DB, alertRecordId string) error
}
