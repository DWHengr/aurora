package alert

import (
	"time"
)

type interval struct {
	SendTime     int64
	IntervalTime int64
}

type alerter struct {
	messages        chan *AlertMessage
	alertIntervals  map[string]interval
	alerterHandlers []AlertHandler
}

//alertHandlerRegister register alert handler
func (a *alerter) alertHandlerRegister(handler AlertHandler) {
	a.alerterHandlers = append(a.alerterHandlers, handler)
}

//verifyInterval verify interval time
func (a *alerter) verifyInterval(name string) bool {
	interval, ok := a.alertIntervals[name]
	if ok {
		nowTime := time.Now().Unix()
		if nowTime-interval.SendTime >= interval.IntervalTime {
			return true
		}
		return false
	}
	return true
}

//verifySilence verify Silence time
func (a *alerter) verifySilence(name string) bool {
	interval, ok := a.alertIntervals[name]
	if ok {
		nowTime := time.Now().Unix()
		if nowTime-interval.SendTime >= interval.IntervalTime {
			return true
		}
		return false
	}
	return true
}

//work the work thread used for call handler
func (a *alerter) work() {
	for {
		message := <-a.messages
		for _, handler := range a.alerterHandlers {
			go handler(message)
		}
	}
}

func (a *alerter) run(c *AlertConfig) {
	a.messages = make(chan *AlertMessage, c.Buffer)
	for index := 0; index < c.Thread; index++ {
		go a.work()
	}
}

func (a *alerter) Receive(msg *AlertMessage) error {
	// todo  verify silence period
	//if a.verifyInterval(msg.Name) {
	//	// todo send alert
	//}
	a.messages <- msg
	return nil
}
