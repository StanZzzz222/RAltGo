package enums

/*
   Create by zyx
   Date Time: 2024/10/7
   File: voice_connection_state_type.go
*/

type VoiceConnectionState = uint

const (
	Disconnected VoiceConnectionState = iota
	Connecting
	Connected
)
