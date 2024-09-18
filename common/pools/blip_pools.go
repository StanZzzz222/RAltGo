package pools

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/18
   File: blip_pools.go
*/

var blipPools = newBlipPool()

type BlipPool struct {
	blips *sync.Map
}

func newBlipPool() *BlipPool {
	return &BlipPool{
		blips: &sync.Map{},
	}
}

func (p *BlipPool) Put(blip *models.IBlip) {
	if _, ok := p.blips.Load(blip.GetId()); !ok {
		p.blips.Store(blip.GetId(), blip)
	}
}

func (p *BlipPool) Remove(blip *models.IBlip) {
	if _, ok := p.blips.Load(blip.GetId()); ok {
		p.blips.Delete(blip.GetId())
	}
}

func GetBlip(id uint32) *models.IBlip {
	blip, ok := blipPools.blips.Load(id)
	if ok {
		return blip.(*models.IBlip)
	}
	return nil
}

func GetBlips() []*models.IBlip {
	var blips []*models.IBlip
	blipPools.blips.Range(func(key, value any) bool {
		blips = append(blips, value.(*models.IBlip))
		return true
	})
	return blips
}

func GetBlipIterator() <-chan *models.IBlip {
	blipChan := make(chan *models.IBlip)
	go func() {
		defer close(blipChan)
		blipPools.blips.Range(func(key, value any) bool {
			blipChan <- value.(*models.IBlip)
			return true
		})
	}()
	return blipChan
}

func GetBlipPools() *BlipPool {
	return blipPools
}
