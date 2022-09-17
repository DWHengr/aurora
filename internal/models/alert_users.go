package models

import (
	"gorm.io/gorm"
)

type AlertUsers struct {
	BaseModel

	Name       string `json:"name"`
	Department string `json:"department"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

type AlertUsersRepo interface {
	Create(db *gorm.DB, user *AlertUsers) error
	Deletes(db *gorm.DB, ids []string) error
	Update(db *gorm.DB, alertUser *AlertUsers) error
}
