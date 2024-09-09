package models

import (
	"fmt"
	"strconv"
	"strings"
)

/*
   Create by zyx
   Date Time: 2024/9/5
   File: vec3.go
*/

type Vector3 struct{ X, Y, Z float32 }

func NewVector3(x, y, z float32) *Vector3 {
	return &Vector3{x, y, z}
}

func NewVector3Round(x, y, z float32, round float32) *Vector3 {
	return &Vector3{x, y, z}
}

func NewVector3FromStr(position string) (*Vector3, error) {
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
