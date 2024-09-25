package enum

/*
   Create by zyx
   Date Time: 2024/9/12
   File: blip_data_type.go
*/

type PedDataType int32

const (
	PedPosition PedDataType = iota
	PedRotation
	PedHealth
	PedArmour
	PedDimension
	PedCollision
	PedFrozen
	PedVisible
	PedMaxHealth
	PedCurrentWeapon
	PedDestroy
)
