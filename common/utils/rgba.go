package utils

/*
   Create by zyx
   Date Time: 2024/9/12
   File: rgba.go
*/

type Rgba struct {
	R, G, B, A uint8
}

func NewRGBA(r, g, b, a uint8) *Rgba {
	return &Rgba{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
