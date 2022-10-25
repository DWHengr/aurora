package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/page"
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

func (r *alertUsersRepo) All(db *gorm.DB) ([]*models.AlertUsers, error) {
	entity := make([]*models.AlertUsers, 0)
	err := db.Table(r.TableName()).
		Find(&entity).
		Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *alertUsersRepo) Page(db *gorm.DB, pageData *page.ReqPage) (*page.RespPage, error) {
	users := make([]*models.AlertUsers, 0)
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
	err := db.Find(&users).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}
	return &page.RespPage{
		Page:     pageData.Page,
		Size:     pageData.Size,
		Total:    count,
		DataList: users,
	}, nil
}

func (r *alertUsersRepo) GetUserByIds(db *gorm.DB, ids []string) ([]*models.AlertUsers, error) {
	users := make([]*models.AlertUsers, 0)
	err := db.Raw("select * from "+r.TableName()+" where id in ?", ids).Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
