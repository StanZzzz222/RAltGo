package models

import (
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"github.com/StanZzzz222/RAltGo/hash_enums/check_point_type"
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

type ICheckpoint struct {
	id             uint32
	checkpointType check_point_type.CheckPointType
	position       *entities.Vector3
	dimension      int32
	visible        bool
	playersOnly    bool
	height         float32
	radius         float32
	nextPosition   *entities.Vector3
	color          *entities.Rgba
	iconColor      *entities.Rgba
	datas          *sync.Map
}

func (c *ICheckpoint) GetId() uint32                                      { return c.id }
func (c *ICheckpoint) GetCheckpointType() check_point_type.CheckPointType { return c.checkpointType }
func (c *ICheckpoint) GetPosition() *entities.Vector3                     { return c.position }
func (c *ICheckpoint) GetDimension() int32                                { return c.dimension }
func (c *ICheckpoint) GetVisible() bool                                   { return c.visible }
func (c *ICheckpoint) GetPlayersOnly() bool                               { return c.playersOnly }
func (c *ICheckpoint) GetHeight() float32                                 { return c.height }
func (c *ICheckpoint) GetRadius() float32                                 { return c.radius }
func (c *ICheckpoint) GetNextPosition() *entities.Vector3                 { return c.nextPosition }
func (c *ICheckpoint) GetColor() *entities.Rgba                           { return c.color }
func (c *ICheckpoint) GetIconColor() *entities.Rgba                       { return c.iconColor }

func (c *ICheckpoint) NewICheckPoint(id uint32, checkPointType uint8, position *entities.Vector3, height, radius float32) *ICheckpoint {
	return &ICheckpoint{
		id:             id,
		checkpointType: check_point_type.CheckPointType(checkPointType),
		position:       position,
		dimension:      hash_enums.DefaultDimension,
		visible:        true,
		playersOnly:    false,
		height:         height,
		radius:         radius,
		nextPosition:   nil,
		color:          nil,
		iconColor:      nil,
		datas:          &sync.Map{},
	}
}

func (c *ICheckpoint) SetDimension(dimension int32) {
	c.dimension = dimension
	w.SetCheckpointData(c.id, enum.CheckpointDimension, int64(dimension), 0, 0, 0, 0, 0, 0)
}

func (c *ICheckpoint) SetVisible(visible bool) {
	c.visible = visible
	value := 0
	if visible {
		value = 1
	}
	w.SetCheckpointData(c.id, enum.CheckpointVisible, int64(value), 0, 0, 0, 0, 0, 0)
}

func (c *ICheckpoint) SetPosition(position *entities.Vector3) {
	c.position = position
	posData, posMetaData := int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	w.SetCheckpointData(c.id, enum.CheckpointPosition, posData, posMetaData, 0, 0, 0, 0, 0)
}

func (c *ICheckpoint) SetCheckpointType(checkpointType check_point_type.CheckPointType) {
	c.checkpointType = checkpointType
	w.SetCheckpointData(c.id, enum.CheckpointType, int64(checkpointType), 0, 0, 0, 0, 0, 0)
}

func (c *ICheckpoint) SetHeight(height float32) {
	c.height = height
	w.SetCheckpointData(c.id, enum.CheckpointHeight, 0, 0, height, 0, 0, 0, 0)
}

func (c *ICheckpoint) SetRadius(radius float32) {
	c.radius = radius
	w.SetCheckpointData(c.id, enum.CheckpointRadius, 0, 0, radius, 0, 0, 0, 0)
}

func (c *ICheckpoint) SetPlayersOnly(playersOnly bool) {
	c.playersOnly = playersOnly
	value := 0
	if playersOnly {
		value = 1
	}
	w.SetCheckpointData(c.id, enum.CheckpointPlayersOnly, int64(value), 0, 0, 0, 0, 0, 0)
}

func (c *ICheckpoint) SetNextPosition(position *entities.Vector3) {
	c.nextPosition = position
	posData, posMetaData := int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	w.SetCheckpointData(c.id, enum.CheckpointNextPosition, posData, posMetaData, 0, 0, 0, 0, 0)
}

func (c *ICheckpoint) SetColor(color *entities.Rgba) {
	c.color = color
	w.SetCheckpointData(c.id, enum.CheckpointColor, 0, 0, 0, color.R, color.G, color.B, color.A)
}

func (c *ICheckpoint) SetIconColor(iconColor *entities.Rgba) {
	c.iconColor = iconColor
	w.SetCheckpointData(c.id, enum.CheckpointIconColor, 0, 0, 0, iconColor.R, iconColor.G, iconColor.B, iconColor.A)
}

func (c *ICheckpoint) Destroy() {
	w.SetCheckpointData(c.id, enum.CheckpointDestory, 0, 0, 0, 0, 0, 0, 0)
	pools.DestroyCheckpoint(c)
}

func (c *ICheckpoint) SetData(key string, value any) {
	c.datas.Store(key, value)
}

func (c *ICheckpoint) DelData(key string) {
	_, ok := c.datas.Load(key)
	if ok {
		c.datas.Delete(key)
	}
}

func (c *ICheckpoint) DelAllData() {
	c.datas.Range(func(key, value any) bool {
		c.datas.Delete(key)
		return true
	})
}

func (c *ICheckpoint) HasData(key string) bool {
	_, ok := c.datas.Load(key)
	if ok {
		return true
	}
	return false
}

func (c *ICheckpoint) GetData(key string) any {
	value, ok := c.datas.Load(key)
	if ok {
		return value
	}
	return value
}

func (c *ICheckpoint) GetDatas() []any {
	var datas []any
	c.datas.Range(func(key, value any) bool {
		datas = append(datas, key)
		return true
	})
	return datas
}