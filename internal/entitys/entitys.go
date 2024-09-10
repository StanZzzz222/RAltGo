package entitys

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_player.h"
	#include "c_vehicle.h"
*/
import "C"
import (
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"github.com/StanZzzz222/RAltGo/models"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/9
   File: entitys.go
*/

var w = &lib.Warrper{}

type CPlayer struct {
	ID         uint32
	Name       string
	IP         string
	AuthToken  string
	HWIDHash   uint64
	HWIDExHash uint64
	Position   *models.Vector3
	Rotation   *models.Vector3
}

type CVehicle struct {
	ID           uint32
	Model        uint32
	PrimaryColor uint8
	SecondColor  uint8
	Position     *models.Vector3
	Rotation     *models.Vector3
}

func ConvertCPlayer(cPtr uintptr) *CPlayer {
	cPlayer := (*C.CPlayer)(unsafe.Pointer(cPtr))
	return &CPlayer{
		ID:         uint32(cPlayer.id),
		Name:       w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.name))),
		IP:         w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.ip))),
		AuthToken:  w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.auth_token))),
		HWIDHash:   uint64(cPlayer.hwid_hash),
		HWIDExHash: uint64(cPlayer.hwid_ex_hash),
		Position:   (*models.Vector3)(unsafe.Pointer(cPlayer.position)),
		Rotation:   (*models.Vector3)(unsafe.Pointer(cPlayer.rotation)),
	}
}

func ConvertCVehicle(cvPtr uintptr) *CVehicle {
	cVehicle := (*C.CVehicle)(unsafe.Pointer(cvPtr))
	return &CVehicle{
		ID:           uint32(cVehicle.id),
		Model:        uint32(cVehicle.model),
		PrimaryColor: uint8(cVehicle.primary_color),
		SecondColor:  uint8(cVehicle.second_color),
		Position:     (*models.Vector3)(unsafe.Pointer(cVehicle.position)),
		Rotation:     (*models.Vector3)(unsafe.Pointer(cVehicle.position)),
	}
}
