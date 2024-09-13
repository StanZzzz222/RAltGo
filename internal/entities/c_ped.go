package entities

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_ped.h"
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

type CPed struct {
	ID       uint32
	Model    uint32
	Position *Vector3
	Rotation *Vector3
}

func ConvertCPed(cpPtr uintptr) *CPed {
	cPed := (*C.CPed)(unsafe.Pointer(cpPtr))
	if cPed == nil {
		return nil
	}
	return &CPed{
		ID:       uint32(cPed.id),
		Model:    uint32(cPed.model),
		Position: (*Vector3)(unsafe.Pointer(cPed.position)),
		Rotation: (*Vector3)(unsafe.Pointer(cPed.rotation)),
	}
}
