package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"gorm.io/gorm"
)

type alertUsersRepo struct{}

func NewAlertUsersRepo() models.AlertUsersRepo {
	return &alertUsersRepo{}
}

func (r *alertUsersRepo) TableName() string {
	return AlertUsers
}

func (r *alertUsersRepo) Create(db *gorm.DB, user *models.AlertUsers) error {
	return db.Table(r.TableName()).Create(user).Error
}

func (r *alertUsersRepo) Deletes(db *gorm.DB, ids []string) error {
	return db.Table(r.TableName()).Where("id in ?", ids).Delete(&models.AlertUsers{}).Error
}

func (r *alertUsersRepo) Update(db *gorm.DB, alertUser *models.AlertUsers) error {
	err := db.Table(r.TableName()).Updates(alertUser).Error
	return err
}
