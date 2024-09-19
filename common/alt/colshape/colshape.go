package colshape

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_type"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/19
   File: colshape.go
*/

func CreateColshapeCircle(position *entities.Vector3, radius float32) *models.IColshape {
	var w = &lib.Warrper{}
	var c = &models.IColshape{}
	posData, posMetaData := uint64(math.Float32bits(position.X))|(uint64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	ret, freePtrFunc := w.CreateColshape(colshape_type.Circle, posData, posMetaData, 0, 0, radius, 0)
	cColshape := entities.ConvertCColshape(ret)
	if cColshape != nil {
		freePtrFunc()
		c = c.NewIColshape(cColshape.ID, cColshape.ColshapeType, cColshape.Position)
		pools := models.GetPools()
		pools.PutColshape(c)
		return c
	}
	return nil
}

func CreateColshapeSphere(position *entities.Vector3, radius float32) *models.IColshape {
	var w = &lib.Warrper{}
	var c = &models.IColshape{}
	posData, posMetaData := uint64(math.Float32bits(position.X))|(uint64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	ret, freePtrFunc := w.CreateColshape(colshape_type.Sphere, posData, posMetaData, 0, 0, radius, 0)
	cColshape := entities.ConvertCColshape(ret)
	if cColshape != nil {
		freePtrFunc()
		c = c.NewIColshape(cColshape.ID, cColshape.ColshapeType, cColshape.Position)
		pools := models.GetPools()
		pools.PutColshape(c)
		return c
	}
	return nil
}

func CreateColshapeRectangle(pointPosition *entities.Vector3, secondPointPosition *entities.Vector3) *models.IColshape {
	var w = &lib.Warrper{}
	var c = &models.IColshape{}
	posData, posMetaData := uint64(math.Float32bits(pointPosition.X))|(uint64(math.Float32bits(pointPosition.Y))<<32), uint64(math.Float32bits(pointPosition.Z))<<32
	secondPosData, secondPosMetaData := uint64(math.Float32bits(secondPointPosition.X))|(uint64(math.Float32bits(secondPointPosition.Y))<<32), uint64(math.Float32bits(secondPointPosition.Z))<<32
	ret, freePtrFunc := w.CreateColshape(colshape_type.Rectangle, posData, posMetaData, secondPosData, secondPosMetaData, 0, 0)
	cColshape := entities.ConvertCColshape(ret)
	if cColshape != nil {
		freePtrFunc()
		c = c.NewIColshape(cColshape.ID, cColshape.ColshapeType, cColshape.Position)
		pools := models.GetPools()
		pools.PutColshape(c)
		return c
	}
	return nil
}

func CreateColshapeCuboid(pointPosition *entities.Vector3, secondPointPosition *entities.Vector3) *models.IColshape {
	var w = &lib.Warrper{}
	var c = &models.IColshape{}
	posData, posMetaData := uint64(math.Float32bits(pointPosition.X))|(uint64(math.Float32bits(pointPosition.Y))<<32), uint64(math.Float32bits(pointPosition.Z))<<32
	secondPosData, secondPosMetaData := uint64(math.Float32bits(secondPointPosition.X))|(uint64(math.Float32bits(secondPointPosition.Y))<<32), uint64(math.Float32bits(secondPointPosition.Z))<<32
	ret, freePtrFunc := w.CreateColshape(colshape_type.Cuboid, posData, posMetaData, secondPosData, secondPosMetaData, 0, 0)
	cColshape := entities.ConvertCColshape(ret)
	if cColshape != nil {
		freePtrFunc()
		c = c.NewIColshape(cColshape.ID, cColshape.ColshapeType, cColshape.Position)
		pools := models.GetPools()
		pools.PutColshape(c)
		return c
	}
	return nil
}

func CreateColshapeCylinder(pointPosition *entities.Vector3, radius, height float32) *models.IColshape {
	var w = &lib.Warrper{}
	var c = &models.IColshape{}
	posData, posMetaData := uint64(math.Float32bits(pointPosition.X))|(uint64(math.Float32bits(pointPosition.Y))<<32), uint64(math.Float32bits(pointPosition.Z))<<32
	ret, freePtrFunc := w.CreateColshape(colshape_type.Cylinder, posData, posMetaData, 0, 0, radius, height)
	cColshape := entities.ConvertCColshape(ret)
	if cColshape != nil {
		freePtrFunc()
		c = c.NewIColshape(cColshape.ID, cColshape.ColshapeType, cColshape.Position)
		pools := models.GetPools()
		pools.PutColshape(c)
		return c
	}
	return nil
}
