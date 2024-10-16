package timer

import (
	"sync"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/10/16
   File: retimer.go
*/

type ITimer struct {
	Key        string    `json:"key"`
	NotifyUnix int64     `json:"notify_unix"`
	Duration   int64     `json:"duration"`
	Expr       string    `json:"expr"`
	Loop       bool      `json:"loop"`
	LoopCount  int       `json:"loop_count"`
	IsPause    bool      `json:"stop"`
	Datas      *sync.Map `json:"datas"`
}

func NewTimer(key string, duration int64, loop bool, loopCount int) *ITimer {
	return &ITimer{
		Key:        key,
		NotifyUnix: time.Now().Add(time.Duration(duration)).Unix(),
		Duration:   duration,
		Expr:       "",
		Loop:       loop,
		IsPause:    false,
		LoopCount:  loopCount,
		Datas:      &sync.Map{},
	}
}

func (t *ITimer) SetData(key string, value any) {
	t.Datas.Store(key, value)
}

func (t *ITimer) DelData(key string) {
	_, ok := t.Datas.Load(key)
	if ok {
		t.Datas.Delete(key)
	}
}

func (t *ITimer) DelAllData() {
	t.Datas.Range(func(key, value any) bool {
		t.Datas.Delete(key)
		return true
	})
}

func (t *ITimer) HasData(key string) bool {
	_, ok := t.Datas.Load(key)
	if ok {
		return true
	}
	return false
}

func (t *ITimer) GetData(key string) any {
	value, ok := t.Datas.Load(key)
	if ok {
		return value
	}
	return value
}

func (t *ITimer) GetDatas() []any {
	var datas []any
	t.Datas.Range(func(key, value any) bool {
		datas = append(datas, key)
		return true
	})
	return datas
}
