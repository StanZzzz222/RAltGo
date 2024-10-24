package retimer

import (
	"github.com/StanZzzz222/RAltGo/common/core/retimer/internal/hooks"
	"github.com/StanZzzz222/RAltGo/common/core/retimer/internal/scripts"
	"github.com/StanZzzz222/RAltGo/common/core/retimer/timer"
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

func OnTimerEvent(key string, preffixMatch bool, callback hooks.OnTimerEventCallback) {
	hooks.OnTimerEvent(key, preffixMatch, callback)
}

func CreateTimer(key string, duration time.Duration, loop bool) *timer.ITimer {
	t := timer.NewTimer(key, duration, loop, 0)
	timers.Store(t.Key, t)
	return t
}

func CreateLoopCountTimer(key string, duration time.Duration, loopCount int) *timer.ITimer {
	t := timer.NewTimer(key, duration, true, loopCount)
	timers.Store(t.Key, t)
	return t
}

func CreateScriptTimer(key string, expr string, duration time.Duration, loop bool) *timer.ITimer {
	t := timer.NewTimer(key, duration, loop, 0)
	t.Expr = expr
	scripts.ExecuteTimerExpression(t)
	timers.Store(t.Key, t)
	return t
}

func CreateLoopCountScriptTimer(key string, expr string, duration time.Duration, loopCount int) *timer.ITimer {
	t := timer.NewTimer(key, duration, true, loopCount)
	t.Expr = expr
	scripts.ExecuteTimerExpression(t)
	timers.Store(t.Key, t)
	return t
}

func GetTimer(key string) *timer.ITimer {
	t, ok := timers.Load(key)
	if ok {
		return t.(*timer.ITimer)
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

func StopTimer(timer *timer.ITimer) {
	if timer != nil {
		timer.IsPause = true
	}
}

func PlayTimer(timer *timer.ITimer) {
	if timer != nil {
		timer.IsPause = false
	}
}

func TTLTimer(key string) int64 {
	t := GetTimer(key)
	if t != nil {
		second := t.NotifyUnix - time.Now().UTC().Unix()
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
			t := value.(*timer.ITimer)
			if t.Duration < 1000 {
				millsecondTimers.Store(key, t)
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
			t := value.(*timer.ITimer)
			if t.Duration >= 1000 {
				secondTimers.Store(key, t)
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
			t := value.(*timer.ITimer)
			if t.Duration >= (60 * 1000) {
				minuteTimers.Store(key, t)
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
			t := value.(*timer.ITimer)
			if t.Duration >= (60 * 60 * 1000) {
				hourTimers.Store(key, t)
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
			t := value.(*timer.ITimer)
			if t.Duration >= (60 * 60 * 1000) {
				dayTimers.Store(key, t)
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
			t := value.(*timer.ITimer)
			if !t.IsPause {
				if time.Now().Add(time.Millisecond).UTC().Unix() >= t.NotifyUnix {
					if !t.Loop {
						Destroy(t.Key)
						hooks.TriggerTimer(t)
					} else {
						if t.LoopCount > 0 {
							lastLoopCount := t.LoopCount - 1
							if lastLoopCount == 0 {
								Destroy(t.Key)
								hooks.TriggerTimer(t)
							} else {
								t.LoopCount = lastLoopCount
								t.NotifyUnix = time.Now().Add(t.Duration).Unix()
								hooks.TriggerTimer(t)
							}
						} else {
							t.NotifyUnix = time.Now().Add(t.Duration).Unix()
							hooks.TriggerTimer(t)
						}
					}
				}
			}
		}
		return true
	})
}
