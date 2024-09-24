package numplate_style_type

/*
   Create by zyx
   Date Time: 2024/9/24
   File: numplate_style_type.go
*/

// Extend: coreclr-module - NumberPlateStyle

//go:generate stringer -type=NumberPlateStyle
type NumberPlateStyle int

const (
	BlueWhite NumberPlateStyle = iota
	YellowBlack
	YellowBlue
	BlueWhite2
	BlueWhite3
	Yankton
)
