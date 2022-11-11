package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
	"strings"
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
func (r *alertUsersGroupRepo) All(db *gorm.DB) ([]*models.AlertUsersGroup, error) {
	entity := make([]*models.AlertUsersGroup, 1)
	err := db.Table(r.TableName()).
		Find(&entity).
		Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *alertUsersGroupRepo) Update(db *gorm.DB, userGroup *models.AlertUsersGroup) error {
	err := db.Table(r.TableName()).Updates(userGroup).Error
	return err
}

func (r *alertUsersGroupRepo) Deletes(db *gorm.DB, ids []string) error {
	return db.Table(r.TableName()).Where("id in ?", ids).Delete(&models.AlertUsersGroup{}).Error
}

func (r *alertUsersGroupRepo) Page(db *gorm.DB, pageData *page.ReqPage) (*page.RespPage, error) {
	userGroups := make([]*models.AlertUsersGroup, 0)
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
	err := db.Find(&userGroups).Offset(-1).Limit(-1).Count(&count).Error
	alertUsersRepo := NewAlertUsersRepo()
	for _, userGroup := range userGroups {
		users := make([]*models.AlertUsers, 0)
		users, err := alertUsersRepo.GetUserByIds(db, strings.Split(userGroup.UserIds, ","))
		if err == nil {
			userGroup.UserIdsDetail = users
		}
	}
	if err != nil {
		return nil, err
	}
	return &page.RespPage{
		Page:     pageData.Page,
		Size:     pageData.Size,
		Total:    count,
		DataList: userGroups,
	}, nil
}
