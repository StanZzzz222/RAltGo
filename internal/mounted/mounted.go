package mounted

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_player.h"
	#include "c_vehicle.h"
*/
import "C"
import (
	"github.com/StanZzzz222/RAltGo/alt_events"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/internal/entitys"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"github.com/StanZzzz222/RAltGo/logger"
	"time"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/9
   File: mounted.go
*/

var w = &lib.Warrper{}
var cb = &alt_events.Callback{}

func Mounted() {}

//export onModuleInit
func onModuleInit(cAltvVersion, core, cResourceName, cResourceHandlers, cModuleHandlers unsafe.Pointer) bool {
	logger.LogInfo(":: Go module Initialize mounting")
	defer logger.LogSuccess(":: Go module Initialize mounting done")
	return w.ModuleMain(uintptr(cAltvVersion), uintptr(core), uintptr(cResourceName), uintptr(cResourceHandlers), uintptr(cModuleHandlers))
}

//export onStart
func onStart() {
	defer func() {
		time.AfterFunc(time.Millisecond*100, func() {
			cb.TriggerOnStart()
		})
	}()
}

//export onServerStarted
func onServerStarted() {
	defer func() {
		time.AfterFunc(time.Millisecond*100, func() {
			cb.TriggerOnServerStarted()
		})
	}()
}

//export onStop
func onStop() {
	defer func() {
		time.AfterFunc(time.Millisecond*100, func() {
			cb.TriggerOnStop()
		})
	}()
}

//export onPlayerConnect
func onPlayerConnect(cPtr uintptr) {
	var player = &models.IPlayer{}
	var cPlayer = entitys.ConvertCPlayer(cPtr)
	defer w.FreePlayer(cPtr)
	player = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
	cb.TriggerOnPlayerConnect(player)
}

//export onEnterVehicle
func onEnterVehicle(cPtr, cvPtr uintptr, seat uint8) {
	var player = &models.IPlayer{}
	var veh = &models.IVehicle{}
	var cPlayer = entitys.ConvertCPlayer(cPtr)
	var cVehicle = entitys.ConvertCVehicle(cvPtr)
	defer w.FreePlayer(cPtr)
	defer w.FreeVehicle(cvPtr)
	player = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
	veh = veh.NewIVehicle(cVehicle.ID, cVehicle.Model, cVehicle.PrimaryColor, cVehicle.SecondColor, cVehicle.Position, cVehicle.Rotation)
	cb.TriggerOnEnterVehicle(player, veh, seat)
}

//export onLeaveVehicle
func onLeaveVehicle(cPtr, cvPtr uintptr, seat uint8) {
	var player = &models.IPlayer{}
	var veh = &models.IVehicle{}
	var cPlayer = entitys.ConvertCPlayer(cPtr)
	var cVehicle = entitys.ConvertCVehicle(cvPtr)
	defer w.FreePlayer(cPtr)
	defer w.FreeVehicle(cvPtr)
	player = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
	veh = veh.NewIVehicle(cVehicle.ID, cVehicle.Model, cVehicle.PrimaryColor, cVehicle.SecondColor, cVehicle.Position, cVehicle.Rotation)
	cb.TriggerOnLeaveVehicle(player, veh, seat)
}
