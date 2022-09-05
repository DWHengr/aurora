package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"gorm.io/gorm"
)

type alertRecordsRepo struct{}

func NewAlertRecordsRepo() models.AlertRecordsRepo {
	return &alertRecordsRepo{}
}

func (r *alertRecordsRepo) TableName() string {
	return AlertRecords
}

func (r *alertRecordsRepo) Create(db *gorm.DB, record *models.AlertRecords) error {
	return db.Table(r.TableName()).Create(record).Error
}
