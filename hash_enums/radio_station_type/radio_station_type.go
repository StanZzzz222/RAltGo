package radio_station_type

/*
   Create by zyx
   Date Time: 2024/9/24
   File: radio_station_type.go
*/

// Extend: coreclr-module - RadioStation

//go:generate stringer -type=RadioStation
type RadioStation uint32

const (
	LosSantosRockRadio RadioStation = iota
	NonStopPopFm
	RadioLosSantos
	ChannelX
	WestCoastTalkRadio
	RebelRadio
	SoulwaxFm
	EastLosFm
	WestCoastClassics
	BlaineCountyRadio
	TheBlueArk
	WorldWideFm
	FlyloFm
	TheLowdown
	RadioMirrorPark
	Space
	VinewoodBoulevardRadio
	SelfRadio
	TheLab
	RadioOff RadioStation = 255
)
