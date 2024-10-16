package hooks

import (
	"github.com/StanZzzz222/RAltGo/common/core/retimer/internal/entities"
	"sync"
)

/*
	Create by zyx
	Date Time: 2024/10/16
	File: hooks.go
*/

var hooks = &sync.Map{}

type onTimerEventCallback func(timer *entities.Timer)

func OnTimerEvent(key string, callback onTimerEventCallback) {
	hooks.Store(key, callback)
}

func TriggerTimer(timer *entities.Timer) {
	if timer != nil {
		if hook, ok := hooks.Load(timer.Key); ok {
			hook.(onTimerEventCallback)(timer)
		}
	}
}
