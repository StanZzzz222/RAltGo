package pools

import (
	"github.com/StanZzzz222/RAltGo/common/models"
)

/*
   Create by zyx
   Date Time: 2024/9/18
   File: checkpoint.go
*/

func GetColshape(id uint32) *models.IColshape {
	return models.GetPools().GetColshape(id)
}

func GetColshapes() []*models.IColshape {
	var colshapes []*models.IColshape
	var pools = models.GetPools().GetColshapePools()
	pools.Range(func(key, value any) bool {
		colshapes = append(colshapes, value.(*models.IColshape))
		return true
	})
	return colshapes
}

func GetColshapeIterator() <-chan *models.IColshape {
	var colshapeChan = make(chan *models.IColshape)
	var pools = models.GetPools().GetColshapePools()
	go func() {
		defer close(colshapeChan)
		pools.Range(func(key, value any) bool {
			colshapeChan <- value.(*models.IColshape)
			return true
		})
	}()
	return colshapeChan
}
