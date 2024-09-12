package entities

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_vehicle.h"
*/
import "C"
import (
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/12
   File: c_vehicle.go
*/

type CVehicle struct {
	ID           uint32
	Model        uint32
	PrimaryColor uint8
	SecondColor  uint8
	Position     *Vector3
	Rotation     *Vector3
}

func ConvertCVehicle(cvPtr uintptr) *CVehicle {
	cVehicle := (*C.CVehicle)(unsafe.Pointer(cvPtr))
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
