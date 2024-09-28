package alt_ped

import (
	"github.com/StanZzzz222/RAltGo/common"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/ped_hash"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/13
   File: ped_hash.go
*/

func CreatePed(model string, position, rotation *entities.Vector3) *models.IPed {
	var w = lib.GetWarpper()
	var p = &models.IPed{}
	var posData, posMetaData = int64(math.Float32bits(position.X)) | (int64(math.Float32bits(position.Y)) << 32), uint64(math.Float32bits(position.Z)) << 32
	var rotData, rotMetaData = int64(math.Float32bits(rotation.X)) | (int64(math.Float32bits(rotation.Y)) << 32), uint64(math.Float32bits(rotation.Z)) << 32
	ret, freePtrFunc := w.CreatePed(common.Hash(model), uint64(posData), posMetaData, uint64(rotData), rotMetaData, 0)
	cPed := entities.ConvertCPed(ret)
	if cPed != nil {
		freePtrFunc()
		p = p.NewIPed(cPed.ID, cPed.Model, cPed.Position, cPed.Rotation)
		pools := models.GetPools()
		pools.PutPed(p)
		return p
	}
	return nil
}

func CreateStreamPed(model string, position, rotation *entities.Vector3, streamingDistance uint32) *models.IPed {
	var w = lib.GetWarpper()
	var p = &models.IPed{}
	var posData, posMetaData = int64(math.Float32bits(position.X)) | (int64(math.Float32bits(position.Y)) << 32), uint64(math.Float32bits(position.Z)) << 32
	var rotData, rotMetaData = int64(math.Float32bits(rotation.X)) | (int64(math.Float32bits(rotation.Y)) << 32), uint64(math.Float32bits(rotation.Z)) << 32
	ret, freePtrFunc := w.CreatePed(common.Hash(model), uint64(posData), posMetaData, uint64(rotData), rotMetaData, streamingDistance)
	cPed := entities.ConvertCPed(ret)
	if cPed != nil {
		freePtrFunc()
		p = p.NewIPed(cPed.ID, cPed.Model, cPed.Position, cPed.Rotation)
		pools := models.GetPools()
		pools.PutPed(p)
		return p
	}
	return nil
}

func CreatePedByHash(model ped_hash.ModelHash, position, rotation *entities.Vector3) *models.IPed {
	var w = lib.GetWarpper()
	var p = &models.IPed{}
	var posData, posMetaData = int64(math.Float32bits(position.X)) | (int64(math.Float32bits(position.Y)) << 32), uint64(math.Float32bits(position.Z)) << 32
	var rotData, rotMetaData = int64(math.Float32bits(rotation.X)) | (int64(math.Float32bits(rotation.Y)) << 32), uint64(math.Float32bits(rotation.Z)) << 32
	ret, freePtrFunc := w.CreatePed(uint32(model), uint64(posData), posMetaData, uint64(rotData), rotMetaData, 0)
	cPed := entities.ConvertCPed(ret)
	if cPed != nil {
		freePtrFunc()
		p = p.NewIPed(cPed.ID, cPed.Model, cPed.Position, cPed.Rotation)
		pools := models.GetPools()
		pools.PutPed(p)
		return p
	}
	return nil
}

func CreateStreamPedByHash(model ped_hash.ModelHash, position, rotation *entities.Vector3, streamingDistance uint32) *models.IPed {
	var w = lib.GetWarpper()
	var p = &models.IPed{}
	var posData, posMetaData = int64(math.Float32bits(position.X)) | (int64(math.Float32bits(position.Y)) << 32), uint64(math.Float32bits(position.Z)) << 32
	var rotData, rotMetaData = int64(math.Float32bits(rotation.X)) | (int64(math.Float32bits(rotation.Y)) << 32), uint64(math.Float32bits(rotation.Z)) << 32
	ret, freePtrFunc := w.CreatePed(uint32(model), uint64(posData), posMetaData, uint64(rotData), rotMetaData, streamingDistance)
	cPed := entities.ConvertCPed(ret)
	if cPed != nil {
		freePtrFunc()
		p = p.NewIPed(cPed.ID, cPed.Model, cPed.Position, cPed.Rotation)
		pools := models.GetPools()
		pools.PutPed(p)
		return p
	}
	return nil
}
