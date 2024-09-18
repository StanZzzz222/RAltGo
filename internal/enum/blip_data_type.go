package enum

/*
   Create by zyx
   Date Time: 2024/9/12
   File: blip_data_type.go
*/

type BlipDataType int32

const (
	Sprite BlipDataType = iota
	Color
	RgbaColor
	BlipVisible
	Display
	Alpha
	Friendly
	HighDetail
	MissionCreator
	ShortRange
	Bright
	CrewIndicatorVisible
	Category
	FlashInterval
	FlashTimer
	Flashes
	FlashesAlternate
	Global
	MinimalOnEdge
	Route
	Pulse
	HiddenOnLegend
	OutlineIndicatorVisible
	BlipRot
	Shrinked
	ShowCone
	TickVisible
	UseHeightIndicatorOnEdge
	BlipPosition
	Name
	RouteColor
	HeadingIndicatorVisible
	ShortHeightThreshold
	Number
	Scale
	GxtName
	BlipType
	BlipDestroy
)
