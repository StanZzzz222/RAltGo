package pools

import (
	"github.com/StanZzzz222/RAltGo/common/models"
)

/*
   Create by zyx
   Date Time: 2024/9/18
   File: blip.go
*/

func GetVirtualEntity(id uint32) *models.IVirtualEntity {
	return models.GetPools().GetVirtualEntity(id)
}

func GetVirtualEntitys() []*models.IVirtualEntity {
	var virtualEntitys []*models.IVirtualEntity
	var pools = models.GetPools().GetVirtualEntityPools()
	pools.Range(func(key, value any) bool {
		virtualEntitys = append(virtualEntitys, value.(*models.IVirtualEntity))
		return true
	})
	return virtualEntitys
}

func GetVirtualEntityIterator() <-chan *models.IVirtualEntity {
	var virtualEntityChan = make(chan *models.IVirtualEntity)
	var pools = models.GetPools().GetVirtualEntityPools()
	go func() {
		defer close(virtualEntityChan)
		pools.Range(func(key, value any) bool {
			virtualEntityChan <- value.(*models.IVirtualEntity)
			return true
		})
	}()
	return virtualEntityChan
}
