package models

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/utils"
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"github.com/StanZzzz222/RAltGo/hash_enums/ped_hash"
	"github.com/StanZzzz222/RAltGo/hash_enums/weapon_hash"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"math"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/13
   File: ped_hash.go
*/

type IPed struct {
	id                uint32
	model             ped_hash.ModelHash
	health            uint16
	armour            uint16
	maxHealth         uint16
	currentWeapon     weapon_hash.ModelHash
	streamingDistance uint32
	streamed          bool
	datas             *sync.Map
	*BaseObject
}

func (p *IPed) NewIPed(id uint32, model uint32, position, rotation *entities.Vector3, streamingDistance uint32, streamed bool) *IPed {
	return &IPed{
		id:                id,
		model:             ped_hash.ModelHash(model),
		currentWeapon:     weapon_hash.Fist,
		streamingDistance: streamingDistance,
		armour:            0,
		maxHealth:         200,
		streamed:          streamed,
		datas:             &sync.Map{},
		BaseObject:        NewBaseObject(position, rotation, hash_enums.DefaultDimension, false, true, true),
	}
}

func (p *IPed) GetId() uint32                           { return p.id }
func (p *IPed) GetModel() ped_hash.ModelHash            { return p.model }
func (p *IPed) GetCurrentWeapon() weapon_hash.ModelHash { return p.currentWeapon }
func (p *IPed) GetStreamingDistance() uint32            { return p.streamingDistance }
func (p *IPed) GetStreamed() bool                       { return p.streamed }
func (p *IPed) GetMaxHealth() uint16                    { return p.maxHealth }
func (p *IPed) GetDimension() int32                     { return p.dimension }
func (p *IPed) GetFrozen() bool                         { return p.frozen }
func (p *IPed) GetCollision() bool                      { return p.collision }
func (p *IPed) GetVisible() bool                        { return p.visible }
func (p *IPed) GetPosition() *entities.Vector3 {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Ped, uint8(enum.PedPosition))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.Vector3Val
	}
	return nil
}
func (p *IPed) GetRotation() *entities.Vector3 {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Ped, uint8(enum.PedRotation))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.Vector3Val
	}
	return nil
}
func (p *IPed) GetPositionString() string {
	position := p.GetPosition()
	return fmt.Sprintf("%v,%v,%v", position.X, position.Y, position.Z)
}
func (p *IPed) GetRotationString() string {
	rotation := p.GetRotation()
	return fmt.Sprintf("%v,%v,%v", rotation.X, rotation.Y, rotation.Z)
}
func (p *IPed) GetPositionRotation() (*entities.Vector3, *entities.Vector3) {
	return p.GetPosition(), p.GetRotation()
}
func (p *IPed) GetPositionRotationString() (string, string) {
	return p.GetPositionString(), p.GetRotationString()
}

func (p *IPed) SetPosition(position *entities.Vector3) {
	p.position = position
	w.SetPedMetaData(p.id, enum.PedPosition, int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32)
}

func (p *IPed) SetRotation(rotation *entities.Vector3) {
	p.rotation = rotation
	w.SetPedMetaData(p.id, enum.PedRotation, int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32)
}

func (p *IPed) SetHealth(health uint16) {
	if health > hash_enums.MaxHealth {
		p.health = health
		w.SetPedData(p.id, enum.PedMaxHealth, int64(health))
		return
	}
	p.health = health
	w.SetPedData(p.id, enum.PedHealth, int64(health))
}

func (p *IPed) SetArmour(armour uint16) {
	p.armour = armour
	w.SetPedData(p.id, enum.PedArmour, int64(armour))
}

func (p *IPed) SetDimension(dimension int32) {
	p.dimension = dimension
	w.SetPedData(p.id, enum.PedDimension, int64(dimension))
}

func (p *IPed) SetCollision(collision bool) {
	p.collision = collision
	value := 0
	if collision {
		value = 1
	}
	w.SetPedData(p.id, enum.PedCollision, int64(value))
}

func (p *IPed) SetFrozen(frozen bool) {
	p.frozen = frozen
	value := 0
	if frozen {
		value = 1
	}
	w.SetPedData(p.id, enum.PedFrozen, int64(value))
}

func (p *IPed) SetVisible(visible bool) {
	p.visible = visible
	value := 0
	if visible {
		value = 1
	}
	w.SetPedData(p.id, enum.PedVisible, int64(value))
}

func (p *IPed) SetStreamed(streamed bool) {
	p.streamed = streamed
	value := 0
	if streamed {
		value = 1
	}
	w.SetPedData(p.id, enum.PedStreamed, int64(value))
}

func (p *IPed) SetStreamingDistance(streamingDistance uint32) {
	p.streamingDistance = streamingDistance
	w.SetPedData(p.id, enum.PedStreamingDistance, int64(streamingDistance))
}

func (p *IPed) SetMaxHealth(maxHealth uint16) {
	if maxHealth > hash_enums.MaxHealth {
		p.maxHealth = maxHealth
		w.SetPedData(p.id, enum.PedMaxHealth, int64(maxHealth))
		return
	}
	p.maxHealth = maxHealth
	w.SetPedData(p.id, enum.PedMaxHealth, int64(maxHealth))
}

func (p *IPed) SetCurrentWeapon(currentWeapon weapon_hash.ModelHash) {
	p.currentWeapon = currentWeapon
	w.SetPedData(p.id, enum.PedCurrentWeapon, int64(currentWeapon))
}

func (p *IPed) SetCurrentWeaponByName(model string) {
	modelHash := weapon_hash.ModelHash(utils.Hash(model))
	p.currentWeapon = modelHash
	w.SetPedData(p.id, enum.PedCurrentWeapon, int64(modelHash))
}

func (p *IPed) Destroy() {
	w.SetPedData(p.id, enum.PedDestroy, int64(0))
	pools.DestroyPed(p)
}

func (p *IPed) SetData(key string, value any) {
	p.datas.Store(key, value)
}

func (p *IPed) DelData(key string) {
	_, ok := p.datas.Load(key)
	if ok {
		p.datas.Delete(key)
	}
}

func (p *IPed) DelAllData() {
	p.datas.Range(func(key, value any) bool {
		p.datas.Delete(key)
		return true
	})
}

func (p *IPed) HasData(key string) bool {
	_, ok := p.datas.Load(key)
	if ok {
		return true
	}
	return false
}

func (p *IPed) GetData(key string) any {
	value, ok := p.datas.Load(key)
	if ok {
		return value
	}
	return value
}

func (p *IPed) GetDatas() []any {
	var datas []any
	p.datas.Range(func(key, value any) bool {
		datas = append(datas, key)
		return true
	})
	return datas
}
