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
	logger.LogInfof("TODO: not implement")
}

func (e *EntityData) SetStreamed(streamed bool) {
	logger.LogInfof("TODO: not implement")
}

func (e *EntityData) SetStreamingDistance(streamingDistance uint32) {
	logger.LogInfof("TODO: not implement")
}

func (e *EntityData) GetStreamed(streamed bool) {
	logger.LogInfof("TODO: not implement")
}

func (e *EntityData) GetStreamingDistance(streamingDistance uint32) {
	logger.LogInfof("TODO: not implement")
}

func (e *EntityData) Detach() {
	logger.LogInfof("TODO: not implement")
}

func (e *EntityData) AttachToEntityBoneName(targetEntity any, bone *AttachToEntityBoneName) {
	logger.LogInfof("TODO: not implement")
}

func (e *EntityData) AttachToEntityBoneIndex(targetEntity any, boneIndex *AttachToEntityBoneIndex) {
	logger.LogInfof("TODO: not implement")
}
