package alert

import "github.com/DWHengr/aurora/internal/alertcore"

func NewAlerter() alertcore.Alerter {
	alert := alertcore.NewAlerterSingle()
	// register user-defined alert handler
	alert.AlertHandlerRegister(PrintfHandler)
	return alert
}
