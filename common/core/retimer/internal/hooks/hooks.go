package hooks

import (
	"github.com/StanZzzz222/RAltGo/common/core/alt/alt_timers"
	"github.com/StanZzzz222/RAltGo/common/core/retimer/timer"
	"sync"
	"time"
)

/*
	Create by zyx
	Date Time: 2024/10/16
	File: hooks.go
*/

var hooks = &sync.Map{}

type OnTimerEventCallback func(timer *timer.ITimer)

func OnTimerEvent(key string, callback OnTimerEventCallback) {
	hooks.Store(key, callback)
}

func TriggerTimer(timer *timer.ITimer) {
	if timer != nil {
		if hook, ok := hooks.Load(timer.Key); ok {
			alt_timers.SetTimeout(time.Microsecond, func() {
				hook.(OnTimerEventCallback)(timer)
			})
		}
	}
}
