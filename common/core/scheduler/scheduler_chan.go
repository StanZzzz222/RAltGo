package scheduler

import (
	"github.com/StanZzzz222/RAltGo/internal/lib"
)

/*
   Create by zyx
   Date Time: 2024/10/15
   File: scheduler_chan.go
*/

var schedulerChan = make(chan func())

func init() {
	var w = lib.GetWrapper()
	go func() {
		for {
			select {
			case callback := <-schedulerChan:
				w.PushTask(callback)
			}
		}
	}()
}

func TaskScheduler(task func()) {
	schedulerChan <- task
}
