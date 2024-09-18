package pools

import "github.com/StanZzzz222/RAltGo/common/models"

/*
   Create by zyx
   Date Time: 2024/9/18
   File: vehicle_pools.go
*/

func GetVehicle(id uint32) *models.IVehicle {
	vehicle, ok := vehiclePools.vehicles.Load(id)
	if ok {
		return vehicle.(*IVehicle)
	}
	return nil
}

func GetVehicles() []*IVehicle {
	var vehicles []*IVehicle
	vehiclePools.vehicles.Range(func(key, value any) bool {
		vehicles = append(vehicles, value.(*IVehicle))
		return true
	})
	return vehicles
}

func GetVehicleIterator() <-chan *IVehicle {
	vehicleChan := make(chan *IVehicle)
	go func() {
		defer close(vehicleChan)
		playerPools.players.Range(func(key, value any) bool {
			vehicleChan <- value.(*IVehicle)
			return true
		})
	}()
	return vehicleChan
}

func GetVehiclePools() *VehiclePool {
	return vehiclePools
}
