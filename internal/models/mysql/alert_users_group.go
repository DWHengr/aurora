package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"gorm.io/gorm"
)

type alertUsersGroupRepo struct{}

func NewAlertUsersGroupRepo() models.AlertUsersGroupRepo {
	return &alertUsersGroupRepo{}
}

func (r *alertUsersGroupRepo) TableName() string {
	return AlertUsersGroup
}

func (r *alertUsersGroupRepo) Create(db *gorm.DB, userGroup *models.AlertUsersGroup) error {
	return db.Table(r.TableName()).Create(userGroup).Error
}
