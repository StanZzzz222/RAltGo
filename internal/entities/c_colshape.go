package entities

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_colshape.h"
*/
import "C"
import (
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/19
   File: c_colshape.go
*/

type CColshape struct {
	ID           uint32
	ColshapeType uint32
	Position     *Vector3
}

func ConvertCColshape(ccPtr uintptr) *CColshape {
	cColshape := (*C.CColshape)(unsafe.Pointer(ccPtr))
	if cColshape == nil {
		return nil
	}
	return &CColshape{
		ID:           uint32(cColshape.id),
		ColshapeType: uint32(cColshape.colshape_type),
		Position:     (*Vector3)(unsafe.Pointer(cColshape.position)),
	}
}
