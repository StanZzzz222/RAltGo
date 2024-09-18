package weather

/*
   Create by zyx
   Date Time: 2024/9/9
   File: enum.go
*/

// Extend: coreclr-module - WeatherType

//go:generate stringer -type=WeatherType
type WeatherType uint32

const (
	ExtraSunny WeatherType = iota
	Clear
	Clouds
	Smog
	Foggy
	Overcast
	Rain
	Thunder
	Clearing
	Neutral
	Snow
	Blizzard
	Snowlight
	Xmas
	Halloween
)
