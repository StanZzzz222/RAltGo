package init

import "C"
import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/alt_events"
	"github.com/StanZzzz222/RAltGo/internal/entity"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"github.com/StanZzzz222/RAltGo/logger"
	"github.com/StanZzzz222/RAltGo/models"
	"os"
	"time"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/9
   File: init.go
*/

var w = &lib.Warrper{}

func Mounted() {}

//export onModuleInit
func onModuleInit(cAltvVersion, core, cResourceName, cResourceHandlers, cModuleHandlers unsafe.Pointer) bool {
	path, _ := os.Getwd()
	path = fmt.Sprintf("%v/resources/rs-go-module/server/target/debug/server.dll", path)
	//path = fmt.Sprintf("%v/modules/rs-go-module.dll", path)
	logger.LogInfo(":: Go module Initialize mounting")
	defer logger.LogSuccess(":: Go module Initialize mounting done")
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		logger.LogErrorf(":: Please check if %v exists", path)
		time.Sleep(time.Second * 3)
		os.Exit(-1)
		return false
	}
	w.InitDLL(path)
	return w.ModuleMain(uintptr(cAltvVersion), uintptr(core), uintptr(cResourceName), uintptr(cResourceHandlers), uintptr(cModuleHandlers))
}

//export onStart
func onStart() {
	var cb = &alt_events.Callback{}
	cb = cb.New()
	cb.TriggerOnStart()
}

//export onStop
func onStop() {
	var cb = &alt_events.Callback{}
	cb = cb.New()
	cb.TriggerOnStop()
}

//export onPlayerConnect
func onPlayerConnect(cPtr *C.CPlayer) {
	var player = &models.IPlayer{}
	var cb = &alt_events.Callback{}
	var ptr = uintptr(unsafe.Pointer(cPtr))
	defer w.FreePlayer(ptr)
	cPlayer := (*entity.CPlayer)(unsafe.Pointer(cPtr))
	player = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
	cb = cb.New()
	cb.TriggerOnPlayerConnect(player)
}

//export onEnterVehicle
func onEnterVehicle(cPtr *C.CPlayer, cvPtr *C.CVehicle, seat uint8) {
	var player = &models.IPlayer{}
	var vehicle = &models.IVehicle{}
	var cb = &alt_events.Callback{}
	var ptr = uintptr(unsafe.Pointer(cvPtr))
	defer w.FreeVehicle(ptr)
	cPlayer := (*entity.CPlayer)(unsafe.Pointer(cPtr))
	cVehicle := (*entity.CVehicle)(unsafe.Pointer(cvPtr))
	player = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
	vehicle = vehicle.NewIVehicle(cVehicle.ID, cVehicle.Model, cVehicle.PrimaryColor, cVehicle.SecondColor, cVehicle.Position, cVehicle.Rotation)
	cb.TriggerOnEnterVehicle(player, vehicle, seat)
}
