package models

import (
	"github.com/StanZzzz222/RAltGo/internal/enums"
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
	networkObjectType    enums.ObjectType
	datas                *sync.Map
	syncedMetaData       *sync.Map
	streamSyncedMetaData *sync.Map
}

func NewNetworkData(id uint32, objectType enums.ObjectType) *NetworkData {
	return &NetworkData{id, objectType, &sync.Map{}, &sync.Map{}, &sync.Map{}}
}

func (n *NetworkData) SetSyncedMetaData(key string, value any) {
	var warpper = lib.GetWarpper()
	n.syncedMetaData.Store(key, value)
	mvalues := NewMValues(value)
	warpper.SetNetworkData(n.networdId, n.networkObjectType, enums.NetworkSyncedMeta, key, mvalues.Dump())
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
	case enums.Player, enums.Vehicle, enums.Ped:
		var warpper = lib.GetWarpper()
		n.streamSyncedMetaData.Store(key, value)
		mvalues := NewMValues(value)
		warpper.SetNetworkData(n.networdId, n.networkObjectType, enums.NetworkStreamSyncedMeta, key, mvalues.Dump())
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

func (n *NetworkData) SetData(key string, value any) {
	n.datas.Store(key, value)
}

func (n *NetworkData) DelData(key string) {
	_, ok := n.datas.Load(key)
	if ok {
		n.datas.Delete(key)
	}
}

func (n *NetworkData) DelAllData() {
	n.datas.Range(func(key, value any) bool {
		n.datas.Delete(key)
		return true
	})
}

func (n *NetworkData) HasData(key string) bool {
	_, ok := n.datas.Load(key)
	if ok {
		return true
	}
	return false
}

func (n *NetworkData) GetData(key string) any {
	value, ok := n.datas.Load(key)
	if ok {
		return value
	}
	return value
}

func (n *NetworkData) GetDatas() []any {
	var datas []any
	n.datas.Range(func(key, value any) bool {
		datas = append(datas, key)
		return true
	})
	return datas
}
