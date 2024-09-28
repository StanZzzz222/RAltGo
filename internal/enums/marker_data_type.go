package enums

/*
   Create by zyx
   Date Time: 2024/9/24
   File: checkpoint_data_type.go
*/

type MarkerDataType int32

const (
	MarkerDimension MarkerDataType = iota
	MarkerPosition
	MarkerRotation
	MarkerVisible
	MarkerColor
	MarkerScale
	MarkerBobUpDown
	MarkerDir
	MarkerFaceCamera
	MarkerRotating
	MarkerType
	MarkerTarget
	MarkerDestory
)
