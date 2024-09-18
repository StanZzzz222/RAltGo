package blip_type

/*
   Create by zyx
   Date Time: 2024/9/12
   File: blip_type.go
*/

type BlipType int32

const (
	Point BlipType = iota
	Area
	Radius
)
