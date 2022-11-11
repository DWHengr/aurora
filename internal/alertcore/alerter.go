package alertcore

import (
	"sync"
)

var alertSingleInstance Alerter
var m sync.Mutex

type Alerter interface {
	Receive(msg *AlertMessage) error
	AlertHandlerRegister(handler AlertHandler)
	AlertIntervalRegister(name string, interval *Interval)
	AlertSilenceRegister(name string, silence *Silence)
	ReloadHandlerRegister(handler ReloadHandler)
	Reload()
	Run()
}

func NewAlerterSingle(c *AlertConfig) Alerter {
	m.Lock()
	if alertSingleInstance == nil {
		alertSingleInstance = &alerter{
			thread:          c.Thread,
			buffer:          c.Buffer,
			alerterHandlers: []AlertHandler{},
			alertSilences:   make(map[string]*Silence),
			alertIntervals:  make(map[string]*Interval),
		}
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
