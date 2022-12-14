package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/internal/page"
	"github.com/DWHengr/aurora/internal/service/utils"
	"github.com/DWHengr/aurora/pkg/id"
	"gorm.io/gorm"
)

type AlertRecordsService interface {
	CreateRecord(records *models.AlertRecords) error
	Delete(metricId string) error
	Page(page *page.ReqPage) (*page.RespPage, error)
	Deletes(ids []string) error
}

type alertRecordsService struct {
	db               *gorm.DB
	alertRecordsRepo models.AlertRecordsRepo
}

func NewAlertRecordsService() (AlertRecordsService, error) {
	db := utils.GetMysqlInstance()

	return &alertRecordsService{
		db:               db,
		alertRecordsRepo: mysql.NewAlertRecordsRepo(),
	}, nil
}
func (s *alertRecordsService) CreateRecord(records *models.AlertRecords) error {
	records.ID = "rc-" + id.ShortID(8)
	err := s.alertRecordsRepo.Create(s.db, records)
	return err
}

func (s *alertRecordsService) Delete(metricId string) error {
	err := s.alertRecordsRepo.Delete(s.db, metricId)
	if err != nil {
		return err
	}
	return err

}

func (s *alertRecordsService) Deletes(ids []string) error {
	err := s.alertRecordsRepo.Deletes(s.db, ids)
	return err
}

func (s *alertRecordsService) Page(page *page.ReqPage) (*page.RespPage, error) {
	return s.alertRecordsRepo.Page(s.db, page)
}
