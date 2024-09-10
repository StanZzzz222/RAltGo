package entity

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

func ConvertCPlayer(cPtr *C.CPlayer) *CPlayer {
	return &CPlayer{
		ID:         uint32(cPtr.id),
		Name:       w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPtr.name))),
		IP:         w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPtr.ip))),
		AuthToken:  w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPtr.auth_token))),
		HWIDHash:   uint64(cPtr.hwid_hash),
		HWIDExHash: uint64(cPtr.hwid_ex_hash),
		Position:   (*models.Vector3)(unsafe.Pointer(cPtr.position)),
		Rotation:   (*models.Vector3)(unsafe.Pointer(cPtr.rotation)),
	}
}

func ConvertCVehicle(cvPtr *C.CVehicle) *CVehicle {
	return &CVehicle{
		ID:           uint32(cvPtr.id),
		Model:        uint32(cvPtr.model),
		PrimaryColor: uint8(cvPtr.primary_color),
		SecondColor:  uint8(cvPtr.second_color),
		Position:     (*models.Vector3)(unsafe.Pointer(cvPtr.position)),
		Rotation:     (*models.Vector3)(unsafe.Pointer(cvPtr.position)),
	}
}
