package models

import (
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
	"time"
)

type AlertSilences struct {
	BaseModel

	Name        string `json:"name"`
	Type        string `json:"type"`
	StartTime   int64  `json:"startTime"`
	EndTime     int64  `json:"endTime"`
	Description string `json:"description"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
}

func (a *AlertSilences) BeforeSave(tx *gorm.DB) error {
	a.UpdateTime = time.Now().Unix()
	return nil
}

func (a *AlertSilences) BeforeCreate(tx *gorm.DB) error {
	a.CreateTime = time.Now().Unix()
	return nil
}

type AlertSilencesRepo interface {
	GetAll(db *gorm.DB) ([]*AlertSilences, error)
	Create(db *gorm.DB, silence *AlertSilences) error
	Deletes(db *gorm.DB, ids []string) error
	Update(db *gorm.DB, silence *AlertSilences) error
	Page(db *gorm.DB, page *page.ReqPage) (*page.RespPage, error)
}
