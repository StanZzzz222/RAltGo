package utils

import (
	"github.com/gammazero/deque"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/12
   File: task_queue.go
*/

type TaskQueue struct {
	tasks *deque.Deque[func()]
	rw    sync.RWMutex
}

func NewTaskQueue() *TaskQueue {
	q := &TaskQueue{
		tasks: deque.New[func()](),
	}
	return q
}

func (q *TaskQueue) AddTask(task func()) {
	q.rw.Lock()
	defer q.rw.Unlock()
	q.tasks.PushBack(task)
}

func (q *TaskQueue) PopCheck() bool {
	q.rw.RLock()
	defer q.rw.RUnlock()
	if q.tasks.Len() == 0 {
		return false
	}
	return true
}

func (q *TaskQueue) Pop() func() {
	q.rw.Lock()
	defer q.rw.Unlock()
	return q.tasks.PopFront()
}
