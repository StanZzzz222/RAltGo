package timers

import (
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/9/12
   File: timers.go
*/

var w = &lib.Warrper{}

func SetInterval(duration time.Duration, callback func()) {
	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ticker.C:
				w.PushTask(callback)
			}
		}
	}()
}

func SetTimeout(duration time.Duration, callback func()) {
	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ticker.C:
				w.PushTask(callback)
				ticker.Stop()
				break
			}
		}
	}()
}
