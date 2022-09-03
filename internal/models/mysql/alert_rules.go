package mysql

import (
	"github.com/DWHengr/aurora/internal/Page"
	"github.com/DWHengr/aurora/internal/models"
	"gorm.io/gorm"
)

var cache = make(map[string]*models.AlertRules)

type alterRulesRepo struct {
}

func NewAlterRulesRepo() models.AlertRulesRepo {
	return &alterRulesRepo{}
}

func (r *alterRulesRepo) TableName() string {
	return AlertRules
}

func (r *alterRulesRepo) setCache(ruleId string, rule *models.AlertRules) {
	cache[ruleId] = rule
}

func (r *alterRulesRepo) deleteCache(ruleId string) {
	delete(cache, ruleId)
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
	rule, ok := cache[id]
	if ok && rule != nil {
		return rule, nil
	}
	rule = &models.AlertRules{}
	err := db.Table(r.TableName()).Where("id = ?", id).Find(rule).Error
	if err == nil {
		r.setCache(rule.ID, rule)
	}
	return rule, err
}

func (r *alterRulesRepo) Create(db *gorm.DB, alertRule *models.AlertRules) error {
	tx := db.Begin()
	err := tx.Table(r.TableName()).Create(alertRule).Error
	if err != nil {
		return err
	}
	for _, v := range alertRule.RulesArr {
		v.RuleId = alertRule.ID
	}
	err = tx.Table(RuleMetricRelation).Omit("expression").CreateInBatches(alertRule.RulesArr, len(alertRule.RulesArr)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *alterRulesRepo) Delete(db *gorm.DB, alertRuleId string) error {
	entity := &models.AlertRules{
		BaseModel: models.BaseModel{
			ID: alertRuleId,
		},
	}
	err := db.Table(r.TableName()).Delete(entity).Error
	if err == nil {
		r.deleteCache(alertRuleId)
	}
	return err
}

func (r *alterRulesRepo) Update(db *gorm.DB, alertRule *models.AlertRules) error {
	err := db.Table(r.TableName()).Updates(alertRule).Error
	if err != nil {
		return err
	}
	err = db.Table(r.TableName()).Where("id = ?", alertRule.ID).Find(alertRule).Error
	if err == nil {
		r.setCache(alertRule.ID, nil)
	}
	return err
}

func (r *alterRulesRepo) Page(db *gorm.DB, page *Page.ReqPage) (*Page.RespPage, error) {
	rules := make([]*models.AlertRules, 0)
	var count int64
	db = db.Table(r.TableName())
	for _, filter := range page.Filters {
		db = db.Where(filter.Column, filter.Value)
	}
	for _, order := range page.Orders {
		db = db.Order(order.Column + " " + order.Direction)
	}
	if page.Page > 0 && page.Size > 0 {
		db = db.Limit(page.Size).Offset((page.Page - 1) * page.Size)
	}
	err := db.Find(&rules).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}
	return &Page.RespPage{
		Page:     page.Page,
		Size:     page.Size,
		Total:    count,
		DataList: rules,
	}, nil
}
