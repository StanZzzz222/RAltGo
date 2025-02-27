package alt_blip

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/blip_type"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
)

/*
Create by zyx
Date Time: 2024/9/12
File: blip.go
*/

func CreateBlipPoint(global bool, spriteId, color uint32, name string, position *models.Vector3) *models.IBlip {
	var w = lib.GetWrapper()
	var blip = &models.IBlip{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	ret, freePtrFunc := w.CreateBlip(blip_type.Point, spriteId, color, name, posData, posMetadata, 0, 0, 0)
	cBlip := entities.ConvertCBlip(ret)
	if cBlip != nil {
		freePtrFunc()
		blip = blip.NewIBlip(cBlip.ID, cBlip.BlipType, cBlip.SpriteId, cBlip.Color, cBlip.Name, cBlip.Rot, (*models.Vector3)(cBlip.Position))
		pools := models.GetPools()
		pools.PutBlip(blip)
		blip.SetGlobal(global)
		return blip
	}
	return nil
}

func CreateBlipPointSomePlayers(players []*models.IPlayer, spriteId, color uint32, name string, position *models.Vector3) *models.IBlip {
	var w = lib.GetWrapper()
	var blip = &models.IBlip{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	ret, freePtrFunc := w.CreateBlip(blip_type.Point, spriteId, color, name, posData, posMetadata, 0, 0, 0)
	cBlip := entities.ConvertCBlip(ret)
	if cBlip != nil {
		freePtrFunc()
		blip = blip.NewIBlip(cBlip.ID, cBlip.BlipType, cBlip.SpriteId, cBlip.Color, cBlip.Name, cBlip.Rot, (*models.Vector3)(cBlip.Position))
		pools := models.GetPools()
		pools.PutBlip(blip)
		blip.SetGlobal(false)
		blip.SetSomePlayers(players)
		return blip
	}
	return nil
}

func CreateBlipArea(global bool, color uint32, name string, position *models.Vector3, width, height float32) *models.IBlip {
	var w = lib.GetWrapper()
	var blip = &models.IBlip{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	ret, freePtrFunc := w.CreateBlip(blip_type.Area, 5, color, name, posData, posMetadata, width, height, 0)
	cBlip := entities.ConvertCBlip(ret)
	if cBlip != nil {
		freePtrFunc()
		blip = blip.NewIBlip(cBlip.ID, cBlip.BlipType, cBlip.SpriteId, cBlip.Color, cBlip.Name, cBlip.Rot, (*models.Vector3)(cBlip.Position))
		pools := models.GetPools()
		pools.PutBlip(blip)
		return blip
	}
	return nil
}

func CreateBlipAreaSomePlayers(players []*models.IPlayer, color uint32, name string, position *models.Vector3, width, height float32) *models.IBlip {
	var w = lib.GetWrapper()
	var blip = &models.IBlip{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	ret, freePtrFunc := w.CreateBlip(blip_type.Area, 5, color, name, posData, posMetadata, width, height, 0)
	cBlip := entities.ConvertCBlip(ret)
	if cBlip != nil {
		freePtrFunc()
		blip = blip.NewIBlip(cBlip.ID, cBlip.BlipType, cBlip.SpriteId, cBlip.Color, cBlip.Name, cBlip.Rot, (*models.Vector3)(cBlip.Position))
		pools := models.GetPools()
		pools.PutBlip(blip)
		blip.SetGlobal(false)
		blip.SetSomePlayers(players)
		return blip
	}
	return nil
}

func CreateBlipRadius(global bool, color uint32, name string, position *models.Vector3, radius float32, outline bool) *models.IBlip {
	var w = lib.GetWrapper()
	var blip = &models.IBlip{}
	var spriteId = 9
	if outline {
		spriteId = 10
	}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	ret, freePtrFunc := w.CreateBlip(blip_type.Radius, uint32(spriteId), color, name, posData, posMetadata, 0, 0, radius)
	cBlip := entities.ConvertCBlip(ret)
	if cBlip != nil {
		freePtrFunc()
		blip = blip.NewIBlip(cBlip.ID, cBlip.BlipType, cBlip.SpriteId, cBlip.Color, cBlip.Name, cBlip.Rot, (*models.Vector3)(cBlip.Position))
		pools := models.GetPools()
		pools.PutBlip(blip)
		blip.SetGlobal(global)
		return blip
	}
	return nil
}

func CreateBlipRadiusSomePlayers(players []*models.IPlayer, color uint32, name string, position *models.Vector3, radius float32, outline bool) *models.IBlip {
	var w = lib.GetWrapper()
	var blip = &models.IBlip{}
	var spriteId = 9
	if outline {
		spriteId = 10
	}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	ret, freePtrFunc := w.CreateBlip(blip_type.Radius, uint32(spriteId), color, name, posData, posMetadata, 0, 0, radius)
	cBlip := entities.ConvertCBlip(ret)
	if cBlip != nil {
		freePtrFunc()
		blip = blip.NewIBlip(cBlip.ID, cBlip.BlipType, cBlip.SpriteId, cBlip.Color, cBlip.Name, cBlip.Rot, (*models.Vector3)(cBlip.Position))
		pools := models.GetPools()
		pools.PutBlip(blip)
		blip.SetGlobal(false)
		blip.SetSomePlayers(players)
		return blip
	}
	return nil
}
