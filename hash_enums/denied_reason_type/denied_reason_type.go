package denied_reason_type

/*
   Create by zyx
   Date Time: 2024/10/10
   File: denied_reason_type.go
*/

type DeniedReason uint8

const (
	WrongVersion DeniedReason = iota
	WrongBranch
	DebugNotAllowed
	WrongPassword
	WrongCdnUrl
)
