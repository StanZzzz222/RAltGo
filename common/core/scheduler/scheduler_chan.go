package scheduler

import (
	"github.com/StanZzzz222/RAltGo/common/core/alt/alt_timers"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/10/15
   File: scheduler_chan.go
*/

var schedulerChan = make(chan func())

func init() {
	go func() {
		for {
			select {
			case task := <-schedulerChan:
				alt_timers.SetTimeout(time.Microsecond, func() {
					task()
				})
			}
		}
	}()
}

func TaskScheduler(task func()) {
	schedulerChan <- task
}
