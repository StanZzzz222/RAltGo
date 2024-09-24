package vehicle_door_type

/*
   Create by zyx
   Date Time: 2024/9/24
   File: vehicle_door_type.go
*/

// Extend: coreclr-module - VehicleDoor

//go:generate stringer -type=VehicleDoorType
type VehicleDoorType byte

const (
	DriverFront VehicleDoorType = iota
	PassengerFront
	DriverRear
	PassengerRear
	Hood
	Trunk
)
