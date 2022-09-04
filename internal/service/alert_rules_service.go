package service

import (
	"github.com/DWHengr/aurora/internal/Page"
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/pkg/httpclient"
	"github.com/DWHengr/aurora/pkg/id"
	"gorm.io/gorm"
)

type AlertRulesService interface {
	GetAllAlertRules() ([]*models.AlertRules, error)
	FindById(id string) (*models.AlertRules, error)
	Create(rule *models.AlertRules) (*CreateAlertRuleResp, error)
	Update(rule *models.AlertRules) (*CreateAlertRuleResp, error)
	Delete(ruleId string) error
	Page(page *Page.ReqPage) (*Page.RespPage, error)
}

type alertRulesService struct {
	db                     *gorm.DB
	alertRulesRepo         models.AlertRulesRepo
	ruleMetricRelationRepo models.RuleMetricRelationRepo
}

func NewAlertRulesService() (AlertRulesService, error) {
	db := GetMysqlInstance()

	return &alertRulesService{
		db:                     db,
		alertRulesRepo:         mysql.NewAlterRulesRepo(),
		ruleMetricRelationRepo: mysql.NewRuleMetricRelationRepo(),
	}, nil
}
func (s *alertRulesService) GetAllAlertRules() ([]*models.AlertRules, error) {
	tables, err := s.alertRulesRepo.GetAll(s.db)
	if err != nil {
		return nil, err
	}
	return tables, err
}

func (s *alertRulesService) FindById(id string) (*models.AlertRules, error) {
	return s.alertRulesRepo.FindById(s.db, id)
}

type CreateAlertRuleResp struct {
	ID string `json:"id"`
}

func (s *alertRulesService) Create(rule *models.AlertRules) (*CreateAlertRuleResp, error) {
	rule.ID = "rul-" + id.ShortID(8)
	// begin transaction
	tx := s.db.Begin()
	// create rule
	err := s.alertRulesRepo.Create(tx, rule)
	if err != nil {
		return nil, err
	}
	// create a relationship between rule and metric
	for _, v := range rule.RulesArr {
		v.RuleId = rule.ID
	}
	err = s.ruleMetricRelationRepo.Batches(tx, rule.RulesArr)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	s.setMetricExpressionValue(rule)
	err = ModifyPrometheusRuleAndReload(rule)
	if err == nil {
		httpclient.Request("http://127.0.0.1:9090/-/reload", "POST", nil, nil, nil)
	}
	return &CreateAlertRuleResp{
		ID: rule.ID,
	}, nil
}

func (s *alertRulesService) Delete(ruleId string) error {
	tx := s.db.Begin()
	err := s.alertRulesRepo.Delete(tx, ruleId)
	if err != nil {
		return err
	}
	err = s.ruleMetricRelationRepo.DeleteByRuleId(tx, ruleId)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	err = DeletePrometheusRuleAndReload(ruleId)
	if err == nil {
		httpclient.Request("http://127.0.0.1:9090/-/reload", "POST", nil, nil, nil)
	}
	return err
}

func (s *alertRulesService) Update(rule *models.AlertRules) (*CreateAlertRuleResp, error) {
	tx := s.db.Begin()
	err := s.alertRulesRepo.Update(tx, rule)
	if err != nil {
		return nil, err
	}
	for _, v := range rule.RulesArr {
		err = s.ruleMetricRelationRepo.Update(tx, v)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	s.setMetricExpressionValue(rule)
	err = ModifyPrometheusRuleAndReload(rule)
	if err == nil {
		httpclient.Request("http://127.0.0.1:9090/-/reload", "POST", nil, nil, nil)
	}
	return &CreateAlertRuleResp{
		ID: rule.ID,
	}, nil
}

func (s *alertRulesService) Page(page *Page.ReqPage) (*Page.RespPage, error) {
	return s.alertRulesRepo.Page(s.db, page)
}

func (s *alertRulesService) setMetricExpressionValue(rule *models.AlertRules) {
	ruleMetric, err := s.ruleMetricRelationRepo.GetRuleMetricByRuleId(s.db, rule.ID)
	if err == nil {
		rule.RulesArr = ruleMetric
	}
}
