package models

import (
	"github.com/StanZzzz222/RAltGo/hash_enums"
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

type IObject struct {
	id                uint32
	model             uint32
	alpha             uint8
	textureVariation  uint8
	streamingDistance uint32
	streamed          bool
	datas             *sync.Map
	*BaseObject
}

func (o *IObject) GetId() uint32                { return o.id }
func (o *IObject) GetDimension() int32          { return o.dimension }
func (o *IObject) GetVisible() bool             { return o.visible }
func (o *IObject) GetFrozen() bool              { return o.frozen }
func (o *IObject) GetCollision() bool           { return o.collision }
func (o *IObject) GetModel() uint32             { return o.model }
func (o *IObject) GetAlpha() uint8              { return o.alpha }
func (o *IObject) GetTextureVariation() uint8   { return o.textureVariation }
func (o *IObject) GetStreamingDistance() uint32 { return o.streamingDistance }
func (o *IObject) GetStreamed() bool            { return o.streamed }
func (o *IObject) GetLodDistance() uint16 {
	ret, freeDataResultFunc := w.GetData(o.id, enum.Object, uint8(enum.ObjectLodDistance))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U16Val
	}
	return 0
}
func (o *IObject) GetPosition() *entities.Vector3 {
	ret, freeDataResultFunc := w.GetData(o.id, enum.Object, uint8(enum.ObjectPosition))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.Vector3Val
	}
	return nil
}
func (o *IObject) GetRotation() *entities.Vector3 {
	ret, freeDataResultFunc := w.GetData(o.id, enum.Object, uint8(enum.ObjectRotation))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.Vector3Val
	}
	return nil
}

func (o *IObject) NewIObject(id, model uint32, position, rotation *entities.Vector3) *IObject {
	return &IObject{
		id:                id,
		model:             model,
		alpha:             255,
		textureVariation:  0,
		streamingDistance: 0,
		streamed:          false,
		datas:             &sync.Map{},
		BaseObject:        NewBaseObject(position, rotation, hash_enums.DefaultDimension, false, true, true),
	}
}

func (o *IObject) SetDimension(dimension int32) {
	o.dimension = dimension
	w.SetObjectData(o.id, enum.ObjectDimension, int64(dimension), 0)
}

func (o *IObject) SetVisible(visible bool) {
	o.visible = visible
	value := 0
	if visible {
		value = 1
	}
	w.SetObjectData(o.id, enum.ObjectVisible, int64(value), 0)
}

func (o *IObject) SetAlpha(alpha uint8) {
	o.alpha = alpha
	w.SetObjectData(o.id, enum.ObjectAlpha, int64(alpha), 0)
}

func (o *IObject) SetLodDistance(lodDistance uint16) {
	w.SetObjectData(o.id, enum.ObjectLodDistance, int64(lodDistance), 0)
}

func (o *IObject) SetTextureVariation(textureVariation uint8) {
	o.textureVariation = textureVariation
	w.SetObjectData(o.id, enum.ObjectTextureVariation, int64(textureVariation), 0)
}

func (o *IObject) SetStreamingDistance(streamingDistance uint32) {
	o.streamingDistance = streamingDistance
	w.SetObjectData(o.id, enum.ObjectStreamingDistance, int64(streamingDistance), 0)
}

func (o *IObject) SetStreamed(streamed bool) {
	o.streamed = streamed
	value := 0
	if streamed {
		value = 1
	}
	w.SetObjectData(o.id, enum.ObjectStreamed, int64(value), 0)
}

func (o *IObject) SetFrozen(frozen bool) {
	o.frozen = frozen
	value := 0
	if frozen {
		value = 1
	}
	w.SetObjectData(o.id, enum.ObjectFrozen, int64(value), 0)
}

func (o *IObject) SetCollision(collision bool) {
	o.collision = collision
	value := 0
	if collision {
		value = 1
	}
	w.SetObjectData(o.id, enum.ObjectCollision, int64(value), 0)
}

func (o *IObject) SetPosition(position *entities.Vector3) {
	o.position = position
	posData, posMetaData := int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	w.SetObjectData(o.id, enum.ObjectPosition, posData, posMetaData)
}

func (o *IObject) SetRotation(rotation *entities.Vector3) {
	o.rotation = rotation
	rotData, rotMetaData := int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32
	w.SetObjectData(o.id, enum.ObjectRotation, rotData, rotMetaData)
}

func (o *IObject) Destroy() {
	w.SetObjectData(o.id, enum.ObjectDestory, int64(0), 0)
	pools.DestroyObject(&o)
}

func (o *IObject) SetData(key string, value any) {
	o.datas.Store(key, value)
}

func (o *IObject) DelData(key string) {
	_, ok := o.datas.Load(key)
	if ok {
		o.datas.Delete(key)
	}
}

func (o *IObject) DelAllData() {
	o.datas.Range(func(key, value any) bool {
		o.datas.Delete(key)
		return true
	})
}

func (o *IObject) HasData(key string) bool {
	_, ok := o.datas.Load(key)
	if ok {
		return true
	}
	return false
}

func (o *IObject) GetData(key string) any {
	value, ok := o.datas.Load(key)
	if ok {
		return value
	}
	return value
}

func (o *IObject) GetDatas() []any {
	var datas []any
	o.datas.Range(func(key, value any) bool {
		datas = append(datas, key)
		return true
	})
	return datas
}
