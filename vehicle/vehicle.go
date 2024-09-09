package vehicle

import "C"
import (
	"fmt"
	"gamemode/enums/vehicle"
	"gamemode/models"
	"gamemode/utils"
	"math"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/9
   File: vehicle.go
*/

func CreateVehicle(model string, numberPlate string, position, rotation *models.Vector3, primaryColor, secondColor uint8) *models.IVehicle {
	var veh = &models.IVehicle{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	rotData := uint64(math.Float32bits(rotation.X)) | (uint64(math.Float32bits(rotation.Y)) << 32)
	rotMetadata := uint64(math.Float32bits(rotation.Z)) << 32
	ret := w.CreateVehicle(utils.Hash(model), posData, posMetadata, rotData, rotMetadata, w.GoStringMarshalPtr(numberPlate), primaryColor, secondColor)
	switch ret {
	case 0:
		return nil
	default:
		cVeh := (*C.Vehicle)(unsafe.Pointer(ret))
		pos := (*models.Vector3)(unsafe.Pointer(player.position))
		rot := (*models.Vector3)(unsafe.Pointer(player.rotation))
		return veh.NewIVehicle(cVeh.id, cVeh.model, cVeh.position)
	}
}

func CreateVehicleByHash(model vehicle.ModelHash, numberPlate string, position, rotation *models.Vector3, primaryColor, secondColor uint8) *models.IVehicle {
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	rotData := uint64(math.Float32bits(rotation.X)) | (uint64(math.Float32bits(rotation.Y)) << 32)
	rotMetadata := uint64(math.Float32bits(rotation.Z)) << 32
	ret := w.CreateVehicle(uint32(model), posData, posMetadata, rotData, rotMetadata, w.GoStringMarshalPtr(numberPlate), primaryColor, secondColor)
	fmt.Println(ret)
	switch ret {
	case 0:
		return nil
	default:

	}
	return nil
}
