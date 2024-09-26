package models

import (
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"github.com/StanZzzz222/RAltGo/logger"
	"reflect"
)

/*
   Create by zyx
   Date Time: 2024/9/25
   File: entity_data.go
*/

type SyncId uint16

type EntityData struct {
	entityId         uint32
	entityObjectType enum.ObjectType
}

func NewEntityData(id uint32, objectType enum.ObjectType) *EntityData {
	return &EntityData{id, objectType}
}

func (e *EntityData) GetNetOwner() *IPlayer {
	switch e.entityObjectType {
	case enum.Player, enum.Vehicle, enum.Ped, enum.Object:
		ret, freeEntityDataFunc := w.GetEntityData(e.entityId, e.entityObjectType, enum.NetOwner)
		res := entities.ConverCDataResult(ret)
		if res != nil {
			defer freeEntityDataFunc()
			return pools.GetPlayer(res.U32Val)
		}
		return nil
	default:
		logger.LogWarnf("ObjectType: %v does not support the GetNetOwner method", e.entityObjectType.String())
		break
	}
	return nil
}

func (e *EntityData) GetSyncId() SyncId {
	switch e.entityObjectType {
	case enum.Player, enum.Vehicle, enum.Ped, enum.Object:
		ret, freeEntityDataFunc := w.GetEntityData(e.entityId, e.entityObjectType, enum.SyncId)
		res := entities.ConverCDataResult(ret)
		if res != nil {
			defer freeEntityDataFunc()
			return SyncId(res.U32Val)
		}
	default:
		logger.LogWarnf("ObjectType: %v does not support the GetNetOwner method", e.entityObjectType.String())
		break
	}
	return SyncId(0)
}

func (e *EntityData) GetStreamed() bool {
	switch e.entityObjectType {
	case enum.Player, enum.Vehicle, enum.Ped, enum.Object:
		ret, freeEntityDataFunc := w.GetEntityData(e.entityId, e.entityObjectType, enum.Streamed)
		res := entities.ConverCDataResult(ret)
		if res != nil {
			defer freeEntityDataFunc()
			return res.BoolVal
		}
	default:
		logger.LogWarnf("ObjectType: %v does not support the GetStreamed method", e.entityObjectType.String())
		break
	}
	return false
}

func (e *EntityData) GetStreamingDistance() uint32 {
	switch e.entityObjectType {
	case enum.Player, enum.Vehicle, enum.Ped, enum.Object:
		ret, freeEntityDataFunc := w.GetEntityData(e.entityId, e.entityObjectType, enum.Streamed)
		res := entities.ConverCDataResult(ret)
		if res != nil {
			defer freeEntityDataFunc()
			return res.U32Val
		}
	default:
		logger.LogWarnf("ObjectType: %v does not support the GetStreamingDistance method", e.entityObjectType.String())
		break
	}
	return 0
}

func (e *EntityData) SetNetOwner(owner *IPlayer, disableMigration bool) {
	switch e.entityObjectType {
	case enum.Player, enum.Vehicle, enum.Ped, enum.Object:
		value := 0
		if disableMigration {
			value = 1
		}
		w.SetEntityData(e.entityId, e.entityObjectType, enum.NetOwner, 0, uint64(owner.GetId()), uint32(value), "")
	default:
		logger.LogWarnf("ObjectType: %v does not support the SetNetOwner method", e.entityObjectType.String())
		break
	}
}

func (e *EntityData) SetStreamed(streamed bool) {
	switch e.entityObjectType {
	case enum.Player, enum.Vehicle, enum.Ped, enum.Object:
		value := 0
		if streamed {
			value = 1
		}
		w.SetEntityData(e.entityId, e.entityObjectType, enum.Streamed, 0, uint64(value), 0, "")
	default:
		logger.LogWarnf("ObjectType: %v does not support the SetStreamed method", e.entityObjectType.String())
		break
	}
}

func (e *EntityData) SetStreamingDistance(streamingDistance uint32) {
	switch e.entityObjectType {
	case enum.Player, enum.Vehicle, enum.Ped, enum.Object:
		w.SetEntityData(e.entityId, e.entityObjectType, enum.StreamingDistance, 0, 0, streamingDistance, "")
	default:
		logger.LogWarnf("ObjectType: %v does not support the SetStreamingDistance method", e.entityObjectType.String())
		break
	}
}

func (e *EntityData) Detach() {
	switch e.entityObjectType {
	case enum.Player, enum.Vehicle, enum.Ped, enum.Object:
		w.SetEntityData(e.entityId, e.entityObjectType, enum.Detach, 0, 0, 0, "")
	default:
		logger.LogWarnf("ObjectType: %v does not support the Detach method", e.entityObjectType.String())
		break
	}
}

func (e *EntityData) AttachToEntityBoneName(targetEntity any, bone *AttachToEntityBoneName) {
	if ok, entityType, id := checkSupport(targetEntity); ok {
		w.SetEntityData(e.entityId, e.entityObjectType, enum.AttachToEntityBoneName, entityType, uint64(id), 0, bone.Dump())
		return
	}
	logger.LogWarnf("ObjectType: %v does not support the AttachToEntityBoneName method", e.entityObjectType.String())
}

func (e *EntityData) AttachToEntityBoneIndex(targetEntity any, boneIndex *AttachToEntityBoneIndex) {
	if ok, entityType, id := checkSupport(targetEntity); ok {
		w.SetEntityData(e.entityId, e.entityObjectType, enum.AttachToEntityBoneIndex, entityType, uint64(id), 0, boneIndex.Dump())
		return
	}
	logger.LogWarnf("ObjectType: %v does not support the AttachToEntityBoneName method", e.entityObjectType.String())
}

func checkSupport(targetEntity any) (bool, enum.ObjectType, uint32) {
	var res = false
	var entityType = enum.ObjectType(0)
	var id uint32 = 0
	t := reflect.TypeOf(targetEntity)
	if t.Kind() == reflect.Ptr {
		elemType := t.Elem()
		switch elemType {
		case reflect.TypeOf((*IPlayer)(nil)).Elem():
			res = true
			entityType = enum.Player
			id = targetEntity.(*IPlayer).GetId()
			break
		case reflect.TypeOf((*IVehicle)(nil)).Elem():
			res = true
			entityType = enum.Vehicle
			id = targetEntity.(*IVehicle).GetId()
			break
		case reflect.TypeOf((*IPed)(nil)).Elem():
			res = true
			entityType = enum.Ped
			id = targetEntity.(*IPed).GetId()
			break
		case reflect.TypeOf((*IObject)(nil)).Elem():
			res = true
			entityType = enum.Object
			id = targetEntity.(*IObject).GetId()
			break
		}
	}
	return res, entityType, id
}
