package models

import "gorm.io/gorm"

type RuleMetricRelation struct {
	Id         int    `json:"id"`
	RuleId     string `json:"ruleId"`
	MetricId   string `json:"metricId"`
	Expression string `json:"expression"`
	Statistics string `json:"statistics"`
	Operator   string `json:"operator"`
	AlertValue string `json:"alertValue"`
}

type RuleMetricRelationRepo interface {
	GetRuleMetricByRuleId(db *gorm.DB, ruleId string) ([]*RuleMetricRelation, error)
}