package hooks

import (
	"github.com/StanZzzz222/RAltGo/common/core/alt/alt_timers"
	"github.com/StanZzzz222/RAltGo/common/core/retimer/timer"
	"strings"
	"sync"
	"time"
)

/*
	Create by zyx
	Date Time: 2024/10/16
	File: hooks.go
*/

var hooks = &sync.Map{}
var hookPreffixMatch = &sync.Map{}

type OnTimerEventCallback func(timer *timer.ITimer)

func OnTimerEvent(key string, preffixMatch bool, callback OnTimerEventCallback) {
	hooks.Store(key, callback)
	if preffixMatch {
		hookPreffixMatch.Store(key, true)
	}
}

func TriggerTimer(timer *timer.ITimer) {
	if timer != nil {
		if hook, ok := hooks.Load(timer.Key); ok {
			alt_timers.SetTimeout(time.Microsecond, func() {
				hook.(OnTimerEventCallback)(timer)
			})
		}
		hookPreffixMatch.Range(func(k, v any) bool {
			key := k.(string)
			if strings.Contains(timer.Key, key) {
				if hook, ok := hooks.Load(key); ok {
					alt_timers.SetTimeout(time.Microsecond, func() {
						hook.(OnTimerEventCallback)(timer)
					})
					return false
				}
			}
			return true
		})
	}
}
