package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/pkg/config"
	"gorm.io/gorm"
)

type AlertSilencesService interface {
	GetAllAlertSilences() ([]*models.AlertSilences, error)
}

type alertSilencesService struct {
	db                *gorm.DB
	alertSilencesRepo models.AlertSilencesRepo
}

func NewAlertSilencesService(conf *config.Config) (AlertSilencesService, error) {
	db, err := CreateMysqlConn(conf)
	if err != nil {
		return nil, err
	}

	return &alertSilencesService{
		db:                db,
		alertSilencesRepo: mysql.NewAlterSilencesRepo(),
	}, nil
}
func (s *alertSilencesService) GetAllAlertSilences() ([]*models.AlertSilences, error) {
	tables, err := s.alertSilencesRepo.GetAll(s.db)
	if err != nil {
		return nil, err
	}
	// TODO
	return tables, err
}
