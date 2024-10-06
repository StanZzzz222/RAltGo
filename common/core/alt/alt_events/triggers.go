package alt_events

import (
	"github.com/StanZzzz222/RAltGo/common"
	"github.com/StanZzzz222/RAltGo/common/core/pools"
	"github.com/StanZzzz222/RAltGo/common/core/scheduler"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_entity_type"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enums"
	"github.com/StanZzzz222/RAltGo/logger"
	"github.com/goccy/go-json"
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
	EmitAllPlayer("chat:message", player.GetGameName(), message)
	if eventBus.onChatMessage != nil {
		eventBus.onChatMessage(player, message)
	}
}

func (t *EventBusTrigger) TriggerOnEnteringVehicle(player *models.IPlayer, vehicle *models.IVehicle, seat uint8) {
	if eventBus.onEnteringVehicle != nil {
		eventBus.onEnteringVehicle(player, vehicle, seat)
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

func (t *EventBusTrigger) TriggerOnCommandError(player *models.IPlayer, existCommand bool, commandName, desc string) {
	if eventBus.onCommandError != nil {
		eventBus.onCommandError(player, existCommand, commandName, desc)
	}
}

func (t *EventBusTrigger) TriggerOnConsoleCommand(name string, args []string) {
	if eventBus.onConsoleCommand != nil {
		eventBus.onConsoleCommand(name, args)
	}
}

func (t *EventBusTrigger) TriggerOnNetOwnerChange(entity any, oldNetOwner *models.IPlayer, newNetOwner *models.IPlayer) {
	if eventBus.onNetOwnerChange != nil {
		eventBus.onNetOwnerChange(entity, oldNetOwner, newNetOwner)
	}
}

func (t *EventBusTrigger) TriggerOnChangeVehicleSeat(player *models.IPlayer, vehicle *models.IVehicle, oldSeat, newSeat uint8) {
	if eventBus.onChangeVehicleSeat != nil {
		eventBus.onChangeVehicleSeat(player, vehicle, oldSeat, newSeat)
	}
}

func (t *EventBusTrigger) TriggerOnPlayerSpawn(player *models.IPlayer) {
	if eventBus.onPlayerSpawn != nil {
		eventBus.onPlayerSpawn(player)
	}
}

func (t *EventBusTrigger) TriggerOnPlayerInteriorChange(player *models.IPlayer, oldInterior, newInterior uint32) {
	if eventBus.onPlayerInteriorChange != nil {
		eventBus.onPlayerInteriorChange(player, oldInterior, newInterior)
	}
}

func (t *EventBusTrigger) TriggerOnPlayerDimensionChange(player *models.IPlayer, oldDimension, newDimension int32) {
	if eventBus.onPlayerDimensionChange != nil {
		eventBus.onPlayerDimensionChange(player, oldDimension, newDimension)
	}
}

func (t *EventBusTrigger) TriggerOnClientEvent(player *models.IPlayer, eventName, eventArgs string) {
	if callback, ok := eventBus.onClientEvents.Load(eventName); ok {
		callbackValue := reflect.ValueOf(callback)
		callbackType := reflect.TypeOf(callback)
		args := t.EventArgsParse(eventArgs)
		if callbackType.NumIn() == len(args)+1 {
			inputs := make([]reflect.Value, 0)
			inputs = append(inputs, reflect.ValueOf(player))
			if len(args) == 0 {
				callbackValue.Call(inputs)
				return
			}
			for _, arg := range args {
				inputs = append(inputs, reflect.ValueOf(arg))
			}
			callbackValue.Call(inputs)
		} else {
			logger.Logger().LogErrorf("ClientEvent trigger falied, expected: %v | received: %v", callbackType.NumIn(), len(args))
		}
	}
}

func (t *EventBusTrigger) TriggerOnLocalEvent(eventName, eventArgs string) {
	if callback, ok := eventBus.onLocalEvents.Load(eventName); ok {
		s := scheduler.NewScheduler()
		s.AddTask(func() {
			callbackValue := reflect.ValueOf(callback)
			args := t.EventArgsParse(eventArgs)
			inputs := make([]reflect.Value, 0)
			if len(args) == 0 {
				callbackValue.Call(inputs)
				return
			}
			for _, arg := range args {
				inputs = append(inputs, reflect.ValueOf(arg))
			}
			callbackValue.Call(inputs)
		})
		s.Run()
	}
}

func (t *EventBusTrigger) EventArgsParse(eventArgs string) []any {
	var args []any
	var result []any
	err := json.Unmarshal([]byte(eventArgs), &args)
	if err != nil {
		logger.Logger().LogErrorf("EventArgsParse falied, %v", err.Error())
		return args
	}
	if len(args) != 0 {
		for _, objAny := range args {
			obj := objAny.(map[string]any)
			argType := obj["type"]
			argValue := obj["value"]
			t := reflect.TypeOf(argType)
			if t.Kind() == reflect.Float64 {
				switch enums.ObjectType(int32(argType.(float64))) {
				case enums.Player:
					p := pools.GetPlayer(uint32(argValue.(float64)))
					result = append(result, p)
					continue
				case enums.Ped:
					p := pools.GetPed(uint32(argValue.(float64)))
					result = append(result, p)
					continue
				case enums.Vehicle:
					v := pools.GetVehicle(uint32(argValue.(float64)))
					result = append(result, v)
					continue
				case enums.Colshape:
					c := pools.GetColshape(uint32(argValue.(float64)))
					result = append(result, c)
					continue
				case enums.Object:
					o := pools.GetObject(uint32(argValue.(float64)))
					result = append(result, o)
					continue
				case enums.CheckPoint:
					c := pools.GetCheckpoint(uint32(argValue.(float64)))
					result = append(result, c)
					continue
				case enums.Marker:
					m := pools.GetMarker(uint32(argValue.(float64)))
					result = append(result, m)
					continue
				case enums.Blip:
					b := pools.GetBlip(uint32(argValue.(float64)))
					result = append(result, b)
					continue
				default:
					logger.Logger().LogErrorf("EventArgsParse falied, Unknow ObjectType: %v", argType)
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
				case reflect.Float64.String():
					result = append(result, argValue.(float64))
					continue
				case reflect.String.String():
					switch argValue.(string) {
					case "rgba":
						var rgba *entities.Rgba
						_ = json.Unmarshal([]byte(argValue.(string)), &rgba)
						result = append(result, rgba)
						continue
					case "vector2":
						var vec3 *entities.Vector3
						_ = json.Unmarshal([]byte(argValue.(string)), &vec3)
						result = append(result, vec3)
						continue
					case "vector3":
						var vec2Map = map[string]any{}
						_ = json.Unmarshal([]byte(argValue.(string)), &vec2Map)
						result = append(result, common.NewVector3(float32(vec2Map["x"].(float64)), float32(vec2Map["y"].(float64)), 0))
						continue
					}
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
	return args
}
