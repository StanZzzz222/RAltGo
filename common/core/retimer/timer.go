package retimer

import (
	"github.com/StanZzzz222/RAltGo/common/core/retimer/entities"
	"github.com/StanZzzz222/RAltGo/common/core/retimer/internal/hooks"
	"sync"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/10/16
   File: retimer.go
*/

var (
	timers = &sync.Map{}
)

func init() {
	tickerMillsecond := time.NewTicker(time.Millisecond * 1)
	tickerSecond := time.NewTicker(time.Second * 1)
	tickerMinute := time.NewTicker(time.Minute * 1)
	tickerHour := time.NewTicker(time.Hour * 1)
	tickerDay := time.NewTicker(time.Hour * 24)
	go func() {
		for {
			select {
			case <-tickerMillsecond.C:
				check(getTimersBySecond())
				continue
			case <-tickerSecond.C:
				check(getTimersBySecond())
				continue
			case <-tickerMinute.C:
				check(getTimersByMinute())
				continue
			case <-tickerHour.C:
				check(getTimersByHour())
				continue
			case <-tickerDay.C:
				check(getTimersByDay())
				continue
			}
		}
	}()
}

func OnTimerEvent(key string, callback hooks.OnTimerEventCallback) {
	hooks.OnTimerEvent(key, callback)
}

func CreateTimer(key string, duration time.Duration, loop bool, loopCount int) {
	timer := entities.NewTimer(key, int64(duration), loop, loopCount)
	timers.Store(timer.Key, timer)
}

func GetTimer(key string) *entities.Timer {
	timer, ok := timers.Load(key)
	if ok {
		return timer.(*entities.Timer)
	}
	return nil
}

func DelTimer(key string) {
	_, ok := timers.Load(key)
	if ok {
		timers.Delete(key)
		return
	}
}

func RestartTimer(key string) {
	timer := GetTimer(key)
	if timer != nil {
		DelTimer(key)
		timer.NotifyMillisecond = time.Now().Add(time.Duration(timer.Millisecond) * time.Millisecond).UTC().Unix()
		CreateTimer(timer.Key, time.Duration(timer.Millisecond), timer.Loop, timer.LoopCount)
	}
}

func TTLTimer(key string) int64 {
	timer := GetTimer(key)
	if timer != nil {
		second := (timer.NotifyMillisecond * 1000) - time.Now().UTC().Unix()
		if second <= 0 {
			return 0
		}
		return second
	}
	return -1
}

func getTimersByMillsecond() *sync.Map {
	millsecondTimers := &sync.Map{}
	timers.Range(func(key, value any) bool {
		if value != nil {
			timer := value.(*entities.Timer)
			if timer.Millisecond < 1000 {
				millsecondTimers.Store(key, timer)
				return true
			}
		}
		return true
	})
	return millsecondTimers
}

func getTimersBySecond() *sync.Map {
	secondTimers := &sync.Map{}
	timers.Range(func(key, value any) bool {
		if value != nil {
			timer := value.(*entities.Timer)
			if timer.Millisecond >= 1000 {
				secondTimers.Store(key, timer)
				return true
			}
		}
		return true
	})
	return secondTimers
}

func getTimersByMinute() *sync.Map {
	minuteTimers := &sync.Map{}
	timers.Range(func(key, value any) bool {
		if value != nil {
			timer := value.(*entities.Timer)
			if timer.Millisecond >= (60 * 1000) {
				minuteTimers.Store(key, timer)
				return true
			}
		}
		return true
	})
	return minuteTimers
}

func getTimersByHour() *sync.Map {
	hourTimers := &sync.Map{}
	timers.Range(func(key, value any) bool {
		if value != nil {
			timer := value.(*entities.Timer)
			if timer.Millisecond >= (60 * 60 * 1000) {
				hourTimers.Store(key, timer)
				return true
			}
		}
		return true
	})
	return hourTimers
}

func getTimersByDay() *sync.Map {
	dayTimers := &sync.Map{}
	timers.Range(func(key, value any) bool {
		if value != nil {
			timer := value.(*entities.Timer)
			if timer.Millisecond >= (60 * 60 * 1000) {
				dayTimers.Store(key, timer)
				return true
			}
		}
		return true
	})
	return dayTimers
}

func check(timers *sync.Map) {
	timers.Range(func(key, value any) bool {
		if value != nil {
			timer := value.(*entities.Timer)
			if !timer.IsPause {
				if time.Now().Add(time.Millisecond*1).UTC().Unix() >= timer.NotifyMillisecond {
					if !timer.Loop {
						DelTimer(timer.Key)
						hooks.TriggerTimer(timer)
					} else {
						if timer.LoopCount > 0 {
							lastLoopCount := timer.LoopCount - 1
							if lastLoopCount == 0 {
								DelTimer(timer.Key)
								hooks.TriggerTimer(timer)
							} else {
								timer.LoopCount = lastLoopCount
								timer.NotifyMillisecond = time.Now().Add(time.Duration(timer.Millisecond) * time.Millisecond).UTC().Unix()
								hooks.TriggerTimer(timer)
							}
						} else {
							timer.NotifyMillisecond = time.Now().Add(time.Duration(timer.Millisecond) * time.Millisecond).UTC().Unix()
							hooks.TriggerTimer(timer)
						}
					}
				}
			}
		}
		return true
	})
}
