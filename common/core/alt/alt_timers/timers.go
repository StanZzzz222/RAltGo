package alt_timers

import (
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/9/12
   File: timers.go
*/

func SetInterval(duration time.Duration, callback func()) *time.Ticker {
	var w = lib.GetWarpper()
	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ticker.C:
				w.PushTask(callback)
			}
		}
	}()
	return ticker
}

func SetTimeout(duration time.Duration, callback func()) {
	var w = lib.GetWarpper()
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
