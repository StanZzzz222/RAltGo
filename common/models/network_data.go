package models

import (
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"github.com/StanZzzz222/RAltGo/logger"
	"strings"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/25
   File: network_data.go
*/

type NetworkData struct {
	networdId            uint32
	networkObjectType    enum.ObjectType
	metaData             *sync.Map
	syncedMetaData       *sync.Map
	streamSyncedMetaData *sync.Map
}

func NewNetworkData(id uint32, objectType enum.ObjectType) *NetworkData {
	return &NetworkData{id, objectType, &sync.Map{}, &sync.Map{}, &sync.Map{}}
}

func (n *NetworkData) SetMetaData(key string, value any) {
	var keys []string
	var datas []any
	var warpper = lib.GetWarpper()
	n.metaData.Store(key, value)
	n.metaData.Range(func(key, value any) bool {
		keys = append(keys, key.(string))
		datas = append(datas, value)
		return true
	})
	mvalues := NewMValues(datas...)
	warpper.SetNetworkData(n.networdId, n.networkObjectType, enum.NetworkMeta, strings.Join(keys, ","), mvalues.Dump())
}

func (n *NetworkData) SetSyncedMetaData(key string, value any) {
	var keys []string
	var datas []any
	var warpper = lib.GetWarpper()
	n.metaData.Store(key, value)
	n.metaData.Range(func(key, value any) bool {
		keys = append(keys, key.(string))
		datas = append(datas, value)
		return true
	})
	mvalues := NewMValues(datas...)
	warpper.SetNetworkData(n.networdId, n.networkObjectType, enum.NetworkSyncedMeta, strings.Join(keys, ","), mvalues.Dump())
}

func (n *NetworkData) SetStreamSyncedMetaData(key string, value any) {
	switch n.networkObjectType {
	case enum.Player, enum.Vehicle, enum.Ped:
		var keys []string
		var datas []any
		var warpper = lib.GetWarpper()
		n.metaData.Store(key, value)
		n.metaData.Range(func(key, value any) bool {
			keys = append(keys, key.(string))
			datas = append(datas, value)
			return true
		})
		mvalues := NewMValues(datas...)
		warpper.SetNetworkData(n.networdId, n.networkObjectType, enum.NetworkStreamSyncedMeta, strings.Join(keys, ","), mvalues.Dump())
		break
	default:
		logger.LogWarnf("ObjectType: %v does not support the SetStreamSyncedMetaData method", n.networkObjectType.String())
		break
	}
}
