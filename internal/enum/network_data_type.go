package enum

/*
   Create by zyx
   Date Time: 2024/9/12
   File: blip_data_type.go
*/

type NetworkDataType int32

const (
	NetworkMeta NetworkDataType = iota
	NetworkSyncedMeta
	NetworkStreamSyncedMeta
)
