package alt_events

import (
	"encoding/json"
	"github.com/StanZzzz222/RAltGo/common/alt/scheduler"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"github.com/StanZzzz222/RAltGo/logger"
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
}

type EventBus struct {
	onStart            OnStartCallback
	onStop             OnStopCallback
	onServerStarted    OnServerStartedCallback
	onPlayerConnect    OnPlayerConnectCallback
	onPlayerDisconnect OnPlayerDisconnectCallback
	onEnterVehicle     OnEnterVehicleCallback
	onLeaveVehicle     OnLeaveVehicleCallback
	onEnterColshape    OnEnterColshapeCallback
	onLeaveColshape    OnLeaveColshapeCallback
	onClientEvents     *sync.Map
}

func Events() *EventBus {
	return eventBus
}

func (bus *EventBus) OnStart(callback OnStartCallback) {
	bus.onStart = callback
}

func (bus *EventBus) OnStop(callback OnStopCallback) {
	bus.onStop = callback
}

func (bus *EventBus) OnServerStarted(callback OnServerStartedCallback) {
	bus.onServerStarted = callback
}

func (bus *EventBus) OnPlayerConnect(callback OnPlayerConnectCallback) {
	bus.onPlayerConnect = callback
}

func (bus *EventBus) OnPlayerDisconnect(callback OnPlayerDisconnectCallback) {
	bus.onPlayerDisconnect = callback
}

func (bus *EventBus) OnEnterVehicle(callback OnEnterVehicleCallback) {
	bus.onEnterVehicle = callback
}

func (bus *EventBus) OnLeaveVehicle(callback OnLeaveVehicleCallback) {
	bus.onLeaveVehicle = callback
}

func (bus *EventBus) OnEnterColshape(callback OnEnterColshapeCallback) {
	bus.onEnterColshape = callback
}

func (bus *EventBus) OnLeaveColshape(callback OnLeaveColshapeCallback) {
	bus.onLeaveColshape = callback
}

func (bus *EventBus) OnClientEvent(eventName string, callback any) {
	var w = &lib.Warrper{}
	t := reflect.TypeOf(callback)
	if t.Kind() == reflect.Func {
		if !checkZeroEventArgs(callback) {
			logger.LogError("OnClientEvent: should not be zero parameters")
			return
		}
		if !checkFirstEventArgs(callback) {
			logger.LogError("OnClientEvent: The first parameter should be *models.IPlayer")
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
	logger.LogErrorf("OnClientEvent: unknown callback type: %v", t.Name())
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
		logger.LogErrorf("Dump event args falied, %v", err.Error())
		return []byte("")
	}
	return dumpBytes
}
