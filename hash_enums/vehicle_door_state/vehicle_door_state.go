package vehicle_door_state

/*
   Create by zyx
   Date Time: 2024/9/24
   File: vehicle_door_state.go
*/

// Extend: coreclr-module - VehicleDoorState

//go:generate stringer -type=VehicleDoorState
type VehicleDoorState byte

const (
	Closed        VehicleDoorState = 0
	OpenedLevel1  VehicleDoorState = 1
	OpenedLevel2  VehicleDoorState = 2
	OpenedLevel3  VehicleDoorState = 3
	OpenedLevel4  VehicleDoorState = 4
	OpenedLevel5  VehicleDoorState = 5
	OpenedLevel6  VehicleDoorState = 6
	OpenedLevel7  VehicleDoorState = 7
	DoesNotExists VehicleDoorState = 255
)
