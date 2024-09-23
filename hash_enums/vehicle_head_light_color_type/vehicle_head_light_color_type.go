package vehicle_head_light_color_type

/*
   Create by zyx
   Date Time: 2024/9/23
   File: vehicle_head_light_color_type.go
*/

type VehicleHeadLightColorType int

const (
	Default VehicleHeadLightColorType = -1
	White   VehicleHeadLightColorType = iota
	Blue
	ElectricBlue
	MintGreen
	LimeGreen
	Yellow
	GoldenShower
	Orange
	Red
	PonyPink
	HotPink
	Purple
	Blacklight
)
