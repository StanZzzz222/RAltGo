package common

import (
	"github.com/StanZzzz222/RAltGo/common/models"
)

/*
   Create by zyx
   Date Time: 2024/9/12
   File: basic.go
*/

func NewRGB(r, g, b uint8) *models.Rgba {
	return &models.Rgba{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}

func NewRGBA(r, g, b, a uint8) *models.Rgba {
	return &models.Rgba{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
