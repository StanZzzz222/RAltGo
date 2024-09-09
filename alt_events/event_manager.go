package alt_events

import (
	"gamemode/models"
)

/*
   Create by zyx
   Date Time: 2024/9/9
   File: alt_events.go
*/

type Callback struct{}

type OnStartCallback func()
type OnStopCallback func()
type OnPlayerConnectCallback func(player *models.IPlayer)
type OnEnterVehicleCallback func(player *models.IPlayer, vehicle *models.IVehicle, seat uint8)

var onStartCallback OnStartCallback
var onStopCallback OnStopCallback
var onPlayerConnectCallback OnPlayerConnectCallback
var onEnterVehicleCallback OnEnterVehicleCallback

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

func OnPlayerConnect(cb OnPlayerConnectCallback) {
	if onPlayerConnectCallback == nil {
		onPlayerConnectCallback = cb
	}
}

func OnEnterVehicle(cb OnEnterVehicleCallback) {
	if onEnterVehicleCallback == nil {
		onEnterVehicleCallback = cb
	}
}

func (cb *Callback) New() *Callback                                { return &Callback{} }
func (cb *Callback) TriggerOnStart()                               { onStartCallback() }
func (cb *Callback) TriggerOnStop()                                { onStopCallback() }
func (cb *Callback) TriggerOnPlayerConnect(player *models.IPlayer) { onPlayerConnectCallback(player) }
func (cb *Callback) TriggerOnEnterVehicle(player *models.IPlayer, vehicle *models.IVehicle, seat uint8) {
	onEnterVehicleCallback(player, vehicle, seat)
}
