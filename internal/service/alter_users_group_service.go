package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/internal/page"
	"github.com/DWHengr/aurora/pkg/id"
	"gorm.io/gorm"
)

type AlertUsersGroupService interface {
	Create(userGroup *models.AlertUsersGroup) (*CreateAlertUserGroupResp, error)
	Page(page *page.ReqPage) (*page.RespPage, error)
	Update(userGroup *models.AlertUsersGroup) (*CreateAlertUserGroupResp, error)
	Deletes(ids []string) error
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

func (s *alertUsersGroupService) Page(page *page.ReqPage) (*page.RespPage, error) {
	return s.alertUsersGroupRepo.Page(s.db, page)
}

func (s *alertUsersGroupService) Update(userGroup *models.AlertUsersGroup) (*CreateAlertUserGroupResp, error) {
	err := s.alertUsersGroupRepo.Update(s.db, userGroup)
	return &CreateAlertUserGroupResp{
		ID: userGroup.ID,
	}, err
}

func (s *alertUsersGroupService) Deletes(ids []string) error {
	err := s.alertUsersGroupRepo.Deletes(s.db, ids)
	return err
}
