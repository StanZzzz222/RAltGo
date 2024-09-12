package entities

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_blip.h"
*/
import "C"
import (
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/12
   File: c_blip.go
*/

type CBlip struct {
	ID       uint32
	BlipType uint32
	SpriteId uint32
	Color    uint32
	Rot      float32
	Position *Vector3
}

func ConvertCBlip(cbPtr uintptr) *CBlip {
	cBlip := (*C.CBlip)(unsafe.Pointer(cbPtr))
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