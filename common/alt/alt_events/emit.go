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
	s := scheduler.NewScheduler()
	for _, player := range players {
		s.AddTask(func() {
			player.Emit(eventName, args...)
		})
	}
	s.Run()
}

func Emit(player *models.IPlayer, eventName string, args ...any) {
	player.Emit(eventName, args...)
}
