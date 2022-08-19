package alertcore

import (
	"sync"
	"time"
)

type Interval struct {
	SendTime     int64
	IntervalTime int64
}

type Silence struct {
	Type      string
	StartTime time.Time
	EndTime   time.Time
}

type alerter struct {
	thread          int
	buffer          int
	intervalMutex   sync.Mutex
	messages        chan *AlertMessage
	alertIntervals  map[string]*Interval
	alertSilences   map[string]*Silence
	alerterHandlers []AlertHandler
}

//alertHandlerRegister register alert handler
func (a *alerter) AlertHandlerRegister(handler AlertHandler) {
	a.alerterHandlers = append(a.alerterHandlers, handler)
}

//AlertIntervalRegister register alert interval
func (a *alerter) AlertIntervalRegister(name string, interval *Interval) {
	a.alertIntervals[name] = interval
}

//AlertIntervalRegister register alert silence
func (a *alerter) AlertSilenceRegister(name string, silence *Silence) {
	a.alertSilences[name] = silence
}

//verifyInterval verify interval time,Interval time when the return value is true
func (a *alerter) verifyInterval(name string) bool {
	nowTime := time.Now().Unix()
	//a.intervalMutex.Lock()
	//defer a.intervalMutex.Unlock()
	interval, ok := a.alertIntervals[name]
	if ok {
		if interval.SendTime != 0 && nowTime-interval.SendTime < interval.IntervalTime {
			return true
		}
		interval.SendTime = nowTime
		//atomic.StoreInt64(&interval.SendTime, nowTime)
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
		case "block":
			return TimeIsBlock(time, silence.StartTime, silence.EndTime)
		case "offday":
			return TimeIsOffDay(time)
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

func (a *alerter) Run() {
	a.messages = make(chan *AlertMessage, a.buffer)
	for index := 0; index < a.thread; index++ {
		go a.work()
	}
}

func (a *alerter) Receive(msg *AlertMessage) error {
	if a.verifySilence(msg.UniqueId) || a.verifyInterval(msg.UniqueId) {
		return nil
	}
	a.messages <- msg
	return nil
}
