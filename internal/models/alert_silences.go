package models

import (
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
	"time"
)

type AlertSilences struct {
	BaseModel

	Name        string    `json:"name"`
	Type        string    `json:"type"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Description string    `json:"description"`
}

type AlertSilencesRepo interface {
	GetAll(db *gorm.DB) ([]*AlertSilences, error)
	Create(db *gorm.DB, silence *AlertSilences) error
	Deletes(db *gorm.DB, ids []string) error
	Update(db *gorm.DB, silence *AlertSilences) error
	Page(db *gorm.DB, page *page.ReqPage) (*page.RespPage, error)
}
