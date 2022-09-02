package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"gorm.io/gorm"
)

type alterSilencesRepo struct{}

func NewAlterSilencesRepo() models.AlertSilencesRepo {
	return &alterSilencesRepo{}
}

func (r *alterSilencesRepo) TableName() string {
	return AlertSilences
}

func (r *alterSilencesRepo) GetAll(db *gorm.DB) ([]*models.AlertSilences, error) {
	entity := make([]*models.AlertSilences, 0)
	err := db.Table(r.TableName()).
		Find(&entity).
		Error
	if err != nil {
		return nil, err
	}

	return entity, nil
}
