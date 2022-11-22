package alert

import (
	"github.com/DWHengr/aurora/internal/alertcore"
	"github.com/DWHengr/aurora/internal/service"
	"github.com/DWHengr/aurora/pkg/httpclient"
	"github.com/DWHengr/aurora/pkg/logger"
)

func WebHookHandler(message *alertcore.AlertMessage, ctx *alertcore.Context) {
	alertRulesService, err := service.NewAlertRulesService()
	if err != nil {
		logger.Logger.Error(err)
	}
	rule, err := alertRulesService.FindById(message.UniqueId)
	if rule == nil || err != nil {
		logger.Logger.Error(err)
		return
	}
	if len(rule.Webhook) >= 0 {
		err = httpclient.POST(nil, rule.Webhook, message, nil)
		if err != nil {
			logger.Logger.Errorw(err.Error())
		}
	}
}
