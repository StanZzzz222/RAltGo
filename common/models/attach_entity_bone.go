package models

import "github.com/StanZzzz222/RAltGo/internal/entities"

/*
   Create by zyx
   Date Time: 2024/9/25
   File: attach_entity_bone.go
*/

type AttachToEntityBoneIndex struct {
	otherBoneIndex  uint16
	myBoneIndex     uint16
	pos             *entities.Vector3
	rot             *entities.Vector3
	collision       bool
	noFixedRotation bool
}

type AttachToEntityBoneName struct {
	otherBoneName   string
	myBoneName      string
	pos             *entities.Vector3
	rot             *entities.Vector3
	collision       bool
	noFixedRotation bool
}

func NewAttachToEntityBoneIndex(otherBonIndex, myBoneIndex uint16, position, rotation *entities.Vector3, collision, noFixedRotation bool) *AttachToEntityBoneIndex {
	return &AttachToEntityBoneIndex{
		otherBoneIndex:  otherBonIndex,
		myBoneIndex:     myBoneIndex,
		pos:             position,
		rot:             rotation,
		collision:       collision,
		noFixedRotation: noFixedRotation,
	}
}

func NewAttachToEntityBoneName(otherBoneName, myBoneName string, position, rotation *entities.Vector3, collision, noFixedRotation bool) *AttachToEntityBoneName {
	return &AttachToEntityBoneName{
		otherBoneName:   otherBoneName,
		myBoneName:      myBoneName,
		pos:             position,
		rot:             rotation,
		collision:       collision,
		noFixedRotation: noFixedRotation,
	}
}
