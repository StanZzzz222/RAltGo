package entities

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_virtual_entity_group.h"
*/
import "C"
import (
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/29
   File: c_virtual_entity_group.go
*/

type CVirtualEntityGroup struct {
	ID                  uint32
	MaxEntitiesInStream uint32
}

func ConvertCVirtualEntityGroup(ccPtr uintptr) *CVirtualEntityGroup {
	cVirtualEntityGroup := (*C.CVirtualEntityGroup)(unsafe.Pointer(ccPtr))
	if cVirtualEntityGroup == nil {
		return nil
	}
	return &CVirtualEntityGroup{
		ID:                  uint32(cVirtualEntityGroup.id),
		MaxEntitiesInStream: uint32(cVirtualEntityGroup.max_entities_in_stream),
	}
}
