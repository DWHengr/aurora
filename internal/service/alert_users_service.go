package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/pkg/id"
	"gorm.io/gorm"
)

type AlertUsersService interface {
	Create(user *models.AlertUsers) (*CreateAlertUserResp, error)
	Deletes(ids []string) error
}

type alertUsersService struct {
	db             *gorm.DB
	alertUsersRepo models.AlertUsersRepo
}

func NewAlertUsersService() (AlertUsersService, error) {
	db := GetMysqlInstance()

	return &alertUsersService{
		db:             db,
		alertUsersRepo: mysql.NewAlertUsersRepo(),
	}, nil
}

type CreateAlertUserResp struct {
	ID string `json:"id"`
}

func (s *alertUsersService) Create(user *models.AlertUsers) (*CreateAlertUserResp, error) {
	user.ID = "usr-" + id.ShortID(8)
	err := s.alertUsersRepo.Create(s.db, user)
	if err != nil {
		return nil, err
	}
	return &CreateAlertUserResp{
		ID: user.ID,
	}, nil
}

func (s *alertUsersService) Deletes(ids []string) error {
	err := s.alertUsersRepo.Deletes(s.db, ids)
	return err
}
