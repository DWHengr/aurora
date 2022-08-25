package alertcore

type AlertHandler func(message *AlertMessage, ctx *Context)
