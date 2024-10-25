package alt_events

import (
	"github.com/StanZzzz222/RAltGo/common/core/scheduler"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"github.com/StanZzzz222/RAltGo/logger"
	"github.com/goccy/go-json"
	"reflect"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/19
   File: event_bus.go
*/

var eventBus = &EventBus{
	onClientEvents: &sync.Map{},
	onLocalEvents:  &sync.Map{},
}

type EventBus struct {
	onStarts                 []OnStartCallback
	onStops                  []OnStopCallback
	onServerStarteds         []OnServerStartedCallback
	onPlayerConnects         []OnPlayerConnectCallback
	onPlayerDisconnects      []OnPlayerDisconnectCallback
	onConsoleCommands        []OnConsoleCommandCallback
	onNetOwnerChanges        []OnNetOwnerChangeCallback
	onChangeVehicleSeats     []OnChangeVehicleSeatCallback
	onPlayerSpawns           []OnPlayerSpawnCallback
	onPlayerInteriorChanges  []OnPlayerInteriorChangeCallback
	onPlayerDimensionChanges []OnPlayerDimensionChangeCallback
	onEnteringVehicles       []OnEnteringVehicleCallback
	onEnterVehicles          []OnEnterVehicleCallback
	onLeaveVehicles          []OnLeaveVehicleCallback
	onEnterColshapes         []OnEnterColshapeCallback
	onLeaveColshapes         []OnLeaveColshapeCallback
	onChatMessages           []OnChatMessageCallback
	onCommandErrors          []OnCommandErrorCallback
	onPlayerDeaths           []OnPlayerDeathCallback
	onPlayerDamages          []OnPlayerDamageCallback
	onPlayerWeaponChanges    []OnPlayerWeaponChangeCallback
	onPlayerConnectDenieds   []OnPlayerConnectDeniedCallback
	onPlayerHeals            []OnPlayerHealCallback
	onPlayerRequestControls  []OnPlayerRequestControlCallback
	onVehicleAttaches        []OnVehicleAttachCallback
	onVehicleDetaches        []OnVehicleDetachCallback
	onVehicleDestroys        []OnVehicleDestroyCallback
	onVehicleDamages         []OnVehicleDamageCallback
	onVehicleHorns           []OnVehicleHornCallback
	onVehicleSirens          []OnVehicleSirenCallback
	onClientEvents           *sync.Map
	onLocalEvents            *sync.Map
}

func Events() *EventBus {
	return eventBus
}

func (bus *EventBus) OnStart(callback OnStartCallback) {
	bus.onStarts = append(bus.onStarts, callback)
}

func (bus *EventBus) OnStop(callback OnStopCallback) {
	bus.onStops = append(bus.onStops, callback)
}

func (bus *EventBus) OnServerStarted(callback OnServerStartedCallback) {
	bus.onServerStarteds = append(bus.onServerStarteds, callback)
}

func (bus *EventBus) OnPlayerConnect(callback OnPlayerConnectCallback) {
	bus.onPlayerConnects = append(bus.onPlayerConnects, callback)
}

func (bus *EventBus) OnPlayerDisconnect(callback OnPlayerDisconnectCallback) {
	bus.onPlayerDisconnects = append(bus.onPlayerDisconnects, callback)
}

func (bus *EventBus) OnEnterVehicle(callback OnEnterVehicleCallback) {
	bus.onEnterVehicles = append(bus.onEnterVehicles, callback)
}

func (bus *EventBus) OnEnteringVehicle(callback OnEnteringVehicleCallback) {
	bus.onEnteringVehicles = append(bus.onEnteringVehicles, callback)
}

func (bus *EventBus) OnLeaveVehicle(callback OnLeaveVehicleCallback) {
	bus.onLeaveVehicles = append(bus.onLeaveVehicles, callback)
}

func (bus *EventBus) OnEnterColshape(callback OnEnterColshapeCallback) {
	bus.onEnterColshapes = append(bus.onEnterColshapes, callback)
}

func (bus *EventBus) OnLeaveColshape(callback OnLeaveColshapeCallback) {
	bus.onLeaveColshapes = append(bus.onLeaveColshapes, callback)
}

func (bus *EventBus) OnChatMessage(callback OnChatMessageCallback) {
	bus.onChatMessages = append(bus.onChatMessages, callback)
}

func (bus *EventBus) OnCommandError(callback OnCommandErrorCallback) {
	bus.onCommandErrors = append(bus.onCommandErrors, callback)
}

func (bus *EventBus) OnConsoleCommand(callback OnConsoleCommandCallback) {
	bus.onConsoleCommands = append(bus.onConsoleCommands, callback)
}

func (bus *EventBus) OnNetOwnerChange(callback OnNetOwnerChangeCallback) {
	bus.onNetOwnerChanges = append(bus.onNetOwnerChanges, callback)
}

func (bus *EventBus) OnChangeVehicleSeat(callback OnChangeVehicleSeatCallback) {
	bus.onChangeVehicleSeats = append(bus.onChangeVehicleSeats, callback)
}

func (bus *EventBus) OnPlayerSpawn(callback OnPlayerSpawnCallback) {
	bus.onPlayerSpawns = append(bus.onPlayerSpawns, callback)
}

func (bus *EventBus) OnPlayerInteriorChange(callback OnPlayerInteriorChangeCallback) {
	bus.onPlayerInteriorChanges = append(bus.onPlayerInteriorChanges, callback)
}

func (bus *EventBus) OnPlayerDimensionChange(callback OnPlayerDimensionChangeCallback) {
	bus.onPlayerDimensionChanges = append(bus.onPlayerDimensionChanges, callback)
}

func (bus *EventBus) OnPlayerDeath(callback OnPlayerDeathCallback) {
	bus.onPlayerDeaths = append(bus.onPlayerDeaths, callback)
}

func (bus *EventBus) OnPlayerDamage(callback OnPlayerDamageCallback) {
	bus.onPlayerDamages = append(bus.onPlayerDamages, callback)
}

func (bus *EventBus) OnPlayerWeaponChange(callback OnPlayerWeaponChangeCallback) {
	bus.onPlayerWeaponChanges = append(bus.onPlayerWeaponChanges, callback)
}

func (bus *EventBus) OnPlayerConnectDenied(callback OnPlayerConnectDeniedCallback) {
	bus.onPlayerConnectDenieds = append(bus.onPlayerConnectDenieds, callback)
}

func (bus *EventBus) OnPlayerHeal(callback OnPlayerHealCallback) {
	bus.onPlayerHeals = append(bus.onPlayerHeals, callback)
}

func (bus *EventBus) OnPlayerRequestControl(callback OnPlayerRequestControlCallback) {
	bus.onPlayerRequestControls = append(bus.onPlayerRequestControls, callback)
}

func (bus *EventBus) OnVehicleAttach(callback OnVehicleAttachCallback) {
	bus.onVehicleAttaches = append(bus.onVehicleAttaches, callback)
}

func (bus *EventBus) OnVehicleDetach(callback OnVehicleDetachCallback) {
	bus.onVehicleDetaches = append(bus.onVehicleDetaches, callback)
}

func (bus *EventBus) OnVehicleDestroy(callback OnVehicleDestroyCallback) {
	bus.onVehicleDestroys = append(bus.onVehicleDestroys, callback)
}

func (bus *EventBus) OnVehicleDamage(callback OnVehicleDamageCallback) {
	bus.onVehicleDamages = append(bus.onVehicleDamages, callback)
}

func (bus *EventBus) OnVehicleHorn(callback OnVehicleHornCallback) {
	bus.onVehicleHorns = append(bus.onVehicleHorns, callback)
}

func (bus *EventBus) OnClientEvent(eventName string, callback any) {
	var w = lib.GetWarpper()
	t := reflect.TypeOf(callback)
	if t.Kind() == reflect.Func {
		if !checkZeroEventArgs(callback) {
			logger.Logger().LogError("OnClientEvent: should not be zero parameters")
			return
		}
		if !checkFirstEventArgs(callback) {
			logger.Logger().LogError("OnClientEvent: The first parameter should be *timer.IPlayer")
			return
		}
		bus.onClientEvents.Store(eventName, callback)
		s := scheduler.NewScheduler()
		s.AddTask(func() {
			data := dumpEventArgs(callback)
			w.OnClientEvent(eventName, string(data))
		})
		s.Run()
		return
	}
	logger.Logger().LogErrorf("OnClientEvent: unknown callback type: %v", t.Name())
}

func (bus *EventBus) OnLocalEvent(eventName string, callback any) {
	t := reflect.TypeOf(callback)
	if t.Kind() == reflect.Func {
		bus.onLocalEvents.Store(eventName, callback)
		return
	}
	logger.Logger().LogErrorf("OnLocalEvent: unknown callback type: %v", t.Name())
}

func checkZeroEventArgs(callback any) bool {
	var callbackType = reflect.TypeOf(callback)
	var count = callbackType.NumIn()
	return count != 0
}

func checkFirstEventArgs(callback any) bool {
	var callbackType = reflect.TypeOf(callback)
	var firstParam = callbackType.In(0)
	if firstParam.Kind() == reflect.Ptr {
		elemType := firstParam.Elem()
		if elemType == reflect.TypeOf((*models.IPlayer)(nil)).Elem() {
			return true
		}
	}
	return false
}

func dumpEventArgs(callback any) []byte {
	var obj []any
	var callbackType = reflect.TypeOf(callback)
	var params = make([]reflect.Type, 0)
	for i := 1; i < callbackType.NumIn(); i++ {
		params = append(params, callbackType.In(i))
	}
	for _, argType := range params {
		switch argType.Kind() {
		case reflect.Ptr:
			switch argType.Elem() {
			case reflect.TypeOf((*models.IPlayer)(nil)).Elem():
				obj = append(obj, "altv::PlayerContainer")
				continue
			case reflect.TypeOf((*models.IVehicle)(nil)).Elem():
				obj = append(obj, "altv::VehicleContainer")
				continue
			case reflect.TypeOf((*models.IBlip)(nil)).Elem():
				obj = append(obj, "altv::BlipContainer")
				continue
			case reflect.TypeOf((*models.IPed)(nil)).Elem():
				obj = append(obj, "altv::PedContainer")
				continue
			case reflect.TypeOf((*models.IColshape)(nil)).Elem():
				obj = append(obj, "altv::ColshapeContainer")
				continue
			case reflect.TypeOf((*models.Vector3)(nil)).Elem():
				obj = append(obj, "altv::Vector3")
				continue
			case reflect.TypeOf((*models.Rgba)(nil)).Elem():
				obj = append(obj, "altv::Rgba")
				continue
			}
		case reflect.Bool:
			obj = append(obj, "bool")
			continue
		case reflect.Invalid, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			obj = append(obj, "i64")
			continue
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			obj = append(obj, "u64")
			continue
		case reflect.Float32, reflect.Float64:
			obj = append(obj, "f64")
			continue
		case reflect.String:
			obj = append(obj, "String")
			continue
		default:
			obj = append(obj, "String")
		}
	}
	dumpBytes, err := json.Marshal(&obj)
	if err != nil {
		logger.Logger().LogErrorf("Dump event args falied, %v", err.Error())
		return []byte("")
	}
	return dumpBytes
}
