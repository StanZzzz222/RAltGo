package mounted

import "C"
import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/alt_events"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"github.com/StanZzzz222/RAltGo/logger"
	"github.com/StanZzzz222/RAltGo/models"
	"os"
	"time"
	"unsafe"
)

/*
	#include <stdint.h>

	typedef struct {
		float x;
		float y;
		float z;
	} Vector3;

	typedef struct {
		uint32_t id;
		const char *name;
		const char *ip;
		const char *auth_token;
		uint64_t hwid_hash;
		uint64_t hwid_ex_hash;
		const Vector3 *position;
		const Vector3 *rotation;
	} CPlayer;

	typedef struct {
		uint32_t id;
		uint32_t model;
		uint8_t primary_color;
		uint8_t second_color;
		const Vector3 *position;
		const Vector3 *rotation;
	} CVehicle;
*/
import "C"

/*
   Create by zyx
   Date Time: 2024/9/9
   File: mounted.go
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
func onPlayerConnect(cPlayer *C.CPlayer) {
	var player = &models.IPlayer{}
	var cb = &alt_events.Callback{}
	var ptr = uintptr(unsafe.Pointer(cPlayer))
	cb = cb.New()
	defer w.FreePlayer(ptr)
	id := uint32(cPlayer.id)
	name := w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.name)))
	ip := w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.ip)))
	authToken := w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.auth_token)))
	hwIdHash := uint64(cPlayer.hwid_hash)
	hwIdExHash := uint64(cPlayer.hwid_ex_hash)
	pos := (*models.Vector3)(unsafe.Pointer(cPlayer.position))
	rot := (*models.Vector3)(unsafe.Pointer(cPlayer.rotation))
	player = player.NewIPlayer(id, name, ip, authToken, hwIdHash, hwIdExHash, pos, rot)
	cb.TriggerOnPlayerConnect(player)
}

//export onEnterVehicle
func onEnterVehicle(cPlayer *C.CPlayer, cVehicle *C.CVehicle, seat uint8) {
	var player = &models.IPlayer{}
	var vehicle = &models.IVehicle{}
	var cb = &alt_events.Callback{}
	var playerPtr = uintptr(unsafe.Pointer(cPlayer))
	defer w.FreePlayer(playerPtr)
	var vehiclePtr = uintptr(unsafe.Pointer(cVehicle))
	defer w.FreeVehicle(vehiclePtr)
	id := uint32(cPlayer.id)
	name := w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.name)))
	ip := w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.ip)))
	authToken := w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.auth_token)))
	hwIdHash := uint64(cPlayer.hwid_hash)
	hwIdExHash := uint64(cPlayer.hwid_ex_hash)
	pos := (*models.Vector3)(unsafe.Pointer(cPlayer.position))
	rot := (*models.Vector3)(unsafe.Pointer(cPlayer.rotation))
	player = player.NewIPlayer(id, name, ip, authToken, hwIdHash, hwIdExHash, pos, rot)
	vehId := uint32(cVehicle.id)
	vehModel := uint32(cVehicle.model)
	vehPrimaryColor := uint8(cVehicle.model)
	vehSecondColor := uint8(cVehicle.model)
	pos = (*models.Vector3)(unsafe.Pointer(cVehicle.position))
	rot = (*models.Vector3)(unsafe.Pointer(cVehicle.rotation))
	vehicle = vehicle.NewIVehicle(vehId, vehModel, vehPrimaryColor, vehSecondColor, pos, rot)
	cb.TriggerOnEnterVehicle(player, vehicle, seat)
}
