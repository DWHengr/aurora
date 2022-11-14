package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/internal/page"
	"github.com/DWHengr/aurora/internal/service/utils"
	"github.com/DWHengr/aurora/pkg/id"
	"gorm.io/gorm"
	"strings"
)

type AlertUsersGroupService interface {
	Create(userGroup *models.AlertUsersGroup) (*CreateAlertUserGroupResp, error)
	Page(page *page.ReqPage) (*page.RespPage, error)
	Update(userGroup *models.AlertUsersGroup) (*CreateAlertUserGroupResp, error)
	Deletes(ids []string) error
	All() ([]*models.AlertUsersGroup, error)
	GetGroupUser(groupIds []string) ([]*models.AlertUsers, error)
}

type alertUsersGroupService struct {
	db                  *gorm.DB
	alertUsersGroupRepo models.AlertUsersGroupRepo
	alertUserRepo       models.AlertUsersRepo
}

func NewAlertUsersGroupService() (AlertUsersGroupService, error) {
	db := utils.GetMysqlInstance()

	return &alertUsersGroupService{
		db:                  db,
		alertUsersGroupRepo: mysql.NewAlertUsersGroupRepo(),
		alertUserRepo:       mysql.NewAlertUsersRepo(),
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

func (s *alertUsersGroupService) GetGroupUser(groupIds []string) ([]*models.AlertUsers, error) {
	groups, err := s.alertUsersGroupRepo.FindByIds(s.db, groupIds)
	if err != nil {
		return nil, err
	}
	var userIds []string
	for _, group := range groups {
		userIds = append(userIds, strings.Split(group.UserIds, ",")...)
	}
	users, _ := s.alertUserRepo.GetUserByIds(s.db, userIds)
	if err != nil {
		return nil, err
	}
	return users, nil
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

func (s *alertUsersGroupService) All() ([]*models.AlertUsersGroup, error) {
	return s.alertUsersGroupRepo.All(s.db)
}
