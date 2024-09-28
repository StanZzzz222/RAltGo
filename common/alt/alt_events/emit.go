package alt_events

import (
	"github.com/StanZzzz222/RAltGo/common/alt/scheduler"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/internal/lib"
)

/*
   Create by zyx
   Date Time: 2024/9/20
   File: emit.go
*/

func EmitLocalEvent(eventName string, args ...any) {
	var mvalues = models.NewMValues(args...)
	Triggers().TriggerOnLocalEvent(eventName, mvalues.Dump())
}

func EmitAllPlayer(eventName string, args ...any) {
	var w = lib.GetWarpper()
	s := scheduler.NewScheduler()
	s.AddTask(func() {
		var mvalues = models.NewMValues(args...)
		w.EmitAllPlayer(eventName, mvalues.Dump())
	})
	s.Run()
}

func EmitSomePlayers(players []*models.IPlayer, eventName string, args ...any) {
	var w = lib.GetWarpper()
	s := scheduler.NewScheduler()
	mvalues := models.NewMValues(args...)
	mvaluesDumps := mvalues.Dump()
	for _, player := range players {
		s.AddTask(func() {
			w.Emit(player.GetId(), eventName, mvaluesDumps)
		})
	}
	s.Run()
}

func Emit(player *models.IPlayer, eventName string, args ...any) {
	player.Emit(eventName, args...)
}
