package models

import (
	"github.com/goccy/go-json"
)

/*
   Create by zyx
   Date Time: 2024/9/25
   File: attach_entity_bone.go
*/

type AttachToEntityBoneIndex struct {
	OtherBoneIndex  uint16   `json:"other_bone_index"`
	MyBoneIndex     uint16   `json:"my_bone_index"`
	Pos             *Vector3 `json:"pos"`
	Rot             *Vector3 `json:"rot"`
	Collision       bool     `json:"collision"`
	NoFixedRotation bool     `json:"no_fixed_rotation"`
}

type AttachToEntityBoneName struct {
	OtherBoneName   string   `json:"other_bone_name"`
	MyBoneName      string   `json:"my_bone_name"`
	Pos             *Vector3 `json:"pos"`
	Rot             *Vector3 `json:"rot"`
	Collision       bool     `json:"collision"`
	NoFixedRotation bool     `json:"no_fixed_rotation"`
}

func NewAttachToEntityBoneIndex(otherBonIndex, myBoneIndex uint16, position, rotation *Vector3, collision, noFixedRotation bool) *AttachToEntityBoneIndex {
	return &AttachToEntityBoneIndex{
		OtherBoneIndex:  otherBonIndex,
		MyBoneIndex:     myBoneIndex,
		Pos:             position,
		Rot:             rotation,
		Collision:       collision,
		NoFixedRotation: noFixedRotation,
	}
}

func NewAttachToEntityBoneName(otherBoneName, myBoneName string, position, rotation *Vector3, collision, noFixedRotation bool) *AttachToEntityBoneName {
	return &AttachToEntityBoneName{
		OtherBoneName:   otherBoneName,
		MyBoneName:      myBoneName,
		Pos:             position,
		Rot:             rotation,
		Collision:       collision,
		NoFixedRotation: noFixedRotation,
	}
}

func (a *AttachToEntityBoneIndex) Dump() string {
	dumpBytes, _ := json.Marshal(&a)
	return string(dumpBytes)
}

func (a *AttachToEntityBoneName) Dump() string {
	dumpBytes, _ := json.Marshal(&a)
	return string(dumpBytes)
}
