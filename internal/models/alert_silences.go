package models

import (
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
}
