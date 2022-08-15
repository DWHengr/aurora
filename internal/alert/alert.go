package alert

import "time"

type Alerter interface {
	Receive(msg AlertMessage) error
}

func NewAlert() Alerter {
	return &alerter{}
}

type interval struct {
	SendTime     int64
	IntervalTime int64
}

type alerter struct {
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

func (a *alerter) Receive(msg AlertMessage) error {
	// todo  verify silence period
	if a.verifyInterval(msg.Name) {
		// todo send alert
	}
	return nil
}
