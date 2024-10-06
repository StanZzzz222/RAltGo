package tests

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common"
	"github.com/StanZzzz222/RAltGo/common/core/alt/alt_events"
	"github.com/StanZzzz222/RAltGo/common/models"
	"testing"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: event_test.go
*/

func TestEvent(t *testing.T) {
	var p = &models.IPlayer{}
	p = p.NewIPlayer(1, "test", "127.0.0.1", "test", "test", 123, 1, 1, common.NewVector3(0, 0, 0), common.NewVector3(0, 0, 0))
	alt_events.Events().OnLocalEvent("test", func(name string, player *models.IPlayer) {
		fmt.Println("Name: ", name)
		fmt.Println("PlayerId: ", player.GetId())
	})
	models.GetPools().PutPlayer(p)
	alt_events.EmitLocalEvent("test", "Evan", p)
}
