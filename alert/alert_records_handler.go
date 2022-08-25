package alert

import (
	"encoding/json"
	"github.com/DWHengr/aurora/internal/alertcore"
	"github.com/DWHengr/aurora/internal/models"
	"github.com/DWHengr/aurora/internal/service"
	"github.com/DWHengr/aurora/pkg/logger"
)

func RecordsHandler(message *alertcore.AlertMessage, ctx *alertcore.Context) {
	alertRecordsService, err := service.NewAlertRecordsService()
	if err != nil {
		logger.Logger.Error(err)
	}
	alertRulesService, err := service.NewAlertRulesService()
	if err != nil {
		logger.Logger.Error(err)
	}
	rule, err := alertRulesService.FindById(message.UniqueId)
	if err != nil {
		logger.Logger.Error(err)
	}

	record := &models.AlertRecords{
		AlertName: message.Name,
		RuleId:    rule.ID,
		RuleName:  rule.Name,
		Value:     message.Value,
		Severity:  rule.Severity,
		Summary:   message.Summary,
		Attribute: mapToString(message.Attribute),
	}
	err = alertRecordsService.CreateRecord(record)
	if err != nil {
		logger.Logger.Error(err)
	}
}

func mapToString(m map[string]interface{}) string {
	data, err := json.Marshal(m)
	if err != nil {
		logger.Logger.Error(err)
	}
	return string(data)
}
