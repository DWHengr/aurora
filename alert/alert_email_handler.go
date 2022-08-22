package alert

import (
	"bytes"
	"github.com/DWHengr/aurora/internal/alertcore"
	"github.com/DWHengr/aurora/internal/service"
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/DWHengr/aurora/pkg/misc/email"
	email2 "github.com/jordan-wright/email"
	"html/template"
	"time"
)

func EmailHandler(message *alertcore.AlertMessage) {
	alertRulesService, err := service.NewAlertRulesService()
	if err != nil {
		logger.Logger.Error(err)
	}
	rule, err := alertRulesService.FindById(message.UniqueId)
	if err != nil {
		logger.Logger.Error(err)
	}

	sender := email.GetEmailSender()
	e := email2.NewEmail()
	t, err := template.ParseFiles("configs/email_template.html")
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	body := new(bytes.Buffer)
	t.Execute(body, struct {
		AlertName string
		RuleName  string
		Value     string
		Severity  string
		Summary   string
		Attribute map[string]interface{}
		NowTime   time.Time
	}{AlertName: message.Name,
		RuleName:  rule.Name,
		Value:     message.Value,
		Severity:  rule.Severity,
		Summary:   message.Summary,
		Attribute: message.Attribute,
		NowTime:   time.Unix(time.Now().Unix(), 0)})
	e.HTML = body.Bytes()
	e.From = "xxxx@163.com"
	e.To = []string{"xxxxx@qq.com"}
	e.Subject = "alert email"
	err = sender.Send(e, 10*time.Second)
	if err != nil {
		logger.Logger.Error(err)
	}
}
