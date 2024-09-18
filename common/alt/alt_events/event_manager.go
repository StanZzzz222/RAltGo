package alt_events

import "github.com/StanZzzz222/RAltGo/common/models"

/*
   Create by zyx
   Date Time: 2024/9/9
   File: alt_events.go
*/

type Callback struct{}

var onStartCallback OnStartCallback
var onStopCallback OnStopCallback
var onServerStartedCallback OnServerStartedCallback
var onPlayerConnectCallback OnPlayerConnectCallback
var onPlayerDisconnectCallback OnPlayerDisconnectCallback
var onEnterVehicleCallback OnEnterVehicleCallback
var onLeaveVehicleCallback OnLeaveVehicleCallback

type OnStartCallback func()
type OnStopCallback func()
type OnServerStartedCallback func()
type OnPlayerConnectCallback func(player *models.IPlayer)
type OnPlayerDisconnectCallback func(player *models.IPlayer, reason string)
type OnEnterVehicleCallback func(player *models.IPlayer, vehicle *models.IVehicle, seat uint8)
type OnLeaveVehicleCallback func(player *models.IPlayer, vehicle *models.IVehicle, seat uint8)

func (cb *Callback) TriggerOnStart() {
	if onStartCallback != nil {
		onStartCallback()
	}
}
func (cb *Callback) TriggerOnStop() {
	if onStopCallback != nil {
		onStopCallback()
	}
}
func (cb *Callback) TriggerOnServerStarted() {
	if onServerStartedCallback != nil {
		onServerStartedCallback()
	}
}
func (cb *Callback) TriggerOnPlayerConnect(player *models.IPlayer) {
	if onPlayerConnectCallback != nil {
		onPlayerConnectCallback(player)
	}
}
func (cb *Callback) TriggerOnPlayerDisconnect(player *models.IPlayer, reason string) {
	if onPlayerDisconnectCallback != nil {
		onPlayerDisconnectCallback(player, reason)
	}
}
func (cb *Callback) TriggerOnEnterVehicle(player *models.IPlayer, vehicle *models.IVehicle, seat uint8) {
	if onEnterVehicleCallback != nil {
		onEnterVehicleCallback(player, vehicle, seat)
	}
}
func (cb *Callback) TriggerOnLeaveVehicle(player *models.IPlayer, vehicle *models.IVehicle, seat uint8) {
	if onLeaveVehicleCallback != nil {
		onLeaveVehicleCallback(player, vehicle, seat)
	}
}

func OnStart(cb OnStartCallback) {
	if onStartCallback == nil {
		onStartCallback = cb
	}
}

func OnStop(cb OnStopCallback) {
	if onStopCallback == nil {
		onStopCallback = cb
	}
}

func OnServerStarted(cb OnServerStartedCallback) {
	if onServerStartedCallback == nil {
		onServerStartedCallback = cb
	}
}

func OnPlayerConnect(cb OnPlayerConnectCallback) {
	if onPlayerConnectCallback == nil {
		onPlayerConnectCallback = cb
	}
}

func OnPlayerDisconnect(cb OnPlayerDisconnectCallback) {
	if onPlayerDisconnectCallback == nil {
		onPlayerDisconnectCallback = cb
	}
}

func OnEnterVehicle(cb OnEnterVehicleCallback) {
	if onEnterVehicleCallback == nil {
		onEnterVehicleCallback = cb
	}
}

func OnLeaveVehicle(cb OnLeaveVehicleCallback) {
	if onLeaveVehicleCallback == nil {
		onLeaveVehicleCallback = cb
	}
}
