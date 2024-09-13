package main

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_player.h"
	#include "c_vehicle.h"
*/
import "C"
import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/alt/alt_events"
	"github.com/StanZzzz222/RAltGo/common/alt/blip"
	"github.com/StanZzzz222/RAltGo/common/alt/scheduler"
	"github.com/StanZzzz222/RAltGo/common/alt/vehicle"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/common/utils"
	"github.com/StanZzzz222/RAltGo/internal/entities"
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

func main() {}

//export onModuleInit
func onModuleInit(cAltvVersion, core, cResourceName, cResourceHandlers, cModuleHandlers unsafe.Pointer) bool {
	logger.LogInfo(":: Go module Initialize mounting")
	defer logger.LogSuccess(":: Go module Initialize mounting done")
	return w.ModuleMain(uintptr(cAltvVersion), uintptr(core), uintptr(cResourceName), uintptr(cResourceHandlers), uintptr(cModuleHandlers))
}

func benchmark() {
	s := scheduler.NewScheduler()
	s.AddTask(func() {
		start := time.Now()
		for i := 0; i < 5000; i++ {
			_ = blip.CreateBlipPoint(10, 1, fmt.Sprintf("test%v", i), utils.NewVector3(0, 0, 0))
			_ = blip.CreateBlipArea(10, 1, fmt.Sprintf("test%v", i), utils.NewVector3(0, 0, 0), 10, 10)
			_ = blip.CreateBlipRadius(10, 1, fmt.Sprintf("test%v", i), utils.NewVector3(0, 0, 0), 10)
			_ = blip.CreateBlipRadius(10, 1, fmt.Sprintf("test%v", i), utils.NewVector3(0, 0, 0), 10)
			_ = blip.CreateBlipRadius(10, 1, fmt.Sprintf("test%v", i), utils.NewVector3(0, 0, 0), 10)
			_ = blip.CreateBlipRadius(10, 1, fmt.Sprintf("test%v", i), utils.NewVector3(0, 0, 0), 10)
			_ = blip.CreateBlipRadius(10, 1, fmt.Sprintf("test%v", i), utils.NewVector3(0, 0, 0), 10)
			_ = blip.CreateBlipRadius(10, 1, fmt.Sprintf("test%v", i), utils.NewVector3(0, 0, 0), 10)
		}
		fmt.Printf("Create blip points: %v ms\n", time.Since(start).Milliseconds())
	})
	s.AddTask(func() {
		start := time.Now()
		vec := utils.NewVector3(0, 0, 0)
		for i := 0; i < 50000; i++ {
			_ = vehicle.CreateVehicle("t20", "test", vec, vec, 1, 1)
		}
		fmt.Printf("Create vehicles: %v ms\n", time.Since(start).Milliseconds())
	})
	s.Run()
	s = scheduler.NewScheduler()
	s.AddTask(func() {
		fmt.Println("test")
	})
	s.Run()
	fmt.Println("继续")
}

//export onStart
func onStart() {
	cb.TriggerOnStart()
	benchmark()
}

//export onServerStarted
func onServerStarted() {
	cb.TriggerOnServerStarted()
}

//export onStop
func onStop() {
	cb.TriggerOnStop()
}

//export onPlayerConnect
func onPlayerConnect(cPtr uintptr) {
	var player = &models.IPlayer{}
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		defer w.FreePlayer(cPtr)
		player = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
		cb.TriggerOnPlayerConnect(player)
	}
}

//export onEnterVehicle
func onEnterVehicle(cPtr, cvPtr uintptr, seat uint8) {
	var player = &models.IPlayer{}
	var veh = &models.IVehicle{}
	var cPlayer = entities.ConvertCPlayer(cPtr)
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	if cPlayer != nil && cVehicle != nil {
		player = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
		veh = veh.NewIVehicle(cVehicle.ID, cVehicle.Model, cVehicle.PrimaryColor, cVehicle.SecondColor, cVehicle.Position, cVehicle.Rotation)
		cb.TriggerOnEnterVehicle(player, veh, seat)
		w.FreePlayer(cPtr)
		w.FreeVehicle(cvPtr)
	}
}

//export onLeaveVehicle
func onLeaveVehicle(cPtr, cvPtr uintptr, seat uint8) {
	var player = &models.IPlayer{}
	var veh = &models.IVehicle{}
	var cPlayer = entities.ConvertCPlayer(cPtr)
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	if cPlayer != nil && cVehicle != nil {
		player = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
		veh = veh.NewIVehicle(cVehicle.ID, cVehicle.Model, cVehicle.PrimaryColor, cVehicle.SecondColor, cVehicle.Position, cVehicle.Rotation)
		cb.TriggerOnLeaveVehicle(player, veh, seat)
		w.FreePlayer(cPtr)
		w.FreeVehicle(cvPtr)
	}
}
