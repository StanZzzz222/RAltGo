package utils

import "github.com/StanZzzz222/RAltGo/common/models"

/*
   Create by zyx
   Date Time: 2024/9/12
   File: rgba.go
*/

func NewRGBA(r, g, b, a uint8) *models.Rgba {
	return &models.Rgba{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
