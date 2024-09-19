package alt_events

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_entity_type"
)

/*
   Create by zyx
   Date Time: 2024/9/19
   File: triggers.go
*/

var eventBusTrigger = &EventBusTrigger{}

type EventBusTrigger struct{}

func Triggers() *EventBusTrigger {
	return eventBusTrigger
}

func (t *EventBusTrigger) TriggerOnStart() {
	if eventBus.onStart != nil {
		eventBus.onStart()
	}
}

func (t *EventBusTrigger) TriggerOnStop() {
	if eventBus.onStop != nil {
		eventBus.onStop()
	}
}

func (t *EventBusTrigger) TriggerOnServerStarted() {
	if eventBus.onServerStarted != nil {
		eventBus.onServerStarted()
	}
}

func (t *EventBusTrigger) TriggerOnPlayerConnect(player *models.IPlayer) {
	if eventBus.onPlayerConnect != nil {
		eventBus.onPlayerConnect(player)
	}
}

func (t *EventBusTrigger) TriggerOnPlayerDisconnect(player *models.IPlayer, reason string) {
	if eventBus.onPlayerDisconnect != nil {
		eventBus.onPlayerDisconnect(player, reason)
	}
}

func (t *EventBusTrigger) TriggerOnEnterVehicle(player *models.IPlayer, vehicle *models.IVehicle, seat uint8) {
	if eventBus.onEnterVehicle != nil {
		eventBus.onEnterVehicle(player, vehicle, seat)
	}
}

func (t *EventBusTrigger) TriggerOnLeaveVehicle(player *models.IPlayer, vehicle *models.IVehicle, seat uint8) {
	if eventBus.onLeaveVehicle != nil {
		eventBus.onLeaveVehicle(player, vehicle, seat)
	}
}

func (t *EventBusTrigger) TriggerOnEnterColshape(colshapeEntityType colshape_entity_type.ColshapeEntityType, player *models.IPlayer, vehicle *models.IVehicle, colshape *models.IColshape) {
	if eventBus.onEnterColshape != nil {
		eventBus.onEnterColshape(colshapeEntityType, player, vehicle, colshape)
	}
}

func (t *EventBusTrigger) TriggerOnLeaveColshape(colshapeEntityType colshape_entity_type.ColshapeEntityType, player *models.IPlayer, vehicle *models.IVehicle, colshape *models.IColshape) {
	if eventBus.onLeaveColshape != nil {
		eventBus.onLeaveColshape(colshapeEntityType, player, vehicle, colshape)
	}
}
