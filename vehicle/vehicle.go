package vehicle

import "C"
import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/common/utils"
	"github.com/StanZzzz222/RAltGo/enums/vehicle"
	"github.com/StanZzzz222/RAltGo/internal/entitys"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/9
   File: vehicle.go
*/

var rw = &sync.RWMutex{}
var w = &lib.Warrper{}

func CreateVehicle(model string, numberPlate string, position, rotation *entitys.Vector3, primaryColor, secondColor uint8) *models.IVehicle {
	rw.Lock()
	defer rw.Unlock()
	var veh = &models.IVehicle{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	rotData := uint64(math.Float32bits(rotation.X)) | (uint64(math.Float32bits(rotation.Y)) << 32)
	rotMetadata := uint64(math.Float32bits(rotation.Z)) << 32
	ret := w.CreateVehicle(utils.Hash(model), posData, posMetadata, rotData, rotMetadata, w.GoStringMarshalPtr(numberPlate), primaryColor, secondColor)
	cVeh := entitys.ConvertCVehicle(ret)
	if cVeh != nil {
		veh = veh.NewIVehicle(cVeh.ID, cVeh.Model, cVeh.PrimaryColor, cVeh.SecondColor, cVeh.Position, cVeh.Rotation)
		return veh
	}
	return nil
}

func CreateVehicleByHash(model vehicle.ModelHash, numberPlate string, position, rotation *entitys.Vector3, primaryColor, secondColor uint8) *models.IVehicle {
	rw.Lock()
	defer rw.Unlock()
	var veh = &models.IVehicle{}
	posData := uint64(math.Float32bits(position.X)) | (uint64(math.Float32bits(position.Y)) << 32)
	posMetadata := uint64(math.Float32bits(position.Z)) << 32
	rotData := uint64(math.Float32bits(rotation.X)) | (uint64(math.Float32bits(rotation.Y)) << 32)
	rotMetadata := uint64(math.Float32bits(rotation.Z)) << 32
	ret := w.CreateVehicle(uint32(model), posData, posMetadata, rotData, rotMetadata, w.GoStringMarshalPtr(numberPlate), primaryColor, secondColor)
	cVeh := entitys.ConvertCVehicle(ret)
	if cVeh != nil {
		veh = veh.NewIVehicle(cVeh.ID, cVeh.Model, cVeh.PrimaryColor, cVeh.SecondColor, cVeh.Position, cVeh.Rotation)
		return veh
	}
	return nil
}
