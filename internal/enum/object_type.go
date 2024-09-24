package enum

/*
   Create by zyx
   Date Time: 2024/9/14
   File: object_type.go
*/

type ObjectType int32

const (
	Player ObjectType = iota
	Vehicle
	Ped
	Blip
	Colshape
	Object
	CheckPoint
	Marker
)
