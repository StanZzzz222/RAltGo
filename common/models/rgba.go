package models

/*
   Create by zyx
   Date Time: 2024/9/11
   File: rgba.go
*/

type Rgba struct {
	r, g, b, a uint8
}

func NewRGBA(r, g, b, a uint8) *Rgba {
	return &Rgba{r, g, b, a}
}
