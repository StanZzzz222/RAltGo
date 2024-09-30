package models

import (
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enums"
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
	entityObjectType enums.ObjectType
}

func NewEntityData(id uint32, objectType enums.ObjectType) *EntityData {
	return &EntityData{id, objectType}
}

func (e *EntityData) GetNetOwner() *IPlayer {
	switch e.entityObjectType {
	case enums.Player, enums.Vehicle, enums.Ped, enums.Object:
		ret, freeEntityDataFunc := w.GetEntityData(e.entityId, e.entityObjectType, enums.NetOwner)
		res := entities.ConverCDataResult(ret)
		if res != nil {
			defer freeEntityDataFunc()
			return pools.GetPlayer(res.U32Val)
		}
		return nil
	default:
		logger.Logger().LogWarnf("ObjectType: %v does not support the GetNetOwner method", e.entityObjectType.String())
		break
	}
	return nil
}

func (e *EntityData) GetSyncId() SyncId {
	switch e.entityObjectType {
	case enums.Player, enums.Vehicle, enums.Ped, enums.Object:
		ret, freeEntityDataFunc := w.GetEntityData(e.entityId, e.entityObjectType, enums.SyncId)
		res := entities.ConverCDataResult(ret)
		if res != nil {
			defer freeEntityDataFunc()
			return SyncId(res.U32Val)
		}
	default:
		logger.Logger().LogWarnf("ObjectType: %v does not support the GetNetOwner method", e.entityObjectType.String())
		break
	}
	return SyncId(0)
}

func (e *EntityData) GetStreamed() bool {
	switch e.entityObjectType {
	case enums.Player, enums.Vehicle, enums.Ped, enums.Object:
		ret, freeEntityDataFunc := w.GetEntityData(e.entityId, e.entityObjectType, enums.Streamed)
		res := entities.ConverCDataResult(ret)
		if res != nil {
			defer freeEntityDataFunc()
			return res.BoolVal
		}
	default:
		logger.Logger().LogWarnf("ObjectType: %v does not support the GetStreamed method", e.entityObjectType.String())
		break
	}
	return false
}

func (e *EntityData) GetStreamingDistance() uint32 {
	switch e.entityObjectType {
	case enums.Player, enums.Vehicle, enums.Ped, enums.Object:
		ret, freeEntityDataFunc := w.GetEntityData(e.entityId, e.entityObjectType, enums.Streamed)
		res := entities.ConverCDataResult(ret)
		if res != nil {
			defer freeEntityDataFunc()
			return res.U32Val
		}
	default:
		logger.Logger().LogWarnf("ObjectType: %v does not support the GetStreamingDistance method", e.entityObjectType.String())
		break
	}
	return 0
}

func (e *EntityData) SetNetOwner(owner *IPlayer, disableMigration bool) {
	switch e.entityObjectType {
	case enums.Player, enums.Vehicle, enums.Ped, enums.Object:
		value := 0
		if disableMigration {
			value = 1
		}
		w.SetEntityData(e.entityId, e.entityObjectType, enums.NetOwner, 0, uint64(owner.GetId()), uint32(value), "")
	default:
		logger.Logger().LogWarnf("ObjectType: %v does not support the SetNetOwner method", e.entityObjectType.String())
		break
	}
}

func (e *EntityData) SetStreamed(streamed bool) {
	switch e.entityObjectType {
	case enums.Player, enums.Vehicle, enums.Ped, enums.Object:
		value := 0
		if streamed {
			value = 1
		}
		w.SetEntityData(e.entityId, e.entityObjectType, enums.Streamed, 0, uint64(value), 0, "")
	default:
		logger.Logger().LogWarnf("ObjectType: %v does not support the SetStreamed method", e.entityObjectType.String())
		break
	}
}

func (e *EntityData) SetStreamingDistance(streamingDistance uint32) {
	switch e.entityObjectType {
	case enums.Player, enums.Vehicle, enums.Ped, enums.Object:
		w.SetEntityData(e.entityId, e.entityObjectType, enums.StreamingDistance, 0, 0, streamingDistance, "")
	default:
		logger.Logger().LogWarnf("ObjectType: %v does not support the SetStreamingDistance method", e.entityObjectType.String())
		break
	}
}

func (e *EntityData) Detach() {
	switch e.entityObjectType {
	case enums.Player, enums.Vehicle, enums.Ped, enums.Object:
		w.SetEntityData(e.entityId, e.entityObjectType, enums.Detach, 0, 0, 0, "")
	default:
		logger.Logger().LogWarnf("ObjectType: %v does not support the Detach method", e.entityObjectType.String())
		break
	}
}

func (e *EntityData) AttachToEntityBoneName(targetEntity any, bone *AttachToEntityBoneName) {
	if ok, entityType, id := checkSupport(targetEntity); ok {
		w.SetEntityData(e.entityId, e.entityObjectType, enums.AttachToEntityBoneName, entityType, uint64(id), 0, bone.Dump())
		return
	}
	logger.Logger().LogWarnf("ObjectType: %v does not support the AttachToEntityBoneName method", e.entityObjectType.String())
}

func (e *EntityData) AttachToEntityBoneIndex(targetEntity any, boneIndex *AttachToEntityBoneIndex) {
	if ok, entityType, id := checkSupport(targetEntity); ok {
		w.SetEntityData(e.entityId, e.entityObjectType, enums.AttachToEntityBoneIndex, entityType, uint64(id), 0, boneIndex.Dump())
		return
	}
	logger.Logger().LogWarnf("ObjectType: %v does not support the AttachToEntityBoneName method", e.entityObjectType.String())
}

func checkSupport(targetEntity any) (bool, enums.ObjectType, uint32) {
	var res = false
	var entityType = enums.ObjectType(0)
	var id uint32 = 0
	t := reflect.TypeOf(targetEntity)
	if t.Kind() == reflect.Ptr {
		elemType := t.Elem()
		switch elemType {
		case reflect.TypeOf((*IPlayer)(nil)).Elem():
			res = true
			entityType = enums.Player
			id = targetEntity.(*IPlayer).GetId()
			break
		case reflect.TypeOf((*IVehicle)(nil)).Elem():
			res = true
			entityType = enums.Vehicle
			id = targetEntity.(*IVehicle).GetId()
			break
		case reflect.TypeOf((*IPed)(nil)).Elem():
			res = true
			entityType = enums.Ped
			id = targetEntity.(*IPed).GetId()
			break
		case reflect.TypeOf((*IObject)(nil)).Elem():
			res = true
			entityType = enums.Object
			id = targetEntity.(*IObject).GetId()
			break
		}
	}
	return res, entityType, id
}
