package colshape_type

/*
   Create by zyx
   Date Time: 2024/9/19
   File: colshape_type.go
*/

type ColshapeType int32

const (
	Circle ColshapeType = iota
	Sphere
	Rectangle
	Cuboid
	Cylinder
	Polygon
)
