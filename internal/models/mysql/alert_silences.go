package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
)

type alterSilencesRepo struct{}

func NewAlertSilencesRepo() models.AlertSilencesRepo {
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

func (r *alterSilencesRepo) Create(db *gorm.DB, silence *models.AlertSilences) error {
	return db.Table(r.TableName()).Create(silence).Error
}

func (r *alterSilencesRepo) Deletes(db *gorm.DB, ids []string) error {
	return db.Table(r.TableName()).Where("id in ?", ids).Delete(&models.AlertSilences{}).Error
}

func (r *alterSilencesRepo) Update(db *gorm.DB, silence *models.AlertSilences) error {
	err := db.Table(r.TableName()).Updates(silence).Error
	return err
}

func (r *alterSilencesRepo) Page(db *gorm.DB, pageData *page.ReqPage) (*page.RespPage, error) {
	silences := make([]*models.AlertSilences, 0)
	var count int64
	db = db.Table(r.TableName())
	for _, filter := range pageData.Filters {
		if filter.Operator == "like" {
			filter.Value = "%" + filter.Value + "%"
		}
		if filter.Operator == "" {
			filter.Operator = "="
		}
		db = db.Where(filter.Column+" "+filter.Operator+" ?", filter.Value)
	}
	for _, order := range pageData.Orders {
		db = db.Order(order.Column + " " + order.Direction)
	}
	if pageData.Page > 0 && pageData.Size > 0 {
		db = db.Limit(pageData.Size).Offset((pageData.Page - 1) * pageData.Size)
	}
	err := db.Find(&silences).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}
	return &page.RespPage{
		Page:     pageData.Page,
		Size:     pageData.Size,
		Total:    count,
		DataList: silences,
	}, nil
}
