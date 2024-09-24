package object

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/common/utils"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/25
   File: object.go
*/

func CreateObject(model uint32, position, rotation *entities.Vector3) *models.IObject {
	var w = lib.GetWarpper()
	var o = &models.IObject{}
	posData, posMetaData := uint64(math.Float32bits(position.X))|(uint64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	rotData, rotMetaData := uint64(math.Float32bits(rotation.X))|(uint64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32
	ret, freePtrFunc := w.CreateObject(model, posData, posMetaData, rotData, rotMetaData)
	cObject := entities.ConvertCObject(ret)
	if cObject != nil {
		freePtrFunc()
		o = o.NewIObject(cObject.ID, model, position, rotation)
		pools := models.GetPools()
		pools.PutObject(o)
		return o
	}
	return nil
}

func CreateObjectByModelName(model string, position, rotation *entities.Vector3) *models.IObject {
	var modelHash = utils.Hash(model)
	var w = lib.GetWarpper()
	var o = &models.IObject{}
	posData, posMetaData := uint64(math.Float32bits(position.X))|(uint64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	rotData, rotMetaData := uint64(math.Float32bits(rotation.X))|(uint64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32
	ret, freePtrFunc := w.CreateObject(modelHash, posData, posMetaData, rotData, rotMetaData)
	cObject := entities.ConvertCObject(ret)
	if cObject != nil {
		freePtrFunc()
		o = o.NewIObject(cObject.ID, modelHash, position, rotation)
		pools := models.GetPools()
		pools.PutObject(o)
		return o
	}
	return nil
}
