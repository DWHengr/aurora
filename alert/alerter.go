package alert

import (
	"github.com/DWHengr/aurora/internal/alertcore"
	"github.com/DWHengr/aurora/internal/service"
	"github.com/DWHengr/aurora/pkg/config"
	"github.com/DWHengr/aurora/pkg/logger"
)

//LoadAlertHandler load user-defined alert handler
func LoadAlertHandler(alert alertcore.Alerter) {
	alert.AlertHandlerRegister(PrintfHandler)
	alert.AlertHandlerRegister(RecordsHandler)
}

//LoadAlertIntervalAndSilenceByMysql load alert interval time and silence by mysql
func LoadAlertIntervalAndSilenceByMysql(alert alertcore.Alerter) {
	alertRuleService, err := service.NewAlertRulesService()
	if err != nil {
		logger.Logger.Error(err)
	}
	alertRules, err := alertRuleService.GetAllAlertRules()
	if err != nil {
		logger.Logger.Error(err)
	}
	alertSilenceService, err := service.NewAlertSilencesService()
	if err != nil {
		logger.Logger.Error(err)
	}
	alertSilenceMap, err := alertSilenceService.GetAllAlertSilencesMap()
	if err != nil {
		logger.Logger.Error(err)
	}
	for _, rule := range alertRules {
		interval := &alertcore.Interval{
			IntervalTime: TimeStringToInt64(rule.AlertInterval),
		}
		alert.AlertIntervalRegister(rule.ID, interval)
		alertSilence := alertSilenceMap[rule.AlertSilencesId]
		silence := &alertcore.Silence{
			Type:      alertSilence.Type,
			StartTime: alertSilence.StartTime,
			EndTime:   alertSilence.EndTime,
		}
		alert.AlertSilenceRegister(rule.ID, silence)
	}
}

func NewAlerter(config *config.Config) alertcore.Alerter {
	alert := alertcore.NewAlerterSingle(&config.Alert)
	LoadAlertHandler(alert)
	LoadAlertIntervalAndSilenceByMysql(alert)
	return alert
}
