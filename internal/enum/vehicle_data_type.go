package enum

/*
   Create by zyx
   Date Time: 2024/9/11
   File: vehicle_data_type.go
*/

type VehicleDataType int32

const (
	VehiclePrimaryColor VehicleDataType = iota
	VehicleSecondColor
	VehicleNumberPlate
	VehicleCollision
	VehicleFrozen
	VehicleVisible
	VehicleEngineOn
	VehicleLockState
	VehicleLightState
	VehicleHeadLightColor
	VehicleNeonColor
	VehicleNeonActive
	VehicleDimension
	VehiclePosition
	VehicleRot
	VehicleDriftMode
	VehicleDoorState
	VehicleDisableTowing
	VehicleDirtLevel
	VehicleBodyHealth
	VehicleEngineHealth
	VehicleLightsMultiplier
	VehicleWheelColor
	VehicleRearWheels
	VehicleMod
	VehicleModKit
	VehicleRepair
	VehicleDestroy
	VehicleInteriorColor
	VehicleBoatAnchorActive
	VehicleCustomTires
	VehicleLightDamaged
	VehicleRadioStation
	VehicleDashboardColor
	VehicleWindowTint
	VehicleDriver
)
