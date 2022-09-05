package service

import (
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/pkg/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

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
	ruleYml.Annotations["value"] = "{{$value}}"
	// generate str: {k1="v1",k2="v2"}
	var alertObjKAndVArr []string
	for alertObjK, alertObjV := range alertRule.AlertObjectArr {
		alertObjKAndVArr = append(alertObjKAndVArr, alertObjK+"=\""+alertObjV+"\"")
	}
	alertObjKAndVStr := "{" + strings.Join(alertObjKAndVArr, ",") + "}"
	// generate str: metric1{k1="v1",k2="v2"}[Statistics1]>Value1 or metric2{k1="v1",k2="v2"}[Statistics2]<Value2
	var exprArr []string
	for _, rule := range alertRule.RulesArr {
		itemMetric := strings.Replace(rule.Expression, "${}", alertObjKAndVStr, 1)
		itemMetric = strings.Replace(itemMetric, "$[]", "["+rule.Statistics+"]", 1)
		exprArr = append(exprArr, itemMetric+" "+rule.Operator+" "+rule.AlertValue)
	}
	ruleYml.Expr = strings.Join(exprArr, " or ")
	return ruleYml
}

func ModifyPrometheusRuleAndReload(alertRules []*models.AlertRules) error {
	path := ""
	yamlFile, err := ioutil.ReadFile(path)
	prometheusYml := PrometheusYml{}
	if err == nil {
		if err := yaml.Unmarshal(yamlFile, &prometheusYml); err != nil {
			logger.Logger.Error(err)
			return err
		}
	}
	var groupYml = &GroupsYml{
		Name: "aurora.custom.defaults",
	}
	ruleYml := &RuleYml{
		Labels:      make(map[string]string),
		Annotations: make(map[string]string),
	}
	for _, alertRule := range alertRules {
		groupIsNotExist := true
		ruleIsNotExit := true
		for _, group := range prometheusYml.Groups {
			// find group content
			if group.Name == "aurora.custom.defaults" {
				groupIsNotExist = false
				groupYml = group
			}
			for _, rule := range groupYml.Rules {
				uniqueId, ok := rule.Labels["uniqueid"]
				// find rule content
				if ok && uniqueId == alertRule.ID {
					ruleIsNotExit = false
					ruleYml = rule
					break
				}
			}

		}
		CreatAndUpdateRule(ruleYml, alertRule)
		if ruleIsNotExit {
			groupYml.Rules = append(groupYml.Rules, ruleYml)
		}
		if groupIsNotExist {
			prometheusYml.Groups = append(prometheusYml.Groups, groupYml)
		}
	}
	//prometheus rule file out
	out, err := yaml.Marshal(prometheusYml)
	if err != nil {
		logger.Logger.Error(err)
		return err
	}
	if err = ioutil.WriteFile(path, out, 0666); err != nil {
		logger.Logger.Error(err)
		return err
	}
	return nil
}

func DeletePrometheusRuleAndReload(id string) error {
	path := ""
	yamlFile, err := ioutil.ReadFile(path)
	prometheusYml := PrometheusYml{}
	if err == nil {
		if err := yaml.Unmarshal(yamlFile, &prometheusYml); err != nil {
			logger.Logger.Error(err)
			return err
		}
	}
	for _, group := range prometheusYml.Groups {
		// find group content
		if group.Name == "aurora.custom.defaults" {
			for index, rule := range group.Rules {
				uniqueId, ok := rule.Labels["uniqueid"]
				// delete rule
				if ok && uniqueId == id {
					group.Rules = append(group.Rules[:index], group.Rules[index+1:]...)
					break
				}
			}
		}
	}
	//prometheus rule file out
	out, err := yaml.Marshal(prometheusYml)
	if err != nil {
		logger.Logger.Error(err)
		return err
	}
	if err = ioutil.WriteFile(path, out, 0666); err != nil {
		logger.Logger.Error(err)
		return err
	}
	return nil
}
