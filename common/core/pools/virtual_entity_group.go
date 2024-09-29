package pools

import "github.com/StanZzzz222/RAltGo/common/models"

/*
   Create by zyx
   Date Time: 2024/9/18
   File: blip.go
*/

func GetVirtualEntityGroup(id uint32) *models.IVirtualEntityGroup {
	return models.GetPools().GetVirtualEntityGroup(id)
}

func GetVirtualEntityGroups() []*models.IVirtualEntityGroup {
	var virtualEntityGroups []*models.IVirtualEntityGroup
	var pools = models.GetPools().GetVirtualEntityGroupPools()
	pools.Range(func(key, value any) bool {
		virtualEntityGroups = append(virtualEntityGroups, value.(*models.IVirtualEntityGroup))
		return true
	})
	return virtualEntityGroups
}

func GetVirtualEntityGroupIterator() <-chan *models.IVirtualEntityGroup {
	var virtualEntityGroupChan = make(chan *models.IVirtualEntityGroup)
	var pools = models.GetPools().GetVirtualEntityGroupPools()
	go func() {
		defer close(virtualEntityGroupChan)
		pools.Range(func(key, value any) bool {
			virtualEntityGroupChan <- value.(*models.IVirtualEntityGroup)
			return true
		})
	}()
	return virtualEntityGroupChan
}
