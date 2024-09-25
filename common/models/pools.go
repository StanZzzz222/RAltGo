package models

import "sync"

/*
   Create by zyx
   Date Time: 2024/9/18
   File: pools.go
*/

var pools = newPools()

type Pools struct {
	blips       *sync.Map
	vehicles    *sync.Map
	players     *sync.Map
	peds        *sync.Map
	colshapes   *sync.Map
	checkpoints *sync.Map
	markers     *sync.Map
	objects     *sync.Map
}

func newPools() *Pools {
	return &Pools{
		blips:       &sync.Map{},
		vehicles:    &sync.Map{},
		players:     &sync.Map{},
		peds:        &sync.Map{},
		colshapes:   &sync.Map{},
		checkpoints: &sync.Map{},
		markers:     &sync.Map{},
		objects:     &sync.Map{},
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

func (p *Pools) PutColshape(colshape *IColshape) {
	if _, ok := p.colshapes.Load(colshape.GetId()); !ok {
		p.colshapes.Store(colshape.GetId(), colshape)
	}
}

func (p *Pools) PutCheckpoint(checkpoint *ICheckpoint) {
	if _, ok := p.checkpoints.Load(checkpoint.GetId()); !ok {
		p.checkpoints.Store(checkpoint.GetId(), checkpoint)
	}
}

func (p *Pools) PutMarker(marker *IMarker) {
	if _, ok := p.markers.Load(marker.GetId()); !ok {
		p.markers.Store(marker.GetId(), marker)
	}
}

func (p *Pools) PutObject(object *IObject) {
	if _, ok := p.objects.Load(object.GetId()); !ok {
		p.objects.Store(object.GetId(), object)
	}
}

func (p *Pools) DestroyBlip(blip **IBlip) {
	if _, ok := p.blips.Load((*blip).GetId()); ok {
		p.blips.Delete((*blip).GetId())
		*blip = nil
	}
}

func (p *Pools) DestroyPlayer(player **IPlayer) {
	if _, ok := p.players.Load((*player).GetId()); ok {
		p.players.Delete((*player).GetId())
		*player = nil
	}
}

func (p *Pools) DestroyVehicle(vehicle **IVehicle) {
	if _, ok := p.vehicles.Load((*vehicle).GetId()); ok {
		p.vehicles.Delete((*vehicle).GetId())
	}
	*vehicle = nil
}

func (p *Pools) DestroyPed(ped **IPed) {
	if _, ok := p.peds.Load((*ped).GetId()); ok {
		p.peds.Delete((*ped).GetId())
	}
	*ped = nil
}

func (p *Pools) DestroyColshape(colshape **IColshape) {
	if _, ok := p.colshapes.Load((*colshape).GetId()); ok {
		p.colshapes.Delete((*colshape).GetId())
	}
	colshape = nil
}

func (p *Pools) DestroyCheckpoint(checkpoint **ICheckpoint) {
	if _, ok := p.checkpoints.Load((*checkpoint).GetId()); ok {
		p.checkpoints.Delete((*checkpoint).GetId())
	}
	*checkpoint = nil
}

func (p *Pools) DestroyMarker(marker **IMarker) {
	if _, ok := p.markers.Load((*marker).GetId()); ok {
		p.markers.Delete((*marker).GetId())
	}
	*marker = nil
}

func (p *Pools) DestroyObject(object **IObject) {
	if _, ok := p.objects.Load((*object).GetId()); ok {
		p.objects.Delete((*object).GetId())
	}
	*object = nil
}

func (p *Pools) GetPed(id uint32) *IPed {
	if value, ok := p.peds.Load(id); ok {
		return value.(*IPed)
	}
	return nil
}

func (p *Pools) GetBlip(id uint32) *IBlip {
	if value, ok := p.blips.Load(id); ok {
		return value.(*IBlip)
	}
	return nil
}

func (p *Pools) GetVehicle(id uint32) *IVehicle {
	if value, ok := p.vehicles.Load(id); ok {
		return value.(*IVehicle)
	}
	return nil
}

func (p *Pools) GetPlayer(id uint32) *IPlayer {
	if value, ok := p.players.Load(id); ok {
		return value.(*IPlayer)
	}
	return nil
}

func (p *Pools) GetColshape(id uint32) *IColshape {
	if value, ok := p.colshapes.Load(id); ok {
		return value.(*IColshape)
	}
	return nil
}

func (p *Pools) GetCheckpoint(id uint32) *ICheckpoint {
	if value, ok := p.checkpoints.Load(id); ok {
		return value.(*ICheckpoint)
	}
	return nil
}

func (p *Pools) GetMarker(id uint32) *IMarker {
	if value, ok := p.markers.Load(id); ok {
		return value.(*IMarker)
	}
	return nil
}

func (p *Pools) GetObject(id uint32) *IObject {
	if value, ok := p.objects.Load(id); ok {
		return value.(*IObject)
	}
	return nil
}

func (p *Pools) GetVehiclePools() *sync.Map {
	return p.vehicles
}

func (p *Pools) GetPedPools() *sync.Map {
	return p.peds
}

func (p *Pools) GetBlipPools() *sync.Map {
	return p.blips
}

func (p *Pools) GetPlayerPools() *sync.Map {
	return p.players
}

func (p *Pools) GetColshapePools() *sync.Map {
	return p.colshapes
}

func (p *Pools) GetCheckpointPools() *sync.Map {
	return p.checkpoints
}

func (p *Pools) GetMarkerPools() *sync.Map {
	return p.markers
}

func (p *Pools) GetObjectPools() *sync.Map {
	return p.objects
}

func GetPools() *Pools {
	return pools
}
