package tests

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/core/retimer"
	"github.com/StanZzzz222/RAltGo/common/core/retimer/timer"
	"testing"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/10/16
   File: timer_test.go
*/

func TestTimer(t *testing.T) {
	retimer.CreateTimer("test_1", time.Second*4, true)
	retimer.CreateTimer("test_2", time.Second*2, true)
	retimer.OnTimerEvent("test", false, func(timer *timer.ITimer) {
		fmt.Println(timer.Key)
	})
	select {}
}
