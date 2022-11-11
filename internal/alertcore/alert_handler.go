package alertcore

type AlertHandler func(message *AlertMessage, ctx *Context)
type ReloadHandler func(alert Alerter)
