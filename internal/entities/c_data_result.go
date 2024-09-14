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
	var w = &lib.Warrper{}
	cDataResult := (*C.CDataResult)(unsafe.Pointer(cresPtr))
	if cDataResult == nil {
		return nil
	}
	return &CDataResult{
		Tag:        Tag(uint8(cDataResult.tag)),
		U8Val:      uint8(cDataResult.union_CDataResult_data.u8_val),
		U16Val:     uint16(cDataResult.union_CDataResult_data.u16_val),
		U32Val:     uint32(cDataResult.union_CDataResult_data.u32_val),
		U64Val:     uint64(cDataResult.union_CDataResult_data.u64_val),
		BoolVal:    cDataResult.union_CDataResult_data.bool_value != 0,
		StringVal:  w.PtrMarshalGoString(uintptr(unsafe.Pointer(cDataResult.data.cstring_val))),
		Vector3Val: (*Vector3)(unsafe.Pointer(cDataResult.data.vector3_val)),
	}
}
