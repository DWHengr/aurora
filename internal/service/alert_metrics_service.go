package service

import (
	"errors"
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/internal/page"
	"github.com/DWHengr/aurora/pkg/config"
	"github.com/DWHengr/aurora/pkg/httpclient"
	"github.com/DWHengr/aurora/pkg/id"
	"gorm.io/gorm"
	"strings"
)

type AlertMetricsService interface {
	GetAllAlertMetrics() ([]*models.AlertMetrics, error)
	Create(rule *models.AlertMetrics) (*CreateAlertMetricResp, error)
	Delete(metricId string) error
	Page(page *page.ReqPage) (*page.RespPage, error)
	Deletes(ids []string) error
	Update(metric *models.AlertMetrics) (*CreateAlertMetricResp, error)
}

type alertMetricsService struct {
	db                     *gorm.DB
	alertMetricsRepo       models.AlertMetricsRepo
	alertRulesRepo         models.AlertRulesRepo
	ruleMetricRelationRepo models.RuleMetricRelationRepo
}

func NewAlertMetricsService() (AlertMetricsService, error) {
	db := GetMysqlInstance()

	return &alertMetricsService{
		db:                     db,
		alertMetricsRepo:       mysql.NewAlertMetricsRepo(),
		alertRulesRepo:         mysql.NewAlertRulesRepo(),
		ruleMetricRelationRepo: mysql.NewRuleMetricRelationRepo(),
	}, nil
}

func (s *alertMetricsService) GetAllAlertMetrics() ([]*models.AlertMetrics, error) {
	tables, err := s.alertMetricsRepo.GetAll(s.db)
	if err != nil {
		return nil, err
	}
	// TODO
	return tables, err
}

type CreateAlertMetricResp struct {
	ID string `json:"id"`
}

func (s *alertMetricsService) Create(metric *models.AlertMetrics) (*CreateAlertMetricResp, error) {
	metric.ID = "mtc-" + id.ShortID(8)
	err := s.alertMetricsRepo.Create(s.db, metric)
	if err != nil {
		return nil, err
	}
	return &CreateAlertMetricResp{
		ID: metric.ID,
	}, nil
}

func (s *alertMetricsService) Delete(metricId string) error {
	count, err := s.ruleMetricRelationRepo.GetCountByMetricID(s.db, metricId)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("this metric has alert in use")
	}
	err = s.alertMetricsRepo.Delete(s.db, metricId)
	if err != nil {
		return err
	}
	return err

}

func (s *alertMetricsService) Deletes(ids []string) error {
	usedIds := make([]string, 0)
	for _, id := range ids {
		count, err := s.ruleMetricRelationRepo.GetCountByMetricID(s.db, id)
		if err == nil && count > 0 {
			usedIds = append(usedIds, id)
		}
	}
	if len(usedIds) > 0 {
		return errors.New(strings.Join(usedIds, ",") + " these metric has alert in use")
	}
	err := s.alertMetricsRepo.Deletes(s.db, ids)
	return err
}

func (s *alertMetricsService) Page(page *page.ReqPage) (*page.RespPage, error) {
	return s.alertMetricsRepo.Page(s.db, page)
}

func (s *alertMetricsService) Update(metric *models.AlertMetrics) (*CreateAlertMetricResp, error) {
	dbMetric, _ := s.alertMetricsRepo.FindById(s.db, metric.ID)
	if dbMetric != nil && metric.Expression != dbMetric.Expression {
		// find all rule that use this metric,update prometheus rules file
		ruleIds, err := s.ruleMetricRelationRepo.FindRuleIdsByMetricId(s.db, metric.ID)
		if err != nil {
			return nil, err
		}
		rules, err := s.alertRulesRepo.FindByIds(s.db, *ruleIds)
		if err != nil {
			return nil, err
		}
		for _, rule := range rules {
			s.setMetricExpressionValue(rule)
		}
		err = ModifyPrometheusRuleAndReload(rules)
		if err == nil {
			allConfig, _ := config.GetAllConfig()
			httpclient.Request(allConfig.Aurora.PrometheusUrl+"/-/reload", "POST", nil, nil, nil)
		}
	}
	err := s.alertMetricsRepo.Update(s.db, metric)
	return &CreateAlertMetricResp{
		ID: metric.ID,
	}, err
}

func (s *alertMetricsService) setMetricExpressionValue(rule *models.AlertRules) {
	ruleMetric, err := s.ruleMetricRelationRepo.GetRuleMetricByRuleId(s.db, rule.ID)
	if err == nil {
		rule.RulesArr = ruleMetric
	}
}
