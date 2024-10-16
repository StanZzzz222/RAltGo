package entities

import "time"

/*
   Create by zyx
   Date Time: 2024/10/16
   File: retimer.go
*/

type Timer struct {
	Key               string `json:"key"`
	NotifyMillisecond int64  `json:"notice_millisecond"`
	Millisecond       int64  `json:"millisecond"`
	Message           string `json:"message"`
	Expr              string `json:"expr"`
	Loop              bool   `json:"loop"`
	LoopCount         int    `json:"loop_count"`
	IsPause           bool   `json:"stop"`
}

func NewTimer(key string, milliSecond int64, loop bool, loopCount int) *Timer {
	return &Timer{
		Key:               key,
		NotifyMillisecond: time.Now().Add(time.Duration(milliSecond) * time.Millisecond).UTC().Unix(),
		Millisecond:       milliSecond,
		Message:           "",
		Expr:              "",
		Loop:              loop,
		IsPause:           false,
		LoopCount:         loopCount,
	}
}

func (t *Timer) SetMessage(message string) {
	t.Message = message
}

func (t *Timer) SetExpr(expr string) {
	t.Expr = expr
}

func (t *Timer) SetLoop(loop bool) {
	t.Loop = loop
}

func (t *Timer) SetLoopCount(loopCount int) {
	t.LoopCount = loopCount
}

func (t *Timer) Pause() {
	t.IsPause = true
}

func (t *Timer) Play() {
	t.IsPause = false
}
