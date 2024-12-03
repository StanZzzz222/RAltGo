package models

import (
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_type"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enums"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
	"reflect"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/19
   File: colshape.go
*/

type IColshape struct {
	id           uint32
	colshapeType colshape_type.ColshapeType
	position     *Vector3
	playersOnly  bool
	dimension    int32
	datas        *sync.Map
	warpper      *lib.Warpper
	*NetworkData
}

func (c *IColshape) GetId() uint32                               { return c.id }
func (c *IColshape) GetColshapeType() colshape_type.ColshapeType { return c.colshapeType }
func (c *IColshape) GetPosition() *Vector3                       { return c.position }
func (c *IColshape) GetPlayersOnly() bool                        { return c.playersOnly }
func (c *IColshape) GetDimension() int32                         { return c.dimension }
func (c *IColshape) IsEntityIdIn(syncId SyncId) bool {
	ret, freeDataResultFunc := c.warpper.GetColshapeData(c.id, enums.ColshapeIsEntityIdIn, 0, 0, uint64(syncId))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (c *IColshape) IsPointIn(position *Vector3) bool {
	ret, freeDataResultFunc := c.warpper.GetColshapeData(c.id, enums.ColshapeIsPointIn, 0, int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (c *IColshape) IsEntityIn(entity any) bool {
	if res, entityType, id := checkEntity(entity); res {
		ret, freeDataResultFunc := c.warpper.GetColshapeData(c.id, enums.ColshapeIsEntityIn, entityType, int64(id), 0)
		cDataResult := entities.ConverCDataResult(ret)
		if cDataResult != nil {
			freeDataResultFunc()
			return cDataResult.BoolVal
		}
	}
	return false
}

func (c *IColshape) NewIColshape(id uint32, colshapeType uint32, position *Vector3) *IColshape {
	return &IColshape{
		id:           id,
		colshapeType: colshape_type.ColshapeType(colshapeType),
		position:     position,
		dimension:    hash_enums.DefaultDimension,
		datas:        &sync.Map{},
		warpper:      lib.GetWrapper(),
		NetworkData:  NewNetworkData(id, enums.Colshape),
	}
}

func (c *IColshape) SetDimension(dimension int32) {
	c.dimension = dimension
	c.warpper.SetColshapeData(c.id, enums.ColshapeDimension, int64(dimension), 0)
}

func (c *IColshape) SetPlayersOnly(playersOnly bool) {
	c.playersOnly = playersOnly
	value := 0
	if playersOnly {
		value = 1
	}
	c.warpper.SetColshapeData(c.id, enums.ColshapePlayersOnly, int64(value), 0)
}

func (c *IColshape) SetPosition(position *Vector3) {
	c.position = position
	posData, posMetaData := int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	c.warpper.SetColshapeData(c.id, enums.ColshapePosition, posData, posMetaData)
}

func (c *IColshape) Destroy() {
	c.warpper.SetColshapeData(c.id, enums.ColshapeDestory, 0, 0)
	pools.DestroyColshape(c)
}

func (c *IColshape) SetData(key string, value any) {
	c.datas.Store(key, value)
}

func (c *IColshape) DelData(key string) {
	_, ok := c.datas.Load(key)
	if ok {
		c.datas.Delete(key)
	}
}

func (c *IColshape) DelAllData() {
	c.datas.Range(func(key, value any) bool {
		c.datas.Delete(key)
		return true
	})
}

func (c *IColshape) HasData(key string) bool {
	_, ok := c.datas.Load(key)
	if ok {
		return true
	}
	return false
}

func (c *IColshape) GetData(key string) any {
	value, ok := c.datas.Load(key)
	if ok {
		return value
	}
	return value
}

func (c *IColshape) GetDatas() []any {
	var datas []any
	c.datas.Range(func(key, value any) bool {
		datas = append(datas, key)
		return true
	})
	return datas
}

func checkEntity(targetEntity any) (bool, enums.ObjectType, uint32) {
	var res = false
	var entityType = enums.ObjectType(0)
	var id uint32 = 0
	t := reflect.TypeOf(targetEntity)
	if t.Kind() == reflect.Ptr {
		elemType := t.Elem()
		switch elemType {
		case reflect.TypeOf((*IPlayer)(nil)).Elem():
			res = true
			entityType = enums.Player
			id = targetEntity.(*IPlayer).GetId()
			break
		case reflect.TypeOf((*IVehicle)(nil)).Elem():
			res = true
			entityType = enums.Vehicle
			id = targetEntity.(*IVehicle).GetId()
			break
		case reflect.TypeOf((*IBlip)(nil)).Elem():
			res = true
			entityType = enums.Ped
			id = targetEntity.(*IBlip).GetId()
			break
		case reflect.TypeOf((*IPed)(nil)).Elem():
			res = true
			entityType = enums.Ped
			id = targetEntity.(*IPed).GetId()
			break
		case reflect.TypeOf((*IObject)(nil)).Elem():
			res = true
			entityType = enums.Object
			id = targetEntity.(*IObject).GetId()
			break
		}
	}
	return res, entityType, id
}
