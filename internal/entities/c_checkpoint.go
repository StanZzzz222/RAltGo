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
   File: c_checkpoint.go
*/

type CCheckpoint struct {
	ID             uint32
	CheckpointType uint8
	Position       *Vector3
}

func ConvertCCheckPoint(ccPtr uintptr) *CCheckpoint {
	cCheckpoint := (*C.CCheckpoint)(unsafe.Pointer(ccPtr))
	if cCheckpoint == nil {
		return nil
	}
	return &CCheckpoint{
		ID:             uint32(cCheckpoint.id),
		CheckpointType: uint8(cCheckpoint.check_point_type),
		Position:       (*Vector3)(unsafe.Pointer(cCheckpoint.position)),
	}
}
