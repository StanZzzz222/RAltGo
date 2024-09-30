package enums

/*
   Create by zyx
   Date Time: 2024/9/12
   File: blip_data_type.go
*/

type ServerDataType int

const (
	ServerStreamingDistance ServerDataType = iota
	ServerMigrationDistance
	ServerPassword
	ServerColShapeTickRate
	ServerMigrationTickRate
	ServerMaxStreamingObjects
	ServerMaxStreamingPeds
	ServerMaxStreamingVehicles
	ServerMigrationThreadCount
	ServerStreamerThreadCount
	ServerStreamingTickRate
	ServerSyncReceiveThreadCount
	ServerSyncSendThreadCount
	ServerVoiceExternal
	ServerVoiceExternalPublic
	ServerVoiceConnectionState
	ServerAmmoHashForWeaponHash
)
