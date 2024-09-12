package scheduler

import (
	"github.com/StanZzzz222/RAltGo/internal/timers"
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
	mu    *sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: []func(){},
		mu:    &sync.Mutex{},
	}
}

func (s *Scheduler) AddTask(task func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) RunWait() {
	timers.SetTimeout(time.Microsecond*1, func() {
		for _, task := range s.tasks {
			task()
		}
	})
}
