package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/pkg/id"
	"gorm.io/gorm"
)

type AlertUsersGroupService interface {
	Create(userGroup *models.AlertUsersGroup) (*CreateAlertUserGroupResp, error)
}

type alertUsersGroupService struct {
	db                  *gorm.DB
	alertUsersGroupRepo models.AlertUsersGroupRepo
}

func NewAlertUsersGroupService() (AlertUsersGroupService, error) {
	db := GetMysqlInstance()

	return &alertUsersGroupService{
		db:                  db,
		alertUsersGroupRepo: mysql.NewAlertUsersGroupRepo(),
	}, nil
}

type CreateAlertUserGroupResp struct {
	ID string `json:"id"`
}

func (s *alertUsersGroupService) Create(userGroup *models.AlertUsersGroup) (*CreateAlertUserGroupResp, error) {
	userGroup.ID = "usg-" + id.ShortID(8)
	err := s.alertUsersGroupRepo.Create(s.db, userGroup)
	if err != nil {
		return nil, err
	}
	return &CreateAlertUserGroupResp{
		ID: userGroup.ID,
	}, nil
}
