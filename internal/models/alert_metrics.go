package models

import (
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
	"strings"
	"time"
)

type AlertMetrics struct {
	BaseModel

	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Expression  string   `json:"expression"`
	Unit        string   `json:"unit"`
	Operator    string   `json:"operator"`
	OperatorArr []string `json:"operatorArr" gorm:"-"`
	Description string   `json:"description"`
	CreateTime  int64    `json:"createTime"`
	UpdateTime  int64    `json:"updateTime"`
}

func (a *AlertMetrics) AfterFind(tx *gorm.DB) (err error) {
	a.OperatorArr = strings.Split(a.Operator, ",")
	return
}

func (a *AlertMetrics) BeforeSave(tx *gorm.DB) error {
	if a.OperatorArr != nil {
		a.Operator = strings.Join(a.OperatorArr, ",")
	}
	a.UpdateTime = time.Now().Unix()
	return nil
}

func (a *AlertMetrics) BeforeCreate(tx *gorm.DB) error {
	a.CreateTime = time.Now().Unix()
	return nil
}

type AlertMetricsRepo interface {
	GetAll(db *gorm.DB) ([]*AlertMetrics, error)
	FindById(db *gorm.DB, id string) (*AlertMetrics, error)
	Create(db *gorm.DB, alertMetric *AlertMetrics) error
	Delete(db *gorm.DB, alertMetricId string) error
	Page(db *gorm.DB, page *page.ReqPage) (*page.RespPage, error)
	Update(db *gorm.DB, alertMetric *AlertMetrics) error
	Deletes(db *gorm.DB, ids []string) error
}
