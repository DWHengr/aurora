package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	mysql2 "github.com/DWHengr/aurora/pkg/misc/mysql"
	"gorm.io/gorm"
)

type AlertSilencesService interface {
	GetAllAlertSilences() ([]*models.AlertSilences, error)
	GetAllAlertSilencesMap() (map[string]*models.AlertSilences, error)
}

type alertSilencesService struct {
	db                *gorm.DB
	alertSilencesRepo models.AlertSilencesRepo
}

func NewAlertSilencesService(conf *mysql2.MysqlConfig) (AlertSilencesService, error) {
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
	return tables, err
}

func (s *alertSilencesService) GetAllAlertSilencesMap() (map[string]*models.AlertSilences, error) {
	tables, err := s.alertSilencesRepo.GetAll(s.db)
	if err != nil {
		return nil, err
	}
	tablesMap := make(map[string]*models.AlertSilences)
	for _, table := range tables {
		tablesMap[table.ID] = table
	}
	return tablesMap, err
}
