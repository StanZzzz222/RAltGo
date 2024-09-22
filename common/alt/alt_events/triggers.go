package alt_events

import (
	"encoding/json"
	"github.com/StanZzzz222/RAltGo/common/alt/pools"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_entity_type"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"github.com/StanZzzz222/RAltGo/logger"
	"reflect"
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

func (t *EventBusTrigger) TriggerOnChatMessage(player *models.IPlayer, message string) {
	if len(player.GetChatName()) <= 0 {
		EmitAllPlayer("chat:message", player.GetName(), message)
	} else {
		EmitAllPlayer("chat:message", player.GetChatName(), message)
	}
	if eventBus.onChatMessage != nil {
		eventBus.onChatMessage(player, message)
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

func (t *EventBusTrigger) TriggerOnClientEvent(player *models.IPlayer, eventName, eventArgs string) {
	if callback, ok := eventBus.onClientEvents.Load(eventName); ok {
		callbackValue := reflect.ValueOf(callback)
		args := t.EventArgsParse(eventArgs)
		inputs := make([]reflect.Value, 0)
		inputs = append(inputs, reflect.ValueOf(player))
		for _, arg := range args {
			inputs = append(inputs, reflect.ValueOf(arg))
		}
		callbackValue.Call(inputs)
	}
}

func (t *EventBusTrigger) EventArgsParse(eventArgs string) []any {
	var args []any
	var result []any
	err := json.Unmarshal([]byte(eventArgs), &args)
	if err != nil {
		logger.LogErrorf("EventArgsParse falied, %v", err.Error())
		return args
	}
	for _, objAny := range args {
		obj := objAny.(map[string]any)
		argType := obj["type"]
		argValue := obj["value"]
		t := reflect.TypeOf(argType)
		if t.Kind() == reflect.Float64 {
			switch enum.ObjectType(int32(argType.(float64))) {
			case enum.Player:
				p := pools.GetPlayer(uint32(argValue.(float64)))
				result = append(result, p)
				continue
			case enum.Ped:
				p := pools.GetPed(uint32(argValue.(float64)))
				result = append(result, p)
				continue
			case enum.Vehicle:
				v := pools.GetVehicle(uint32(argValue.(float64)))
				result = append(result, v)
				continue
			case enum.Colshape:
				c := pools.GetColshape(uint32(argValue.(float64)))
				result = append(result, c)
				continue
			case enum.Blip:
				b := pools.GetBlip(uint32(argValue.(float64)))
				result = append(result, b)
				continue
			default:
				logger.LogErrorf("EventArgsParse falied, Unknow ObjectType: %v", argType)
				return result
			}
		} else {
			switch argType.(string) {
			case reflect.Bool.String():
				result = append(result, argValue.(bool))
				continue
			case reflect.Int.String():
				result = append(result, int64(argValue.(float64)))
				continue
			case reflect.Uint64.String():
				result = append(result, uint64(argValue.(float64)))
				continue
			case reflect.Uint64.String():
				result = append(result, uint64(argValue.(float64)))
				continue
			case reflect.Float64.String():
				result = append(result, argValue.(float64))
				continue
			case reflect.String.String():
				result = append(result, argValue.(string))
				continue
			case "null":
				result = append(result, nil)
				continue
			}
		}
	}
	return result
}
