package entities

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_voice_channel.h"
*/
import "C"
import (
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/10/10
   File: c_voice_channel.go
*/

type CVoiceChannel struct {
	ID          uint32
	Spatial     bool
	MaxDistance float32
}

func ConvertCVoiceChannel(cvPtr uintptr) *CVoiceChannel {
	cVoiceChannel := (*C.CVoiceChannel)(unsafe.Pointer(cvPtr))
	if cVoiceChannel == nil {
		return nil
	}
	return &CVoiceChannel{
		ID:          uint32(cVoiceChannel.id),
		Spatial:     cVoiceChannel.spatial != 0,
		MaxDistance: float32(cVoiceChannel.max_distance),
	}
}
