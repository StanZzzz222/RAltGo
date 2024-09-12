package entitys

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_player.h"
	#include "c_vehicle.h"
	#include "c_blip.h"
*/
import "C"
import (
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"sync/atomic"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/9
   File: entitys.go
*/

var w = &lib.Warrper{}

type Vector3 struct {
	X, Y, Z float32
}

type CPlayer struct {
	ID         uint32
	Name       string
	IP         string
	AuthToken  string
	HWIDHash   uint64
	HWIDExHash uint64
	Position   *Vector3
	Rotation   *Vector3
}

type CVehicle struct {
	ID           uint32
	Model        uint32
	PrimaryColor uint8
	SecondColor  uint8
	Position     *Vector3
	Rotation     *Vector3
}

type CBlip struct {
	ID       uint32
	BlipType uint32
	SpriteId uint32
	Color    uint32
	Rot      float32
	Position *Vector3
}

func ConvertCPlayer(cPtr uintptr) *CPlayer {
	cPlayer := (*C.CPlayer)(unsafe.Pointer(atomic.LoadUintptr(&cPtr)))
	if cPlayer == nil {
		return nil
	}
	return &CPlayer{
		ID:         uint32(cPlayer.id),
		Name:       w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.name))),
		IP:         w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.ip))),
		AuthToken:  w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.auth_token))),
		HWIDHash:   uint64(cPlayer.hwid_hash),
		HWIDExHash: uint64(cPlayer.hwid_ex_hash),
		Position:   (*Vector3)(unsafe.Pointer(cPlayer.position)),
		Rotation:   (*Vector3)(unsafe.Pointer(cPlayer.rotation)),
	}
}

func ConvertCVehicle(cvPtr uintptr) *CVehicle {
	cVehicle := (*C.CVehicle)(unsafe.Pointer(atomic.LoadUintptr(&cvPtr)))
	if cVehicle == nil {
		return nil
	}
	return &CVehicle{
		ID:           uint32(cVehicle.id),
		Model:        uint32(cVehicle.model),
		PrimaryColor: uint8(cVehicle.primary_color),
		SecondColor:  uint8(cVehicle.second_color),
		Position:     (*Vector3)(unsafe.Pointer(cVehicle.position)),
		Rotation:     (*Vector3)(unsafe.Pointer(cVehicle.position)),
	}
}

func ConvertCBlip(cbPtr uintptr) *CBlip {
	cBlip := (*C.CBlip)(unsafe.Pointer(atomic.LoadUintptr(&cbPtr)))
	if cBlip == nil {
		return nil
	}
	return &CBlip{
		ID:       uint32(cBlip.id),
		BlipType: uint32(cBlip.blip_type),
		SpriteId: uint32(cBlip.sprite_id),
		Color:    uint32(cBlip.color),
		Rot:      float32(cBlip.rot),
		Position: (*Vector3)(unsafe.Pointer(cBlip.position)),
	}
}
