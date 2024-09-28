package scheduler

import (
	"github.com/StanZzzz222/RAltGo/common/alt"
	"sync"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/9/12
   File: scheduler.go
*/

type Scheduler struct {
	tasks []func()
	rw    sync.RWMutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: []func(){},
	}
}

func (s *Scheduler) AddTask(task func()) {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) Run() {
	alt.SetTimeout(time.Microsecond*10, func() {
		for _, task := range s.tasks {
			task()
		}
		s.tasks = []func(){}
	})
}
