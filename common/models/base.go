package models

import (
	"github.com/StanZzzz222/RAltGo/internal/entitys"
	"github.com/StanZzzz222/RAltGo/internal/lib"
)

/*
   Create by zyx
   Date Time: 2024/9/11
   File: base.go
*/

var w = &lib.Warrper{}

type BaseObject struct {
	position  *entitys.Vector3
	rotation  *entitys.Vector3
	dimension int32
	frozen    bool
	collision bool
	visible   bool
}

func NewBaseObject(position, rotation *entitys.Vector3, dimension int32, frozen, collision, visible bool) *BaseObject {
	return &BaseObject{
		position:  position,
		rotation:  rotation,
		dimension: dimension,
		frozen:    frozen,
		collision: collision,
		visible:   visible,
	}
}
