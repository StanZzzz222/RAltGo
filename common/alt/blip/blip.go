package blip

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/enums/blip_type"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
)

/*
Create by zyx
Date Time: 2024/9/12
File: blip.go
*/

func CreateBlipPoint(spriteId, color uint32, name string, position *entities.Vector3) *models.IBlip {
	var w = &lib.Warrper{}
	var blip = &models.IBlip{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	ret, freePtrFunc := w.CreateBlip(blip_type.Point, spriteId, color, name, posData, posMetadata, 0, 0, 0)
	cBlip := entities.ConvertCBlip(ret)
	if cBlip != nil {
		freePtrFunc()
		blip = blip.NewIBlip(cBlip.ID, cBlip.BlipType, cBlip.SpriteId, cBlip.Color, cBlip.Name, cBlip.Rot, cBlip.Position)
		return blip
	}
	return nil
}

func CreateBlipArea(spriteId, color uint32, name string, position *entities.Vector3, width, height float32) *models.IBlip {
	var w = &lib.Warrper{}
	var blip = &models.IBlip{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	ret, freePtrFunc := w.CreateBlip(blip_type.Area, spriteId, color, name, posData, posMetadata, width, height, 0)
	cBlip := entities.ConvertCBlip(ret)
	if cBlip != nil {
		freePtrFunc()
		blip = blip.NewIBlip(cBlip.ID, cBlip.BlipType, cBlip.SpriteId, cBlip.Color, cBlip.Name, cBlip.Rot, cBlip.Position)
		return blip
	}
	return nil
}

func CreateBlipRadius(spriteId, color uint32, name string, position *entities.Vector3, radius float32) *models.IBlip {
	var w = &lib.Warrper{}
	var blip = &models.IBlip{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	ret, freePtrFunc := w.CreateBlip(blip_type.Radius, spriteId, color, name, posData, posMetadata, 0, 0, radius)
	cBlip := entities.ConvertCBlip(ret)
	if cBlip != nil {
		freePtrFunc()
		blip = blip.NewIBlip(cBlip.ID, cBlip.BlipType, cBlip.SpriteId, cBlip.Color, cBlip.Name, cBlip.Rot, cBlip.Position)
		return blip
	}
	return nil
}
