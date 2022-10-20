package models

import (
	"gorm.io/gorm"
)

type AlertUsersGroup struct {
	BaseModel

	Name        string `json:"name"`
	Description string `json:"description"`
	UserIds     string `json:"userIds"`
}

type AlertUsersGroupRepo interface {
	Create(db *gorm.DB, user *AlertUsersGroup) error
}
