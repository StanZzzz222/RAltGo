package utils

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/internal/entitys"
	"strconv"
	"strings"
)

/*
   Create by zyx
   Date Time: 2024/9/5
   File: vector3.go
*/

func NewVector3(x, y, z float32) *entitys.Vector3 {
	return &entitys.Vector3{X: x, Y: y, Z: z}
}

func NewVector3Round(x, y, z float32, round float32) *entitys.Vector3 {
	return &entitys.Vector3{X: x, Y: y, Z: z}
}

func NewVector3FromStr(position string) (*entitys.Vector3, error) {
	position = strings.TrimSpace(position)
	arr := strings.Split(position, ",")
	if len(arr) != 3 {
		return nil, fmt.Errorf("invalid position format")
	}
	var coords [3]float32
	for i, str := range arr {
		var err error
		coord, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
		if err != nil {
			return nil, err
		}
		coords[i] = float32(coord)
	}
	return NewVector3(coords[0], coords[1], coords[2]), nil
}
