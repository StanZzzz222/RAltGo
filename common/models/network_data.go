package models

import (
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"github.com/StanZzzz222/RAltGo/logger"
)

/*
   Create by zyx
   Date Time: 2024/9/25
   File: network_data.go
*/

type NetworkData struct {
	networdId         uint32
	networkObjectType enum.ObjectType
}

func NewNetworkData(id uint32, objectType enum.ObjectType) *NetworkData {
	return &NetworkData{id, objectType}
}

func (n *NetworkData) SetMetaData(datas map[string]any) {
	switch n.networkObjectType {
	case enum.Player:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Vehicle:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Ped:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Blip:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Colshape:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Object:
		logger.LogInfof("TODO: not implement")
		break
	case enum.CheckPoint:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Marker:
		logger.LogInfof("TODO: not implement")
		break
	default:
		logger.LogWarnf("Network objectType %d not support", n.networkObjectType)
	}
}

func (n *NetworkData) SetSyncedMetaData(datas map[string]any) {
	switch n.networkObjectType {
	case enum.Player:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Vehicle:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Ped:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Blip:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Colshape:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Object:
		logger.LogInfof("TODO: not implement")
		break
	case enum.CheckPoint:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Marker:
		logger.LogInfof("TODO: not implement")
		break
	default:
		logger.LogWarnf("Network objectType %d not support", n.networkObjectType)
	}
}

func (n *NetworkData) SetStreamSyncedMetaData(datas map[string]any) {
	switch n.networkObjectType {
	case enum.Player:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Vehicle:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Ped:
		logger.LogInfof("TODO: not implement")
		break
	case enum.Object:
		logger.LogInfof("TODO: not implement")
		break
	case enum.CheckPoint:
		logger.LogInfof("TODO: not implement")
		break
	default:
		logger.LogWarnf("Network objectType %d not support", n.networkObjectType)
	}
}
