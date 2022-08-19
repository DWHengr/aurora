package alertcore

import (
	"sync"
)

var alertSingleInstance Alerter
var m sync.Mutex

type Alerter interface {
	Receive(msg *AlertMessage) error
	AlertHandlerRegister(handler AlertHandler)
	Run(c *AlertConfig)
}

func NewAlerterSingle() Alerter {
	m.Lock()
	if alertSingleInstance == nil {
		alertSingleInstance = &alerter{}
	}
	m.Unlock()
	return alertSingleInstance
}

func GetAlerterSingle() Alerter {
	if alertSingleInstance == nil {
		panic("alter instance is nil, first call NewAlerterSingle() to initialize alter single instance")
	}
	return alertSingleInstance
}
