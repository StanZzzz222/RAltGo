package blip

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/enums/blip_type"
	"github.com/StanZzzz222/RAltGo/internal/entitys"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/12
   File: blip.go
*/

var rw = &sync.RWMutex{}
var w = &lib.Warrper{}

func CreateBlipPoint(spriteId, color uint32, name string, position *entitys.Vector3) *models.IBlip {
	rw.Lock()
	defer rw.Unlock()
	var blip = &models.IBlip{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	ret := w.CreateBlip(blip_type.Point, spriteId, color, posData, posMetadata, 0, 0, 0)
	cBlip := entitys.ConvertCBlip(ret)
	if cBlip != nil {
		blip = blip.NewIBlip(cBlip.ID, cBlip.BlipType, cBlip.SpriteId, cBlip.Color, cBlip.Rot, cBlip.Position)
		blip.SetBlipName(name)
		return blip
	}
	return nil
}

func CreateBlipArea(spriteId, color uint32, name string, position *entitys.Vector3, width, height float32) *models.IBlip {
	rw.Lock()
	defer rw.Unlock()
	var blip = &models.IBlip{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	ret := w.CreateBlip(blip_type.Area, spriteId, color, posData, posMetadata, width, height, 0)
	cBlip := entitys.ConvertCBlip(ret)
	if cBlip != nil {
		blip = blip.NewIBlip(cBlip.ID, cBlip.BlipType, cBlip.SpriteId, cBlip.Color, cBlip.Rot, cBlip.Position)
		blip.SetBlipName(name)
		return blip
	}
	return nil
}

func CreateBlipRadius(spriteId, color uint32, name string, position *entitys.Vector3, radius float32) *models.IBlip {
	rw.Lock()
	defer rw.Unlock()
	var blip = &models.IBlip{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	ret := w.CreateBlip(blip_type.Radius, spriteId, color, posData, posMetadata, 0, 0, radius)
	cBlip := entitys.ConvertCBlip(ret)
	if cBlip != nil {
		blip = blip.NewIBlip(cBlip.ID, cBlip.BlipType, cBlip.SpriteId, cBlip.Color, cBlip.Rot, cBlip.Position)
		blip.SetBlipName(name)
		return blip
	}
	return nil
}
