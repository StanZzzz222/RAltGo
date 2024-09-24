package checkpoint

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/check_point_type"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/25
   File: checkpoint.go
*/

func CreateCheckpoint(checkPointType check_point_type.CheckPointType, position *entities.Vector3, radius, height float32, color *entities.Rgba, streamingDistance uint32) *models.ICheckpoint {
	var w = lib.GetWarpper()
	var c = &models.ICheckpoint{}
	posData, posMetaData := uint64(math.Float32bits(position.X))|(uint64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	ret, freePtrFunc := w.CreateCheckpoint(uint8(checkPointType), posData, posMetaData, radius, height, color.R, color.G, color.B, color.A, streamingDistance)
	cCheckpoint := entities.ConvertCCheckPoint(ret)
	if cCheckpoint != nil {
		freePtrFunc()
		c = c.NewICheckPoint(cCheckpoint.ID, cCheckpoint.CheckpointType, cCheckpoint.Position, height, radius)
		pools := models.GetPools()
		pools.PutCheckpoint(c)
		return c
	}
	return nil
}
