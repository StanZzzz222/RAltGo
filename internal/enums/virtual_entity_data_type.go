package enums

/*
   Create by zyx
   Date Time: 2024/9/19
   File: colshape_data_type.go
*/

type VirtualEntityDataType int32

const (
	VirtualEntityDimension VirtualEntityDataType = iota
	VirtualEntityPosition
	VirtualEntityVisible
	VirtualEntityDestory
)
