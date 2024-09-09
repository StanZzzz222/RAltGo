package utils

import "strings"

/*
   Create by zyx
   Date Time: 2024/9/5
   File: common.go
*/

func Hash(model string) uint32 {
	k := strings.ToLower(model)
	var h uint32
	var i int
	for i = 0; i < len(k); i++ {
		h += uint32(k[i])
		h += h << 10
		h ^= h >> 6
	}
	h += h << 3
	h ^= h >> 11
	h += h << 15
	return h
}
