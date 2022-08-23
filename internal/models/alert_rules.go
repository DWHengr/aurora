package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type rule struct {
	Metric     string `json:"metric"`
	Statistics string `json:"statistics"`
	Operator   string `json:"operator"`
	AlertValue string `json:"alertValue"`
}

type AlertRules struct {
	BaseModel

	Name            string `json:"name"`
	AlertObject     string `json:"alertObject"`
	Rules           rule   `json:"rules"`
	RulesStatus     string `json:"rulesStatus"`
	Severity        string `json:"severity"`
	Webhook         string `json:"webhook"`
	AlertSilencesId string `json:"alertSilencesId"`
	Persistent      string `json:"persistent"`
	AlertInterval   string `json:"alertInterval"`
	StoreInterval   string `json:"storeInterval"`
	Description     string `json:"description"`
}

func (r *rule) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := rule{}
	err := json.Unmarshal(bytes, &result)
	*r = result
	return err
}

func (r rule) Value() (driver.Value, error) {
	jsonByte, err := json.Marshal(r)
	jsonStr := string(jsonByte)
	return jsonStr, err
}

type AlertRulesRepo interface {
	GetAll(db *gorm.DB) ([]*AlertRules, error)
	FindById(db *gorm.DB, id string) (*AlertRules, error)
}
