package entities

import "math"

/*
   Create by zyx
   Date Time: 2024/9/12
   File: basic.go
*/

type Vector3 struct {
	X, Y, Z float32
}

type Rgba struct {
	R, G, B, A uint8
}

func (v *Vector3) Distance(target *Vector3) float32 {
	return float32(math.Sqrt(float64(
		(target.X-v.X)*(target.X-v.X) +
			(target.Y-v.Y)*(target.Y-v.Y) +
			(target.Z-v.Z)*(target.Z-v.Z))))
}
