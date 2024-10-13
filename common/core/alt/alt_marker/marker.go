package alt_marker

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/marker_type"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/25
   File: marker.go
*/

func CreateMarker(markerType marker_type.MarkerType, position *models.Vector3, color *models.Rgba) *models.IMarker {
	var w = lib.GetWarpper()
	var m = &models.IMarker{}
	posData, posMetaData := uint64(math.Float32bits(position.X))|(uint64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	ret, freePtrFunc := w.CreateMarker(uint8(markerType), posData, posMetaData, color.R, color.G, color.B, color.A)
	cMarker := entities.ConvertCMarker(ret)
	if cMarker != nil {
		freePtrFunc()
		m = m.NewIMarker(cMarker.ID, cMarker.MarkerType, (*models.Vector3)(cMarker.Position))
		pools := models.GetPools()
		pools.PutMarker(m)
		return m
	}
	return nil
}
