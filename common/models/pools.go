package models

import "sync"

/*
   Create by zyx
   Date Time: 2024/9/18
   File: pools.go
*/

var pools = newPools()

type Pools struct {
	blips    *sync.Map
	vehicles *sync.Map
	players  *sync.Map
	peds     *sync.Map
}

func newPools() *Pools {
	return &Pools{
		blips:    &sync.Map{},
		vehicles: &sync.Map{},
		players:  &sync.Map{},
		peds:     &sync.Map{},
	}
}

func (p *Pools) PutBlip(blip *IBlip) {
	if _, ok := p.blips.Load(blip.GetId()); !ok {
		p.blips.Store(blip.GetId(), blip)
	}
}

func (p *Pools) PutPlayer(player *IPlayer) {
	if _, ok := p.players.Load(player.GetId()); !ok {
		p.players.Store(player.GetId(), player)
	}
}

func (p *Pools) PutVehicle(vehicle *IVehicle) {
	if _, ok := p.vehicles.Load(vehicle.GetId()); !ok {
		p.vehicles.Store(vehicle.GetId(), vehicle)
	}
}

func (p *Pools) PutPed(ped *IPed) {
	if _, ok := p.peds.Load(ped.GetId()); !ok {
		p.peds.Store(ped.GetId(), ped)
	}
}

func (p *Pools) DestroyBlip(blip *IBlip) {
	if _, ok := p.blips.Load(blip.GetId()); ok {
		p.blips.Delete(blip.GetId())
	}
}

func (p *Pools) DestroyPlayer(player *IPlayer) {
	if _, ok := p.players.Load(player.GetId()); ok {
		p.players.Delete(player.GetId())
	}
}

func (p *Pools) DestroyVehicle(vehicle *IVehicle) {
	if _, ok := p.vehicles.Load(vehicle.GetId()); ok {
		p.vehicles.Delete(vehicle.GetId())
	}
}

func (p *Pools) DestroyPed(ped *IPed) {
	if _, ok := p.peds.Load(ped.GetId()); ok {
		p.peds.Delete(ped.GetId())
	}
}

func GetPools() *Pools {
	return pools
}
