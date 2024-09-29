package alt_vehicle

import "C"
import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/lib"
)

/*
   Create by zyx
   Date Time: 2024/9/29
   File: virtual_entity_group.go
*/

func CreateVirtualEntityGroup(maxEntitiesInStream uint32) *models.IVirtualEntityGroup {
	var w = lib.GetWarpper()
	var virtualEntityGroup = &models.IVirtualEntityGroup{}
	ret, freePtrFunc := w.CreateVirtualEntityGroup(maxEntitiesInStream)
	cVirtualEntityGroup := entities.ConvertCVirtualEntityGroup(ret)
	if cVirtualEntityGroup != nil {
		freePtrFunc()
		virtualEntityGroup = virtualEntityGroup.NewIVirtualEntityGroup(cVirtualEntityGroup.ID, cVirtualEntityGroup.MaxEntitiesInStream)
		pools := models.GetPools()
		pools.PutVirtualEntityGroup(virtualEntityGroup)
		return virtualEntityGroup
	}
	return nil
}
