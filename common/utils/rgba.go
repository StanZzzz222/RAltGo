package utils

import (
	"github.com/StanZzzz222/RAltGo/internal/entities"
)

/*
   Create by zyx
   Date Time: 2024/9/12
   File: basic.go
*/

func NewRGBA(r, g, b, a uint8) *entities.Rgba {
	return &entities.Rgba{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
