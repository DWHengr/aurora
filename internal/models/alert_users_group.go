package models

import (
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
)

type AlertUsersGroup struct {
	BaseModel

	Name        string `json:"name"`
	Description string `json:"description"`
	UserIds     string `json:"userIds"`
}

type AlertUsersGroupRepo interface {
	Create(db *gorm.DB, userGroup *AlertUsersGroup) error
	Update(db *gorm.DB, userGroup *AlertUsersGroup) error
	Page(db *gorm.DB, page *page.ReqPage) (*page.RespPage, error)
}
