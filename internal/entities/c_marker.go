package entities

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_check_point.h"
*/
import "C"
import (
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/25
   File: c_object.go
*/

type CMarker struct {
	ID         uint32
	MarkerType uint8
	Position   *Vector3
}

func ConvertCMarker(cPtr uintptr) *CMarker {
	cMarker := (*C.CObject)(unsafe.Pointer(cPtr))
	if cMarker == nil {
		return nil
	}
	return &CMarker{
		ID:         uint32(cMarker.id),
		MarkerType: uint8(cMarker.marker_type),
		Position:   (*Vector3)(cMarker.position),
	}
}
