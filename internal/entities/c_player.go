package entities

/*
	#cgo CFLAGS: -I../headers
	#include "c_vector3.h"
	#include "c_player.h"
*/
import "C"
import (
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/12
   File: c_player.go
*/

type CPlayer struct {
	ID         uint32
	Name       string
	IP         string
	AuthToken  string
	SocialName string
	SocialID   uint64
	HWIDHash   uint64
	HWIDExHash uint64
	Position   *Vector3
	Rotation   *Vector3
}

func ConvertCPlayer(cPtr uintptr) *CPlayer {
	var w = lib.GetWrapper()
	cPlayer := (*C.CPlayer)(unsafe.Pointer(cPtr))
	if cPlayer == nil {
		return nil
	}
	return &CPlayer{
		ID:         uint32(cPlayer.id),
		Name:       w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.name))),
		IP:         w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.ip))),
		AuthToken:  w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.auth_token))),
		SocialName: w.PtrMarshalGoString(uintptr(unsafe.Pointer(cPlayer.social_name))),
		SocialID:   uint64(cPlayer.social_id),
		HWIDHash:   uint64(cPlayer.hwid_hash),
		HWIDExHash: uint64(cPlayer.hwid_ex_hash),
		Position:   (*Vector3)(unsafe.Pointer(cPlayer.position)),
		Rotation:   (*Vector3)(unsafe.Pointer(cPlayer.rotation)),
	}
}
