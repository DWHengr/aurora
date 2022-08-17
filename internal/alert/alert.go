package alert

import (
	"fmt"
	"time"
)

var AlertInstance Alerter

type Alerter interface {
	Receive(msg *AlertMessage) error
	run(c *AlertConfig)
}

func NewAlertAndRun(c *AlertConfig) {
	AlertInstance = &alerter{}
	AlertInstance.run(c)
}

type interval struct {
	SendTime     int64
	IntervalTime int64
}

type alerter struct {
	messages       chan *AlertMessage
	alertIntervals map[string]interval
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

func (a *alerter) work(n int) {
	for {
		message := <-a.messages
		//todo send
		fmt.Println(n, message)
		time.Sleep(time.Duration(5) * time.Second)
	}
}

func (a *alerter) run(c *AlertConfig) {
	a.messages = make(chan *AlertMessage, c.Buffer)
	for index := 0; index < c.Thread; index++ {
		go a.work(index)
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
