package pools

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/18
   File: ped_pools.go
*/

var pedPools = newPedPool()

type PedPool struct {
	peds *sync.Map
}

func newPedPool() *PedPool {
	return &PedPool{
		peds: &sync.Map{},
	}
}

func (p *PedPool) Put(ped *models.IPed) {
	if _, ok := p.peds.Load(ped.GetId()); !ok {
		p.peds.Store(ped.GetId(), ped)
	}
}

func (p *PedPool) Remove(ped *models.IPed) {
	if _, ok := p.peds.Load(ped.GetId()); ok {
		p.peds.Delete(ped.GetId())
	}
}

func GetPed(id uint32) *models.IPed {
	ped, ok := pedPools.peds.Load(id)
	if ok {
		return ped.(*models.IPed)
	}
	return nil
}

func GetPeds() []*models.IPed {
	var peds []*models.IPed
	pedPools.peds.Range(func(key, value any) bool {
		peds = append(peds, value.(*models.IPed))
		return true
	})
	return peds
}

func GetPedsIterator() <-chan *models.IPed {
	pedsChan := make(chan *models.IPed)
	go func() {
		defer close(pedsChan)
		playerPools.players.Range(func(key, value any) bool {
			pedsChan <- value.(*models.IPed)
			return true
		})
	}()
	return pedsChan
}

func GetPedPools() *PedPool {
	return pedPools
}
