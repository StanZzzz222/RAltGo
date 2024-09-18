package pools

import "github.com/StanZzzz222/RAltGo/common/models"

/*
   Create by zyx
   Date Time: 2024/9/18
   File: ped.go
*/

func GetPed(id uint32) *models.IPed {
	return models.GetPools().GetPed(id)
}

func GetPeds() []*models.IPed {
	var peds []*models.IPed
	var pools = models.GetPools().GetPedPools()
	pools.Range(func(key, value any) bool {
		peds = append(peds, value.(*models.IPed))
		return true
	})
	return peds
}

func GetPedsIterator() <-chan *models.IPed {
	var pedsChan = make(chan *models.IPed)
	var pools = models.GetPools().GetPedPools()
	go func() {
		defer close(pedsChan)
		pools.Range(func(key, value any) bool {
			pedsChan <- value.(*models.IPed)
			return true
		})
	}()
	return pedsChan
}
