package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/page"
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

func (r *alertRecordsRepo) Delete(db *gorm.DB, alertRecordId string) error {
	entity := &models.AlertRecords{
		BaseModel: models.BaseModel{
			ID: alertRecordId,
		},
	}
	err := db.Table(r.TableName()).Delete(entity).Error
	return err
}

func (r *alertRecordsRepo) Page(db *gorm.DB, pageData *page.ReqPage) (*page.RespPage, error) {
	records := make([]*models.AlertRecords, 0)
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
	err := db.Find(&records).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}
	return &page.RespPage{
		Page:     pageData.Page,
		Size:     pageData.Size,
		Total:    count,
		DataList: records,
	}, nil
}
