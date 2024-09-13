package enum

/*
   Create by zyx
   Date Time: 2024/9/11
   File: vehicle_data_type.go
*/

type VehicleDataType int32

const (
	PrimaryColor VehicleDataType = iota
	SecondColor
	NumberPlate
	VehicleCollision
	VehicleFrozen
	VehicleVisible
	EngineOn
	LockState
	LightState
	HeadLightColor
	NeonColor
	NeonActive
	VehicleDimension
	VehiclePosition
	VehicleRot
	DriftMode
	DoorState
	DisableTowing
	DirtLevel
	BodyHealth
	EngineHealth
	LightsMultiplier
	WheelColor
	RearWheels
	Mod
	ModKit
)
