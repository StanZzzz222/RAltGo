package entities

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_data_result.h"
*/
import "C"
import (
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/14
   File: c_data_result.go
*/

type Tag uint8

const (
	U8Tag Tag = iota
	U16Tag
	U32Tag
	U64Tag
	BoolTag
	StringTag
	Vector3Tag
)

type CDataResult struct {
	Tag        Tag
	U8Val      uint8
	U16Val     uint16
	U32Val     uint32
	U64Val     uint64
	BoolVal    bool
	StringVal  string
	Vector3Val *Vector3
}

func ConverCDataResult(cresPtr uintptr) *CDataResult {
	var w = lib.GetWarpper()
	cDataResult := (*C.CDataResult)(unsafe.Pointer(cresPtr))
	if cDataResult == nil {
		return nil
	}
	res := &CDataResult{Tag: Tag(uint8(cDataResult.tag))}
	switch res.Tag {
	case U8Tag:
		res.U8Val = uint8(*(*C.uint8_t)(unsafe.Pointer(&cDataResult.data)))
		break
	case U16Tag:
		res.U16Val = uint16(*(*C.uint16_t)(unsafe.Pointer(&cDataResult.data)))
		break
	case U32Tag:
		res.U32Val = uint32(*(*C.uint32_t)(unsafe.Pointer(&cDataResult.data)))
		break
	case U64Tag:
		res.U64Val = uint64(*(*C.uint64_t)(unsafe.Pointer(&cDataResult.data)))
		break
	case BoolTag:
		res.BoolVal = int(*(*C.int)(unsafe.Pointer(&cDataResult.data))) != 0
		break
	case StringTag:
		res.StringVal = w.PtrMarshalGoString(uintptr(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&cDataResult.data)))))
		break
	case Vector3Tag:
		res.Vector3Val = (*Vector3)(unsafe.Pointer(*(**C.Vector3)(unsafe.Pointer(&cDataResult.data))))
		break
	default:
	}
	return res
}
