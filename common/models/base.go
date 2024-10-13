package models

import (
	"fmt"
	"math"
	"strings"
)

/*
   Create by zyx
   Date Time: 2024/9/11
   File: base.go
*/

type BaseObject struct {
	position  *Vector3
	rotation  *Vector3
	dimension int32
	frozen    bool
	collision bool
	visible   bool
}

type Vector3 struct {
	X, Y, Z float32
}

type Rgba struct {
	R, G, B, A uint8
}

func NewBaseObject(position, rotation *Vector3, dimension int32, frozen, collision, visible bool) *BaseObject {
	return &BaseObject{
		position:  position,
		rotation:  rotation,
		dimension: dimension,
		frozen:    frozen,
		collision: collision,
		visible:   visible,
	}
}

func (v *Vector3) Distance(target *Vector3) float32 {
	return float32(math.Sqrt(float64(
		(target.X-v.X)*(target.X-v.X) +
			(target.Y-v.Y)*(target.Y-v.Y) +
			(target.Z-v.Z)*(target.Z-v.Z))))
}

func (v *Vector3) ToString() string {
	return fmt.Sprintf("%v,%v,%v", v.X, v.Y, v.Z)
}

func (r *Rgba) ToString() string {
	return fmt.Sprintf("%v,%v,%v,%v", r.R, r.G, r.B, r.A)
}

func hash(model string) uint32 {
	k := strings.ToLower(model)
	var h uint32
	var i int
	for i = 0; i < len(k); i++ {
		h += uint32(k[i])
		h += h << 10
		h ^= h >> 6
	}
	h += h << 3
	h ^= h >> 11
	h += h << 15
	return h
}
