package alertcore

import (
	"sync"
	"time"
)

type interval struct {
	SendTime     int64
	IntervalTime int64
}

type silence struct {
	Type      string
	StartTime time.Time
	EndTime   time.Time
}

type alerter struct {
	intervalMutex   sync.Mutex
	messages        chan *AlertMessage
	alertIntervals  map[string]*interval
	alertSilences   map[string]*silence
	alerterHandlers []AlertHandler
}

//alertHandlerRegister register alertcore handler
func (a *alerter) AlertHandlerRegister(handler AlertHandler) {
	a.alerterHandlers = append(a.alerterHandlers, handler)
}

//verifyInterval verify interval time,Interval time when the return value is true
func (a *alerter) verifyInterval(name string) bool {
	nowTime := time.Now().Unix()
	a.intervalMutex.Lock()
	interval, ok := a.alertIntervals[name]
	if !ok && interval.SendTime == 0 {
		interval.SendTime = nowTime
		a.intervalMutex.Unlock()
		return false
	}
	if ok {
		if nowTime-interval.SendTime >= interval.IntervalTime {
			return true
		}
		return false
	}
	return false
}

//verifySilence verify Silence time,Silent when the return value is true
func (a *alerter) verifySilence(name string) bool {
	silence, ok := a.alertSilences[name]
	time := time.Now()
	if ok {
		switch silence.Type {
		case "everyday":
			return TimeIsEveryday(time, silence.StartTime, silence.EndTime)
			break
		case "block":
			return TimeIsBlock(time, silence.StartTime, silence.EndTime)
			break
		case "offday":
			return TimeIsOffDay(time)
			break
		}
	}
	return false
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

func (a *alerter) Run(c *AlertConfig) {
	a.messages = make(chan *AlertMessage, c.Buffer)
	for index := 0; index < c.Thread; index++ {
		go a.work()
	}
}

func (a *alerter) Receive(msg *AlertMessage) error {
	// todo  verify silence period
	//if a.verifyInterval(msg.Name) {
	//	// todo send alertcore
	//}
	a.messages <- msg
	return nil
}
