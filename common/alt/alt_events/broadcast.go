package alt_events

/*
   Create by zyx
   Date Time: 2024/9/22
   File: broadcast.go
*/

func SendBroadcast(message string) {
	EmitAllPlayer("chat:message", "", message)
}
