package alt_events

/*
   Create by zyx
   Date Time: 2024/9/22
   File: broadcast.go
*/

func BroadcastAll(message string) {
	EmitAllPlayer("chat:message", message)
}
