package models

import (
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
)

type AlertUsersGroup struct {
	BaseModel

	Name          string        `json:"name"`
	Description   string        `json:"description"`
	UserIds       string        `json:"userIds"`
	UserIdsDetail []*AlertUsers `json:"userIdsDetail" gorm:"-"`
}

type AlertUsersGroupRepo interface {
	Create(db *gorm.DB, userGroup *AlertUsersGroup) error
	Update(db *gorm.DB, userGroup *AlertUsersGroup) error
	Page(db *gorm.DB, page *page.ReqPage) (*page.RespPage, error)
	FindByIds(db *gorm.DB, ids []string) ([]*AlertUsersGroup, error)
	Deletes(db *gorm.DB, ids []string) error
	All(db *gorm.DB) ([]*AlertUsersGroup, error)
}
