package alt_events

import (
	"github.com/StanZzzz222/RAltGo/common"
	"github.com/StanZzzz222/RAltGo/common/core/pools"
	"github.com/StanZzzz222/RAltGo/common/core/scheduler"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_entity_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/denied_reason_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/weapon_hash"
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
	if len(eventBus.onStarts) > 0 {
		for _, onStart := range eventBus.onStarts {
			onStart()
		}
	}
}

func (t *EventBusTrigger) TriggerOnStop() {
	if len(eventBus.onStops) > 0 {
		for _, onStop := range eventBus.onStops {
			onStop()
		}
	}
}

func (t *EventBusTrigger) TriggerOnServerStarted() {
	if len(eventBus.onServerStarteds) > 0 {
		for _, onServerStarted := range eventBus.onServerStarteds {
			onServerStarted()
		}
	}
}

func (t *EventBusTrigger) TriggerOnPlayerConnect(player *models.IPlayer) {
	if len(eventBus.onPlayerConnects) > 0 {
		for _, onPlayerConnect := range eventBus.onPlayerConnects {
			onPlayerConnect(player)
		}
	}
}

func (t *EventBusTrigger) TriggerOnPlayerDisconnect(player *models.IPlayer, reason string) {
	if len(eventBus.onPlayerDisconnects) > 0 {
		for _, onPlayerDisconnect := range eventBus.onPlayerDisconnects {
			onPlayerDisconnect(player, reason)
		}
	}
}

func (t *EventBusTrigger) TriggerOnChatMessage(player *models.IPlayer, message string) {
	EmitAllPlayer("chat:message", player.GetGameName(), message)
	if len(eventBus.onChatMessages) > 0 {
		for _, onChatMessage := range eventBus.onChatMessages {
			onChatMessage(player, message)
		}
	}
}

func (t *EventBusTrigger) TriggerOnEnteringVehicle(player *models.IPlayer, vehicle *models.IVehicle, seat uint8) {
	if len(eventBus.onEnteringVehicles) > 0 {
		for _, onEnteringVehicle := range eventBus.onEnteringVehicles {
			onEnteringVehicle(player, vehicle, seat)
		}
	}
}

func (t *EventBusTrigger) TriggerOnEnterVehicle(player *models.IPlayer, vehicle *models.IVehicle, seat uint8) {
	if len(eventBus.onEnterVehicles) > 0 {
		for _, onEnterVehicle := range eventBus.onEnterVehicles {
			onEnterVehicle(player, vehicle, seat)
		}
	}
}

func (t *EventBusTrigger) TriggerOnLeaveVehicle(player *models.IPlayer, vehicle *models.IVehicle, seat uint8) {
	if len(eventBus.onLeaveVehicles) > 0 {
		for _, onLeaveVehicle := range eventBus.onLeaveVehicles {
			onLeaveVehicle(player, vehicle, seat)
		}
	}
}

func (t *EventBusTrigger) TriggerOnEnterColshape(colshapeEntityType colshape_entity_type.ColshapeEntityType, player *models.IPlayer, vehicle *models.IVehicle, colshape *models.IColshape) {
	if len(eventBus.onEnterColshapes) > 0 {
		for _, onEnterColshape := range eventBus.onEnterColshapes {
			onEnterColshape(colshapeEntityType, player, vehicle, colshape)
		}
	}
}

func (t *EventBusTrigger) TriggerOnLeaveColshape(colshapeEntityType colshape_entity_type.ColshapeEntityType, player *models.IPlayer, vehicle *models.IVehicle, colshape *models.IColshape) {
	if len(eventBus.onLeaveColshapes) > 0 {
		for _, onLeaveColshape := range eventBus.onLeaveColshapes {
			onLeaveColshape(colshapeEntityType, player, vehicle, colshape)
		}
	}
}

func (t *EventBusTrigger) TriggerOnCommandError(player *models.IPlayer, existCommand bool, commandName, desc string) {
	if len(eventBus.onCommandErrors) > 0 {
		for _, onCommandError := range eventBus.onCommandErrors {
			onCommandError(player, existCommand, commandName, desc)
		}
	}
}

func (t *EventBusTrigger) TriggerOnConsoleCommand(name string, args []string) {
	if len(eventBus.onConsoleCommands) > 0 {
		for _, onConsoleCommand := range eventBus.onConsoleCommands {
			onConsoleCommand(name, args)
		}
	}
}

func (t *EventBusTrigger) TriggerOnNetOwnerChange(entity any, oldNetOwner *models.IPlayer, newNetOwner *models.IPlayer) {
	if len(eventBus.onNetOwnerChanges) > 0 {
		for _, onNetOwnerChange := range eventBus.onNetOwnerChanges {
			onNetOwnerChange(entity, oldNetOwner, newNetOwner)
		}
	}
}

func (t *EventBusTrigger) TriggerOnChangeVehicleSeat(player *models.IPlayer, vehicle *models.IVehicle, oldSeat, newSeat uint8) {
	if len(eventBus.onChangeVehicleSeats) > 0 {
		for _, onChangeVehicleSeat := range eventBus.onChangeVehicleSeats {
			onChangeVehicleSeat(player, vehicle, oldSeat, newSeat)
		}
	}
}

func (t *EventBusTrigger) TriggerOnPlayerSpawn(player *models.IPlayer) {
	if len(eventBus.onPlayerSpawns) > 0 {
		for _, onPlayerSpawn := range eventBus.onPlayerSpawns {
			onPlayerSpawn(player)
		}
	}
}

func (t *EventBusTrigger) TriggerOnPlayerInteriorChange(player *models.IPlayer, oldInterior, newInterior uint32) {
	if len(eventBus.onPlayerInteriorChanges) > 0 {
		for _, onPlayerInteriorChange := range eventBus.onPlayerInteriorChanges {
			onPlayerInteriorChange(player, oldInterior, newInterior)
		}
	}
}

func (t *EventBusTrigger) TriggerOnPlayerDimensionChange(player *models.IPlayer, oldDimension, newDimension int32) {
	if len(eventBus.onPlayerDimensionChanges) > 0 {
		for _, onPlayerDimensionChange := range eventBus.onPlayerDimensionChanges {
			onPlayerDimensionChange(player, oldDimension, newDimension)
		}
	}
}

func (t *EventBusTrigger) TriggerOnPlayerDeath(player *models.IPlayer, killer any, weaponHash weapon_hash.ModelHash) {
	if len(eventBus.onPlayerDeaths) > 0 {
		for _, onPlayerDeath := range eventBus.onPlayerDeaths {
			onPlayerDeath(player, killer, weaponHash)
		}
	}
}

func (t *EventBusTrigger) TriggerOnPlayerDamage(player *models.IPlayer, attacker any, healthDamage, armourDamage uint16) {
	if len(eventBus.onPlayerDamages) > 0 {
		for _, onPlayerDamage := range eventBus.onPlayerDamages {
			onPlayerDamage(player, attacker, healthDamage, armourDamage)
		}
	}
}

func (t *EventBusTrigger) TriggerOnPlayerWeaponChange(player *models.IPlayer, oldWeaponHash, newWeaponHash uint32) {
	if len(eventBus.onPlayerWeaponChanges) > 0 {
		for _, onPlayerWeaponChange := range eventBus.onPlayerWeaponChanges {
			onPlayerWeaponChange(player, weapon_hash.ModelHash(oldWeaponHash), weapon_hash.ModelHash(newWeaponHash))
		}
	}
}

func (t *EventBusTrigger) TriggerOnPlayerConnectDenied(reason denied_reason_type.DeniedReason, name, ip string, passwordHash uint64, isDebug bool, branch string, majorVersion, minorVersion uint16, cdnUrl string, discordId int64) {
	if len(eventBus.onPlayerConnectDenieds) > 0 {
		for _, onPlayerConnectDenied := range eventBus.onPlayerConnectDenieds {
			onPlayerConnectDenied(reason, name, ip, passwordHash, isDebug, branch, majorVersion, minorVersion, cdnUrl, discordId)
		}
	}
}

func (t *EventBusTrigger) TriggerOnPlayerHeal(player *models.IPlayer, oldHealth, newHealth, oldArmour, newArmour uint16) {
	if len(eventBus.onPlayerHeals) > 0 {
		for _, onPlayerHeal := range eventBus.onPlayerHeals {
			onPlayerHeal(player, oldHealth, newHealth, oldArmour, newArmour)
		}
	}
}

func (t *EventBusTrigger) TriggerOnPlayerRequestControl(player *models.IPlayer, entity any) {
	if len(eventBus.onPlayerRequestControls) > 0 {
		for _, onPlayerRequestControl := range eventBus.onPlayerRequestControls {
			onPlayerRequestControl(player, entity)
		}
	}
}

func (t *EventBusTrigger) TriggerOnVehicleAttach(vehicle *models.IVehicle, attached *models.IVehicle) {
	if len(eventBus.onVehicleAttaches) > 0 {
		for _, onVehicleAttach := range eventBus.onVehicleAttaches {
			onVehicleAttach(vehicle, attached)
		}
	}
}

func (t *EventBusTrigger) TriggerOnVehicleDetach(vehicle *models.IVehicle, attached *models.IVehicle) {
	if len(eventBus.onVehicleDetaches) > 0 {
		for _, onVehicleDetach := range eventBus.onVehicleDetaches {
			onVehicleDetach(vehicle, attached)
		}
	}
}

func (t *EventBusTrigger) TriggerOnVehicleDestroy(vehicle *models.IVehicle) {
	if len(eventBus.onVehicleDestroys) > 0 {
		for _, onVehicleDestroy := range eventBus.onVehicleDestroys {
			onVehicleDestroy(vehicle)
		}
	}
}

func (t *EventBusTrigger) TriggerOnVehicleDamage(vehicle *models.IVehicle, damager any, bodyHealthDamage, bodyAdditionalHealthDamage, engineHealthDamage, petrolTankHealthDamage uint32, weapon weapon_hash.ModelHash) {
	if len(eventBus.onVehicleDamages) > 0 {
		for _, onVehicleDamage := range eventBus.onVehicleDamages {
			onVehicleDamage(vehicle, damager, bodyHealthDamage, bodyAdditionalHealthDamage, engineHealthDamage, petrolTankHealthDamage, weapon)
		}
	}
}

func (t *EventBusTrigger) TriggerOnVehicleHorn(vehicle *models.IVehicle, player *models.IPlayer, state bool) {
	if len(eventBus.onVehicleHorns) > 0 {
		for _, onVehicleHorn := range eventBus.onVehicleHorns {
			onVehicleHorn(vehicle, player, state)
		}
	}
}

func (t *EventBusTrigger) TriggerOnVehicleSiren(vehicle *models.IVehicle, state bool) {
	if len(eventBus.onVehicleSirens) > 0 {
		for _, onVehicleSiren := range eventBus.onVehicleSirens {
			onVehicleSiren(vehicle, state)
		}
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
			for i, arg := range args {
				if arg == nil {
					inputs = append(inputs, reflect.Zero(callbackType.In(i+1)))
					continue
				}
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
			if objAny == nil {
				result = append(result, nil)
				continue
			}
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
						var rgba *models.Rgba
						_ = json.Unmarshal([]byte(argValue.(string)), &rgba)
						result = append(result, rgba)
						continue
					case "vector2":
						var vec3 *models.Vector3
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
