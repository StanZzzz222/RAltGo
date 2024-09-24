package entities

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_object.h"
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

type CObject struct {
	ID uint32
}

func ConvertCObject(cPtr uintptr) *CObject {
	cObject := (*C.CObject)(unsafe.Pointer(cPtr))
	if cObject == nil {
		return nil
	}
	return &CObject{
		ID: uint32(cObject.id),
	}
}
