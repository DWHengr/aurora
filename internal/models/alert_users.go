package models

import (
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
	"time"
)

type AlertUsers struct {
	BaseModel

	Name       string `json:"name"`
	Department string `json:"department"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

func (a *AlertUsers) BeforeSave(tx *gorm.DB) error {
	a.UpdateTime = time.Now().Unix()
	return nil
}

func (a *AlertUsers) BeforeCreate(tx *gorm.DB) error {
	a.CreateTime = time.Now().Unix()
	return nil
}

type AlertUsersRepo interface {
	Create(db *gorm.DB, user *AlertUsers) error
	Deletes(db *gorm.DB, ids []string) error
	Update(db *gorm.DB, alertUser *AlertUsers) error
	Page(db *gorm.DB, page *page.ReqPage) (*page.RespPage, error)
	All(db *gorm.DB) ([]*AlertUsers, error)
	GetUserByIds(db *gorm.DB, ids []string) ([]*AlertUsers, error)
}
