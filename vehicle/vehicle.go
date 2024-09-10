package vehicle

import "C"
import (
	"github.com/StanZzzz222/RAltGo/enums/vehicle"
	"github.com/StanZzzz222/RAltGo/internal/entitys"
	"github.com/StanZzzz222/RAltGo/models"
	"github.com/StanZzzz222/RAltGo/utils"
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
		cVeh := (*entitys.CVehicle)(unsafe.Pointer(ret))
		return veh.NewIVehicle(cVeh.ID, cVeh.Model, cVeh.PrimaryColor, cVeh.SecondColor, cVeh.Position, cVeh.Rotation)
	}
}

func CreateVehicleByHash(model vehicle.ModelHash, numberPlate string, position, rotation *models.Vector3, primaryColor, secondColor uint8) *models.IVehicle {
	var veh = &models.IVehicle{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	rotData := uint64(math.Float32bits(rotation.X)) | (uint64(math.Float32bits(rotation.Y)) << 32)
	rotMetadata := uint64(math.Float32bits(rotation.Z)) << 32
	ret := w.CreateVehicle(uint32(model), posData, posMetadata, rotData, rotMetadata, w.GoStringMarshalPtr(numberPlate), primaryColor, secondColor)
	switch ret {
	case 0:
		return nil
	default:
		cVeh := (*entitys.CVehicle)(unsafe.Pointer(ret))
		return veh.NewIVehicle(cVeh.ID, cVeh.Model, cVeh.PrimaryColor, cVeh.SecondColor, cVeh.Position, cVeh.Rotation)
	}
}
