package enums

/*
   Create by zyx
   Date Time: 2024/9/24
   File: checkpoint_data_type.go
*/

type VoiceChannelDataType int32

const (
	VoiceChannelAddPlayer VoiceChannelDataType = iota
	VoiceChannelRemovePlayer
	VoiceChannelIsPlayerMuted
	VoiceChannelMutePlayer
	VoiceChannelUnmutePlayer
	VoiceChannelFilter
	VoiceChannelPriority
	VoiceChannelDestroy
)
