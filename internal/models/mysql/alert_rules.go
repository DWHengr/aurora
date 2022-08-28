package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"gorm.io/gorm"
)

type alterRulesRepo struct{}

func NewAlterRulesRepo() models.AlertRulesRepo {
	return &alterRulesRepo{}
}

func (r *alterRulesRepo) TableName() string {
	return "alert_rules"
}

func (r *alterRulesRepo) GetAll(db *gorm.DB) ([]*models.AlertRules, error) {
	entity := make([]*models.AlertRules, 0)
	err := db.Table(r.TableName()).
		Find(&entity).
		Error
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *alterRulesRepo) FindById(db *gorm.DB, id string) (*models.AlertRules, error) {
	rule := &models.AlertRules{}
	err := db.Table(r.TableName()).Where("id = ?", id).Find(rule).Error
	return rule, err
}

func (r *alterRulesRepo) Create(db *gorm.DB, alertRule *models.AlertRules) error {
	err := db.Table(r.TableName()).Create(alertRule).Error
	return err
}

func (r *alterRulesRepo) Delete(db *gorm.DB, alertRuleId string) error {
	entity := &models.AlertRules{
		BaseModel: models.BaseModel{
			ID: alertRuleId,
		},
	}
	err := db.Table(r.TableName()).Delete(entity).Error
	return err
}
