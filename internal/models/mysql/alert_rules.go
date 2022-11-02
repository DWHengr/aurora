package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/page"
	"gorm.io/gorm"
)

var cache = make(map[string]*models.AlertRules)

type alterRulesRepo struct {
}

func NewAlertRulesRepo() models.AlertRulesRepo {
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

func (r *alterRulesRepo) FindByIds(db *gorm.DB, ids []string) ([]*models.AlertRules, error) {
	rules := make([]*models.AlertRules, 0)
	err := db.Table(r.TableName()).Where("id in ?", ids).Find(rules).Error
	return rules, err
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
	if err != nil {
		return err
	}
	r.deleteCache(alertRuleId)
	return nil
}

func (r *alterRulesRepo) Deletes(db *gorm.DB, ids []string) error {
	err := db.Table(r.TableName()).Where("id in ?", ids).Delete(&models.AlertRules{}).Error
	if err != nil {
		return err
	}
	for _, id := range ids {
		r.deleteCache(id)
	}
	return nil
}

func (r *alterRulesRepo) Update(db *gorm.DB, alertRule *models.AlertRules) error {
	err := db.Table(r.TableName()).Updates(alertRule).Error
	if err != nil {
		return err
	}
	r.setCache(alertRule.ID, nil)
	err = db.Table(r.TableName()).Where("id = ?", alertRule.ID).Find(alertRule).Error
	return err
}

func (r *alterRulesRepo) Page(db *gorm.DB, pageData *page.ReqPage) (*page.RespPage, error) {
	rules := make([]*models.AlertRules, 0)
	var count int64
	db = db.Table(r.TableName())
	for _, filter := range pageData.Filters {
		if filter.Operator == "like" {
			filter.Value = "%" + filter.Value + "%"
		}
		if filter.Operator == "" {
			filter.Operator = "="
		}
		db = db.Where(filter.Column+" "+filter.Operator+" ?", filter.Value)
	}
	for _, order := range pageData.Orders {
		db = db.Order(order.Column + " " + order.Direction)
	}
	if pageData.Page > 0 && pageData.Size > 0 {
		db = db.Limit(pageData.Size).Offset((pageData.Page - 1) * pageData.Size)
	}
	err := db.Find(&rules).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, err
	}
	return &page.RespPage{
		Page:     pageData.Page,
		Size:     pageData.Size,
		Total:    count,
		DataList: rules,
	}, nil
}
