package models

import (
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"github.com/StanZzzz222/RAltGo/logger"
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
	var warpper = lib.GetWarpper()
	n.metaData.Store(key, value)
	mvalues := NewMValues(value)
	warpper.SetNetworkData(n.networdId, n.networkObjectType, enum.NetworkMeta, key, mvalues.Dump())
}

func (n *NetworkData) GetMetaData(key string) any {
	value, ok := n.metaData.Load(key)
	if ok {
		return value
	}
	return nil
}

func (n *NetworkData) HasMetaData(key string) bool {
	_, ok := n.metaData.Load(key)
	return ok
}

func (n *NetworkData) SetSyncedMetaData(key string, value any) {
	var warpper = lib.GetWarpper()
	n.syncedMetaData.Store(key, value)
	mvalues := NewMValues(value)
	warpper.SetNetworkData(n.networdId, n.networkObjectType, enum.NetworkSyncedMeta, key, mvalues.Dump())
}

func (n *NetworkData) GetSyncedMetaData(key string) any {
	value, ok := n.syncedMetaData.Load(key)
	if ok {
		return value
	}
	return nil
}

func (n *NetworkData) HasSyncedMetaData(key string) bool {
	_, ok := n.syncedMetaData.Load(key)
	return ok
}

func (n *NetworkData) SetStreamSyncedMetaData(key string, value any) {
	switch n.networkObjectType {
	case enum.Player, enum.Vehicle, enum.Ped:
		var warpper = lib.GetWarpper()
		n.streamSyncedMetaData.Store(key, value)
		mvalues := NewMValues(value)
		warpper.SetNetworkData(n.networdId, n.networkObjectType, enum.NetworkStreamSyncedMeta, key, mvalues.Dump())
		break
	default:
		logger.LogWarnf("ObjectType: %v does not support the SetStreamSyncedMetaData method", n.networkObjectType.String())
		break
	}
}

func (n *NetworkData) GetStreamSyncedMetaData(key string) any {
	value, ok := n.streamSyncedMetaData.Load(key)
	if ok {
		return value
	}
	return nil
}

func (n *NetworkData) HasStreamSyncedMetaData(key string) bool {
	_, ok := n.streamSyncedMetaData.Load(key)
	return ok
}
