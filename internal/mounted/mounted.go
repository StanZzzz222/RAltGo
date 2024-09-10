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
	"github.com/StanZzzz222/RAltGo/internal/entitys"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"github.com/StanZzzz222/RAltGo/logger"
	"github.com/StanZzzz222/RAltGo/models"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/9
   File: mounted.go
*/

var w = &lib.Warrper{}

func Mounted() {}

//export onModuleInit
func onModuleInit(cAltvVersion, core, cResourceName, cResourceHandlers, cModuleHandlers unsafe.Pointer) bool {
	logger.LogInfo(":: Go module Initialize mounting")
	defer logger.LogSuccess(":: Go module Initialize mounting done")
	return w.ModuleMain(uintptr(cAltvVersion), uintptr(core), uintptr(cResourceName), uintptr(cResourceHandlers), uintptr(cModuleHandlers))
}

//export onTick
func onTick() {
	w.GetTasks().Range(func(key, value any) bool {
		handler, ok := value.(func())
		if ok {
			handler()
			w.TaskDelete(key.(string))
		}
		return true
	})
}

//export onStart
func onStart() {
	var cb = &alt_events.Callback{}
	cb.TriggerOnStart()
}

//export onStop
func onStop() {
	var cb = &alt_events.Callback{}
	cb.TriggerOnStop()
}

//export onPlayerConnect
func onPlayerConnect(cplayer *C.CPlayer) {
	var player = &models.IPlayer{}
	var cPtr = uintptr(unsafe.Pointer(cplayer))
	var cPlayer = entitys.ConvertCPlayer(cPtr)
	var cb = &alt_events.Callback{}
	defer w.FreePlayer(cPtr)
	player = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
	cb.TriggerOnPlayerConnect(player)
}

//export onEnterVehicle
func onEnterVehicle(cplayer *C.CPlayer, cvehicle *C.CVehicle, seat uint8) {
	var player = &models.IPlayer{}
	var vehicle = &models.IVehicle{}
	var cb = &alt_events.Callback{}
	var cPtr = uintptr(unsafe.Pointer(cplayer))
	var cvPtr = uintptr(unsafe.Pointer(cvehicle))
	var cPlayer = entitys.ConvertCPlayer(cPtr)
	var cVehicle = entitys.ConvertCVehicle(cvPtr)
	defer w.FreePlayer(cPtr)
	defer w.FreeVehicle(cvPtr)
	player = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
	vehicle = vehicle.NewIVehicle(cVehicle.ID, cVehicle.Model, cVehicle.PrimaryColor, cVehicle.SecondColor, cVehicle.Position, cVehicle.Rotation)
	cb.TriggerOnEnterVehicle(player, vehicle, seat)
}
