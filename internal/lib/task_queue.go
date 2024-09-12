package lib

import (
	"container/list"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/12
   File: task_queue.go
*/

type TaskQueue struct {
	tasks *list.List
	rw    sync.RWMutex
}

func NewTaskQueue() *TaskQueue {
	q := &TaskQueue{
		tasks: list.New(),
	}
	return q
}

func (q *TaskQueue) AddTask(task func()) {
	q.rw.Lock()
	defer q.rw.Unlock()
	q.tasks.PushBack(task)
}

func (q *TaskQueue) GetTasks() []func() {
	q.rw.RLock()
	defer q.rw.RUnlock()

	var tasks []func()
	for e := q.tasks.Front(); e != nil; e = e.Next() {
		tasks = append(tasks, e.Value.(func()))
	}
	q.tasks.Init()
	return tasks
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
	elem := q.tasks.Front()
	q.tasks.Remove(elem)
	return elem.Value.(func())
}
