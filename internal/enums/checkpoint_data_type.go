package enums

/*
   Create by zyx
   Date Time: 2024/9/24
   File: checkpoint_data_type.go
*/

type CheckpointDataType int32

const (
	CheckpointDimension CheckpointDataType = iota
	CheckpointVisible
	CheckpointPosition
	CheckpointDestory
	CheckpointType
	CheckpointHeight
	CheckpointPlayersOnly
	CheckpointNextPosition
	CheckpointRadius
	CheckpointColor
	CheckpointIconColor
)
