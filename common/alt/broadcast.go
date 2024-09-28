package alt

import "github.com/StanZzzz222/RAltGo/common/alt/alt_events"

/*
   Create by zyx
   Date Time: 2024/9/22
   File: broadcast.go
*/

func SendBroadcast(message string) {
	alt_events.EmitAllPlayer("chat:message", "", message)
}
