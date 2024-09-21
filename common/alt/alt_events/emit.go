package alt_events

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/internal/lib"
)

/*
   Create by zyx
   Date Time: 2024/9/20
   File: emit.go
*/

func EmitAllPlayer(eventName string, args ...any) {
	var w = &lib.Warrper{}
	var mvalues = models.NewMValues(args...)
	w.EmitAllPlayer(eventName, mvalues.Dump())
}

func EmitSomePlayers(players []*models.IPlayer, eventName string, args ...any) {
	for _, player := range players {
		player.Emit(eventName, args...)
	}
}

func Emit(player *models.IPlayer, eventName string, args ...any) {
	player.Emit(eventName, args...)
}
