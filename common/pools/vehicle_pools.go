package pools

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/18
   File: vehicle_pools.go
*/

var vehiclePools = newVehiclePool()

type VehiclePool struct {
	vehicles *sync.Map
}

func newVehiclePool() *VehiclePool {
	return &VehiclePool{
		vehicles: &sync.Map{},
	}
}

func (p *VehiclePool) Put(vehicle *models.IVehicle) {
	if _, ok := p.vehicles.Load(vehicle.GetId()); !ok {
		p.vehicles.Store(vehicle.GetId(), vehicle)
	}
}

func (p *VehiclePool) Remove(vehicle *models.IVehicle) {
	if _, ok := p.vehicles.Load(vehicle.GetId()); ok {
		p.vehicles.Delete(vehicle.GetId())
	}
}

func GetVehicle(id uint32) *models.IVehicle {
	vehicle, ok := vehiclePools.vehicles.Load(id)
	if ok {
		return vehicle.(*models.IVehicle)
	}
	return nil
}

func GetVehicles() []*models.IVehicle {
	var vehicles []*models.IVehicle
	vehiclePools.vehicles.Range(func(key, value any) bool {
		vehicles = append(vehicles, value.(*models.IVehicle))
		return true
	})
	return vehicles
}

func GetVehicleIterator() <-chan *models.IVehicle {
	vehicleChan := make(chan *models.IVehicle)
	go func() {
		defer close(vehicleChan)
		playerPools.players.Range(func(key, value any) bool {
			vehicleChan <- value.(*models.IVehicle)
			return true
		})
	}()
	return vehicleChan
}

func GetVehiclePools() *VehiclePool {
	return vehiclePools
}
