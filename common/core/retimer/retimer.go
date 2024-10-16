package retimer

import (
	"github.com/StanZzzz222/RAltGo/common/core/retimer/internal/hooks"
	"github.com/StanZzzz222/RAltGo/common/core/retimer/internal/scripts"
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
				check(getTimersByMillsecond())
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

func CreateTimer(key string, duration time.Duration, loop bool) *Timer {
	timer := NewTimer(key, int64(duration), loop, 0)
	timers.Store(timer.Key, timer)
	return timer
}

func CreateLoopCountTimer(key string, duration time.Duration, loopCount int) *Timer {
	timer := NewTimer(key, int64(duration), true, loopCount)
	timers.Store(timer.Key, timer)
	return timer
}

func CreateScriptTimer(key string, expr string, duration time.Duration, loop bool) *Timer {
	timer := NewTimer(key, int64(duration), loop, 0)
	timer.Expr = expr
	scripts.ExecuteTimerExpression(timer)
	timers.Store(timer.Key, timer)
	return timer
}

func CreateLoopCountScriptTimer(key string, expr string, duration time.Duration, loopCount int) *Timer {
	timer := NewTimer(key, int64(duration), true, loopCount)
	timer.Expr = expr
	scripts.ExecuteTimerExpression(timer)
	timers.Store(timer.Key, timer)
	return timer
}

func GetTimer(key string) *Timer {
	timer, ok := timers.Load(key)
	if ok {
		return timer.(*Timer)
	}
	return nil
}

func Destroy(key string) {
	_, ok := timers.Load(key)
	if ok {
		timers.Delete(key)
		return
	}
}

func StopTimer(timer *Timer) {
	if timer != nil {
		timer.IsPause = true
	}
}

func PlayTimer(timer *Timer) {
	if timer != nil {
		timer.IsPause = false
	}
}

func TTLTimer(key string) int64 {
	timer := GetTimer(key)
	if timer != nil {
		second := timer.NotifyUnix - time.Now().UTC().Unix()
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
			timer := value.(*Timer)
			if timer.Duration < 1000 {
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
			timer := value.(*Timer)
			if timer.Duration >= 1000 {
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
			timer := value.(*Timer)
			if timer.Duration >= (60 * 1000) {
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
			timer := value.(*Timer)
			if timer.Duration >= (60 * 60 * 1000) {
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
			timer := value.(*Timer)
			if timer.Duration >= (60 * 60 * 1000) {
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
			timer := value.(*Timer)
			if !timer.IsPause {
				if time.Now().Add(time.Millisecond).UTC().Unix() >= timer.NotifyUnix {
					if !timer.Loop {
						Destroy(timer.Key)
						hooks.TriggerTimer(timer)
					} else {
						if timer.LoopCount > 0 {
							lastLoopCount := timer.LoopCount - 1
							if lastLoopCount == 0 {
								Destroy(timer.Key)
								hooks.TriggerTimer(timer)
							} else {
								timer.LoopCount = lastLoopCount
								timer.NotifyUnix = time.Now().Add(time.Duration(timer.NotifyUnix)).Unix()
								hooks.TriggerTimer(timer)
							}
						} else {
							timer.NotifyUnix = time.Now().Add(time.Duration(timer.NotifyUnix)).Unix()
							hooks.TriggerTimer(timer)
						}
					}
				}
			}
		}
		return true
	})
}
