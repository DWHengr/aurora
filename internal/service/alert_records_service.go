package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/pkg/id"
	"gorm.io/gorm"
	"time"
)

type AlertRecordsService interface {
	CreateRecord(records *models.AlertRecords) error
}

type alertRecordsService struct {
	db               *gorm.DB
	alertRecordsRepo models.AlertRecordsRepo
}

func NewAlertRecordsService() (AlertRecordsService, error) {
	db := GetMysqlInstance()

	return &alertRecordsService{
		db:               db,
		alertRecordsRepo: mysql.NewAlertRecordsRepo(),
	}, nil
}
func (s *alertRecordsService) CreateRecord(records *models.AlertRecords) error {
	records.CreateTime = time.Unix(time.Now().Unix(), 0)
	records.ID = "rc-" + id.ShortID(8)
	err := s.alertRecordsRepo.Create(s.db, records)
	return err
}
