package alt_virtual_entity

import "C"
import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/29
   File: virtual_entity.go
*/

func CreateVirtualEntity(group *models.IVirtualEntityGroup, position *models.Vector3, streamingDistance uint32) *models.IVirtualEntity {
	var w = lib.GetWarpper()
	var virtualEntity = &models.IVirtualEntity{}
	posData, posMetaData := uint64(math.Float32bits(position.X))|(uint64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	ret, freePtrFunc := w.CreateVirtualEntity(group.GetId(), posData, posMetaData, streamingDistance)
	cVirtualEntity := entities.ConvertCVirtualEntity(ret)
	if cVirtualEntity != nil {
		freePtrFunc()
		virtualEntity = virtualEntity.NewIVirtualEntity(cVirtualEntity.ID, cVirtualEntity.StreamingDistance, (*models.Vector3)(cVirtualEntity.Position))
		pools := models.GetPools()
		pools.PutVirtualEntity(virtualEntity)
		return virtualEntity
	}
	return nil
}
