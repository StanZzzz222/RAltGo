package alt_events

/*
   Create by zyx
   Date Time: 2024/9/19
   File: event_bus.go
*/

var eventBus = &EventBus{}

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
