package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/pkg/id"
	"gorm.io/gorm"
)

type AlertSilencesService interface {
	GetAllAlertSilences() ([]*models.AlertSilences, error)
	GetAllAlertSilencesMap() (map[string]*models.AlertSilences, error)
	Create(silence *models.AlertSilences) (*CreateAlertSilenceResp, error)
}

type alertSilencesService struct {
	db                *gorm.DB
	alertSilencesRepo models.AlertSilencesRepo
}

func NewAlertSilencesService() (AlertSilencesService, error) {
	db := GetMysqlInstance()

	return &alertSilencesService{
		db:                db,
		alertSilencesRepo: mysql.NewAlertSilencesRepo(),
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

type CreateAlertSilenceResp struct {
	ID string `json:"id"`
}

func (s *alertSilencesService) Create(silence *models.AlertSilences) (*CreateAlertSilenceResp, error) {
	silence.ID = "usr-" + id.ShortID(8)
	err := s.alertSilencesRepo.Create(s.db, silence)
	if err != nil {
		return nil, err
	}
	return &CreateAlertSilenceResp{
		ID: silence.ID,
	}, nil
}
