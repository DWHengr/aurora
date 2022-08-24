package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/models/mysql"
	"github.com/DWHengr/aurora/pkg/logger"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
)

type AlertRulesService interface {
	GetAllAlertRules() ([]*models.AlertRules, error)
	FindById(id string) (*models.AlertRules, error)
}

type alertRulesService struct {
	db             *gorm.DB
	alertRulesRepo models.AlertRulesRepo
}

func NewAlertRulesService() (AlertRulesService, error) {
	db := GetMysqlInstance()

	return &alertRulesService{
		db:             db,
		alertRulesRepo: mysql.NewAlterRulesRepo(),
	}, nil
}
func (s *alertRulesService) GetAllAlertRules() ([]*models.AlertRules, error) {
	tables, err := s.alertRulesRepo.GetAll(s.db)
	if err != nil {
		return nil, err
	}
	// TODO
	return tables, err
}

func (s *alertRulesService) FindById(id string) (*models.AlertRules, error) {
	return s.alertRulesRepo.FindById(s.db, id)
}

type RuleYml struct {
	Alert       string            `yaml:"alert"`
	Annotations map[string]string `yaml:"annotations"`
	Expr        string            `yaml:"expr"`
	For         string            `yaml:"for"`
	Labels      map[string]string `yaml:"labels"`
}

type GroupsYml struct {
	Name  string     `yaml:"name"`
	Rules []*RuleYml `yaml:"rules"`
}

type PrometheusYml struct {
	Groups []*GroupsYml `yaml:"groups"`
}

func CreatAndUpdateRule(ruleYml *RuleYml, alertRule *models.AlertRules) *RuleYml {
	ruleYml.Labels["uniqueid"] = alertRule.ID
	ruleYml.Alert = alertRule.Name
	ruleYml.For = alertRule.Persistent
	ruleYml.Labels["severity"] = alertRule.Severity
	ruleYml.Annotations["summary"] = alertRule.Name
	ruleYml.Annotations["value"] = "{{value}}"
	// generate str: {k1="v1",k2="v2"}
	var alertObjKAndVArr []string
	for alertObjK, alertObjV := range alertRule.AlertObjectArr {
		alertObjKAndVArr = append(alertObjKAndVArr, alertObjK+"=\""+alertObjV+"\"")
	}
	alertObjKAndVStr := "{" + strings.Join(alertObjKAndVArr, ",") + "}"
	// generate str: metric1{k1="v1",k2="v2"}[Statistics1]>Value1 or metric2{k1="v1",k2="v2"}[Statistics2]<Value2
	var exprArr []string
	for _, rule := range alertRule.RulesArr {
		exprArr = append(exprArr, rule.Metric+alertObjKAndVStr+"["+rule.Statistics+"] "+rule.Operator+" "+rule.AlertValue)
	}
	ruleYml.Expr = strings.Join(exprArr, " or ")
	return ruleYml
}

func ModifyPrometheusRuleAndReload(alertRule *models.AlertRules) {
	path := ""
	yamlFile, err := ioutil.ReadFile(path)
	prometheusYml := PrometheusYml{}
	if err == nil {
		if err := yaml.Unmarshal(yamlFile, &prometheusYml); err != nil {
			logger.Logger.Error(err)
		}
	}

	var groupYml = &GroupsYml{
		Name: "aurora.custom.defaults",
	}
	ruleYml := &RuleYml{
		Labels:      make(map[string]string),
		Annotations: make(map[string]string),
	}
	for _, group := range prometheusYml.Groups {
		// find group content
		if group.Name == "aurora.custom.defaults" {
			groupYml = group
		}
		for _, rule := range groupYml.Rules {
			uniqueId, ok := rule.Labels["uniqueid"]
			// find rule content
			if ok && uniqueId == alertRule.ID {
				ruleYml = rule
				break
			}
		}

	}
	CreatAndUpdateRule(ruleYml, alertRule)
	if len(groupYml.Rules) == 0 {
		groupYml.Rules = append(groupYml.Rules, ruleYml)
	}
	if len(groupYml.Rules) == 0 {
		prometheusYml.Groups = append(prometheusYml.Groups, groupYml)
	}
	//prometheus rule file out
	out, err := yaml.Marshal(prometheusYml)
	if err != nil {
		logger.Logger.Error(err)
	}
	if err = ioutil.WriteFile(path, out, 0666); err != nil {
		logger.Logger.Error(err)
	}
}
