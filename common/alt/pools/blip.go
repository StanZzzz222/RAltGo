package pools

import "github.com/StanZzzz222/RAltGo/common/models"

/*
   Create by zyx
   Date Time: 2024/9/18
   File: blip.go
*/

func GetBlip(id uint32) *models.IBlip {
	return models.GetPools().GetBlip(id)
}

func GetBlips() []*models.IBlip {
	var blips []*models.IBlip
	var pools = models.GetPools().GetBlipPools()
	pools.Range(func(key, value any) bool {
		blips = append(blips, value.(*models.IBlip))
		return true
	})
	return blips
}

func GetBlipIterator() <-chan *models.IBlip {
	var blipChan = make(chan *models.IBlip)
	var pools = models.GetPools().GetBlipPools()
	go func() {
		defer close(blipChan)
		pools.Range(func(key, value any) bool {
			blipChan <- value.(*models.IBlip)
			return true
		})
	}()
	return blipChan
}
