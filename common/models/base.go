package models

import (
	"github.com/StanZzzz222/RAltGo/internal/entities"
)

/*
   Create by zyx
   Date Time: 2024/9/11
   File: base.go
*/

type BaseObject struct {
	position  *entities.Vector3
	rotation  *entities.Vector3
	dimension int32
	frozen    bool
	collision bool
	visible   bool
}

func NewBaseObject(position, rotation *entities.Vector3, dimension int32, frozen, collision, visible bool) *BaseObject {
	return &BaseObject{
		position:  position,
		rotation:  rotation,
		dimension: dimension,
		frozen:    frozen,
		collision: collision,
		visible:   visible,
	}
}
