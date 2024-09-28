package enums

/*
   Create by zyx
   Date Time: 2024/9/24
   File: checkpoint_data_type.go
*/

type ObjectDataType int32

const (
	ObjectDimension ObjectDataType = iota
	ObjectDestory
	ObjectPosition
	ObjectRotation
	ObjectAlpha
	ObjectVisible
	ObjectLodDistance
	ObjectTextureVariation
	ObjectFrozen
	ObjectCollision
)
