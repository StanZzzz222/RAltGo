package enums

/*
   Create by zyx
   Date Time: 2024/9/19
   File: colshape_data_type.go
*/

type ColshapeDataType int32

const (
	ColshapeDimension ColshapeDataType = iota
	ColshapePlayersOnly
	ColshapePosition
	ColshapeDestory
	ColshapeIsPointIn
	ColshapeIsEntityIn
	ColshapeIsEntityIdIn
)
