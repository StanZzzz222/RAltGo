package entities

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_virtual_entity.h"
*/
import "C"
import (
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/29
   File: c_virtual_entity.go
*/

type CVirtualEntity struct {
	ID                uint32
	StreamingDistance uint32
	Position          *Vector3
}

func ConvertCVirtualEntity(ccPtr uintptr) *CVirtualEntity {
	cVirtualEntity := (*C.CVirtualEntity)(unsafe.Pointer(ccPtr))
	if cVirtualEntity == nil {
		return nil
	}
	return &CVirtualEntity{
		ID:                uint32(cVirtualEntity.id),
		StreamingDistance: uint32(cVirtualEntity.streaming_distance),
		Position:          (*Vector3)(unsafe.Pointer(cVirtualEntity.position)),
	}
}
