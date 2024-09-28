package enums

/*
   Create by zyx
   Date Time: 2024/9/14
   File: object_type.go
*/

type EntityDataType int32

const (
	NetOwner EntityDataType = iota
	SyncId
	Streamed
	StreamingDistance
	Detach
	AttachToEntityBoneName
	AttachToEntityBoneIndex
)
