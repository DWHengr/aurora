package alert

var AlertInstance Alerter = &alerter{}

type Alerter interface {
	Receive(msg *AlertMessage) error
	alertHandlerRegister(handler AlertHandler)
	run(c *AlertConfig)
}

func Run(c *AlertConfig) {
	AlertInstance.run(c)
}

func AlertHandlerRegister(handler AlertHandler) {
	AlertInstance.alertHandlerRegister(handler)
}
