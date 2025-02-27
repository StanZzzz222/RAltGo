package alt_vehicle

import "C"
import (
	"github.com/StanZzzz222/RAltGo/common"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_hash"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/9
   File: vehicle_hash.go
*/

func CreateVehicle(model string, numberPlate string, position, rotation *models.Vector3, primaryColor, secondColor uint8) *models.IVehicle {
	var w = lib.GetWrapper()
	var veh = &models.IVehicle{}
	posData, posMetaData := uint64(math.Float32bits(position.X))|(uint64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	rotData, rotMetaData := uint64(math.Float32bits(rotation.X))|(uint64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32
	numplate, freeCStringFunc := w.GoStringMarshalPtr(numberPlate)
	ret, freePtrFunc := w.CreateVehicle(common.Hash(model), posData, posMetaData, rotData, rotMetaData, numplate, primaryColor, secondColor)
	cVeh := entities.ConvertCVehicle(ret)
	if cVeh != nil {
		freePtrFunc()
		freeCStringFunc()
		veh = veh.NewIVehicle(cVeh.ID, cVeh.Model, numberPlate, cVeh.PrimaryColor, cVeh.SecondColor, (*models.Vector3)(cVeh.Position), (*models.Vector3)(cVeh.Rotation))
		pools := models.GetPools()
		pools.PutVehicle(veh)
		return veh
	}
	return nil
}

func CreateVehicleByHash(model vehicle_hash.ModelHash, numberPlate string, position, rotation *models.Vector3, primaryColor, secondColor uint8) *models.IVehicle {
	var w = lib.GetWrapper()
	var veh = &models.IVehicle{}
	posData, posMetaData := uint64(math.Float32bits(position.X))|(uint64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	rotData, rotMetaData := uint64(math.Float32bits(rotation.X))|(uint64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32
	numplate, freeCStringFunc := w.GoStringMarshalPtr(numberPlate)
	ret, freePtrFunc := w.CreateVehicle(uint32(model), posData, posMetaData, rotData, rotMetaData, numplate, primaryColor, secondColor)
	cVeh := entities.ConvertCVehicle(ret)
	if cVeh != nil {
		freePtrFunc()
		freeCStringFunc()
		veh = veh.NewIVehicle(cVeh.ID, cVeh.Model, numberPlate, cVeh.PrimaryColor, cVeh.SecondColor, (*models.Vector3)(cVeh.Position), (*models.Vector3)(cVeh.Rotation))
		pools := models.GetPools()
		pools.PutVehicle(veh)
		return veh
	}
	return nil
}
