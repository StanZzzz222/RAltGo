package models

import (
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_type"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"math"
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
	position     *entities.Vector3
	playersOnly  bool
	dimension    int32
	datas        *sync.Map
}

func (c *IColshape) GetId() uint32                               { return c.id }
func (c *IColshape) GetColshapeType() colshape_type.ColshapeType { return c.colshapeType }
func (c *IColshape) GetPosition() *entities.Vector3              { return c.position }
func (c *IColshape) GetPlayersOnly() bool                        { return c.playersOnly }
func (c *IColshape) GetDimension() int32                         { return c.dimension }

func (c *IColshape) NewIColshape(id uint32, colshapeType uint32, position *entities.Vector3) *IColshape {
	return &IColshape{
		id:           id,
		colshapeType: colshape_type.ColshapeType(colshapeType),
		position:     position,
		dimension:    hash_enums.DefaultDimension,
		datas:        &sync.Map{},
	}
}

func (c *IColshape) SetDimension(dimension int32) {
	c.dimension = dimension
	w.SetColshapeData(c.id, enum.ColshapeDimension, int64(dimension), 0)
}

func (c *IColshape) SetPlayersOnly(playersOnly bool) {
	c.playersOnly = playersOnly
	value := 0
	if playersOnly {
		value = 1
	}
	w.SetColshapeData(c.id, enum.PlayersOnly, int64(value), 0)
}

func (c *IColshape) SetPosition(position *entities.Vector3) {
	c.position = position
	posData, posMetaData := int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	w.SetColshapeData(c.id, enum.ColshapePosition, posData, posMetaData)
}

func (c *IColshape) Destroy() {
	w.SetColshapeData(c.id, enum.ColshapeDestory, 0, 0)
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
