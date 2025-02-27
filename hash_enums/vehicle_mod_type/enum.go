package vehicle_mod_type

/*
   Create by zyx
   Date Time: 2024/9/24
   File: enums.go
*/

//go:generate stringer -type=VehicleModType
type VehicleModType uint32

const (
	Spoilers VehicleModType = iota
	FrontBumper
	RearBumper
	SideSkirt
	Exhaust
	Frame
	Grille
	Hood
	Fender
	RightFender
	Roof
	Engine
	Brakes
	Transmission
	Horns
	Suspension
	Armor
	Turbo = 18
	Xenon = 22
	FrontWheels
	BackWheels
	PlateHolders
	TrimDesign = 27
	Ornaments
	DialDesign    = 30
	Seats         = 32
	SteeringWheel = 33
	ShiftLever
	Plaques
	Hydraulics = 38
	Boost      = 40
	Livery     = 48
	Plate      = 62
	Color1     = 66
	Color2
	WindowTint     = 69
	DashboardColor = 74
	TrimColor
)
