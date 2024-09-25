package models

import (
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"github.com/StanZzzz222/RAltGo/logger"
)

/*
   Create by zyx
   Date Time: 2024/9/25
   File: entity_data.go
*/

type EntityData struct {
	entityId         uint32
	entityObjectType enum.ObjectType
}

func NewEntityData(id uint32, objectType enum.ObjectType) *EntityData {
	return &EntityData{id, objectType}
}

func (e *EntityData) GetNetOwner() *IPlayer {
	return pools.GetPlayer(e.entityId)
}

func (e *EntityData) SetNetOwner(owner *IPlayer, disableMigration bool) {
	switch e.entityObjectType {
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
	}
}
