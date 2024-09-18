package pools

import "github.com/StanZzzz222/RAltGo/common/models"

/*
   Create by zyx
   Date Time: 2024/9/18
   File: vehicle.go
*/

func GetVehicle(id uint32) *models.IVehicle {
	return models.GetPools().GetVehicle(id)
}

func GetVehicles() []*models.IVehicle {
	var vehicles []*models.IVehicle
	var pools = models.GetPools().GetVehiclePools()
	pools.Range(func(key, value any) bool {
		vehicles = append(vehicles, value.(*models.IVehicle))
		return true
	})
	return vehicles
}

func GetVehicleIterator() <-chan *models.IVehicle {
	var vehicleChan = make(chan *models.IVehicle)
	var pools = models.GetPools().GetVehiclePools()
	go func() {
		defer close(vehicleChan)
		pools.Range(func(key, value any) bool {
			vehicleChan <- value.(*models.IVehicle)
			return true
		})
	}()
	return vehicleChan
}
