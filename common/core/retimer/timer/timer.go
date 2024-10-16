package timer

import "time"

/*
   Create by zyx
   Date Time: 2024/10/16
   File: retimer.go
*/

type ITimer struct {
	Key        string `json:"key"`
	NotifyUnix int64  `json:"notify_unix"`
	Duration   int64  `json:"duration"`
	Message    string `json:"message"`
	Expr       string `json:"expr"`
	Loop       bool   `json:"loop"`
	LoopCount  int    `json:"loop_count"`
	IsPause    bool   `json:"stop"`
}

func NewTimer(key string, duration int64, loop bool, loopCount int) *ITimer {
	return &ITimer{
		Key:        key,
		NotifyUnix: time.Now().Add(time.Duration(duration)).Unix(),
		Duration:   duration,
		Message:    "",
		Expr:       "",
		Loop:       loop,
		IsPause:    false,
		LoopCount:  loopCount,
	}
}
