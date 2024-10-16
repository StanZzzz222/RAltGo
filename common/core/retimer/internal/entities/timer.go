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

func (t *Timer) GetKey() string {
	return t.Key
}
func (t *Timer) GetNextNotifyUnix() int64 {
	return t.NotifyMillisecond
}
func (t *Timer) GetMillisecond() int64 {
	return t.Millisecond
}
func (t *Timer) GetMessage() string {
	return t.Message
}
func (t *Timer) GetExpr() string {
	return t.Expr
}
func (t *Timer) GetLoopCount() int {
	return t.LoopCount
}
func (t *Timer) IsLoop() bool {
	return t.Loop
}
func (t *Timer) IsStop() bool {
	return t.IsPause
}

func (t *Timer) SetMessage(message string) {
	t.Message = message
}

func (t *Timer) SetExpr(expr string) {
	t.Expr = expr
}

func (t *Timer) Pause() {
	t.IsStop = true
}

func (t *Timer) Play() {
	t.IsStop = false
}
