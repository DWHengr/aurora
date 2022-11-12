package service

import (
	"fmt"
	"github.com/DWHengr/aurora/internal/alertcore"
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/internal/page"
	"github.com/DWHengr/aurora/pkg/config"
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
	Page(page *page.ReqPage) (*page.RespPage, error)
	Deletes(ids []string) error
	Details(id string) (*models.AlertRules, error)
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
		alertRulesRepo:         mysql.NewAlertRulesRepo(),
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

func (s *alertRulesService) Details(id string) (*models.AlertRules, error) {
	rule, err := s.alertRulesRepo.FindById(s.db, id)
	if err != nil {
		return nil, err
	}
	rulesArr, err := s.ruleMetricRelationRepo.GetRuleMetricByRuleId(s.db, id)
	rule.RulesArr = rulesArr
	return rule, nil
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
	err = ModifyPrometheusRuleAndReload([]*models.AlertRules{rule})
	if err == nil {
		allConfig, _ := config.GetAllConfig()
		httpclient.Request(allConfig.Aurora.PrometheusUrl+"/-/reload", "POST", nil, nil, nil)
		alertcore.Reload()
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
	err = DeletePrometheusRuleAndReload([]string{ruleId})
	if err == nil {
		allConfig, _ := config.GetAllConfig()
		httpclient.Request(allConfig.Aurora.PrometheusUrl+"/-/reload", "POST", nil, nil, nil)
		alertcore.Reload()
	}
	return err
}

func (s *alertRulesService) Deletes(ids []string) error {
	tx := s.db.Begin()
	err := s.alertRulesRepo.Deletes(tx, ids)
	if err != nil {
		return err
	}
	err = s.ruleMetricRelationRepo.DeleteByRuleIds(tx, ids)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	err = DeletePrometheusRuleAndReload(ids)
	if err == nil {
		allConfig, _ := config.GetAllConfig()
		httpclient.Request(allConfig.Aurora.PrometheusUrl+"/-/reload", "POST", nil, nil, nil)
		alertcore.Reload()
	}
	return err
}

func (s *alertRulesService) Update(rule *models.AlertRules) (*CreateAlertRuleResp, error) {
	tx := s.db.Begin()
	err := s.alertRulesRepo.Update(tx, rule)
	if err != nil {
		return nil, err
	}
	s.ruleMetricRelationRepo.DeleteByRuleId(tx, rule.ID)
	for _, v := range rule.RulesArr {
		v.RuleId = rule.ID
	}
	err = s.ruleMetricRelationRepo.Batches(tx, rule.RulesArr)
	tx.Commit()
	//if rule.RulesStatus == RuleStatusDisabled {
	//	err = DeletePrometheusRuleAndReload([]string{rule.ID})
	//} else if rule.RulesStatus == RuleStatusEnable {
	//	s.setMetricExpressionValue(rule)
	//	err = ModifyPrometheusRuleAndReload([]*models.AlertRules{rule})
	//}
	s.setMetricExpressionValue(rule)
	err = ModifyPrometheusRuleAndReload([]*models.AlertRules{rule})
	if err == nil {
		allConfig, _ := config.GetAllConfig()
		err := httpclient.Request(allConfig.Aurora.PrometheusUrl+"/-/reload", "POST", nil, nil, nil)
		if err != nil {
			fmt.Println(err)
		}
		alertcore.Reload()
	}
	return &CreateAlertRuleResp{
		ID: rule.ID,
	}, nil
}

func (s *alertRulesService) Page(page *page.ReqPage) (*page.RespPage, error) {
	return s.alertRulesRepo.Page(s.db, page)
}

func (s *alertRulesService) setMetricExpressionValue(rule *models.AlertRules) {
	ruleMetric, err := s.ruleMetricRelationRepo.GetRuleMetricByRuleId(s.db, rule.ID)
	if err == nil {
		rule.RulesArr = ruleMetric
	}
}
