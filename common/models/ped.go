package models

import (
	"github.com/StanZzzz222/RAltGo/common/utils"
	"github.com/StanZzzz222/RAltGo/enums"
	"github.com/StanZzzz222/RAltGo/enums/ped"
	"github.com/StanZzzz222/RAltGo/enums/weapon"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/13
   File: ped.go
*/

type IPed struct {
	id                uint32
	model             ped.ModelHash
	health            uint16
	armour            uint16
	maxHealth         uint16
	currentWeapon     weapon.ModelHash
	streamingDistance uint32
	streamed          bool
	*BaseObject
}

func (p *IPed) GetId() uint32                      { return p.id }
func (p *IPed) GetModel() ped.ModelHash            { return p.model }
func (p *IPed) GetCurrentWeapon() weapon.ModelHash { return p.currentWeapon }
func (p *IPed) GetStreamingDistance() uint32       { return p.streamingDistance }
func (p *IPed) GetStreamed() bool                  { return p.streamed }
func (p *IPed) GetMaxHealth() uint16               { return p.maxHealth }
func (p *IPed) GetDimension() int32                { return p.dimension }
func (p *IPed) GetFrozen() bool                    { return p.frozen }
func (p *IPed) GetCollision() bool                 { return p.collision }
func (p *IPed) GetVisible() bool                   { return p.visible }

func (p *IPed) NewIPed(id uint32, model uint32, position, rotation *entities.Vector3) *IPed {
	return &IPed{
		id:            id,
		model:         ped.ModelHash(model),
		currentWeapon: weapon.Fist,
		BaseObject:    NewBaseObject(position, rotation, enums.DefaultDimension, false, true, true),
	}
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
	p.maxHealth = maxHealth
	w.SetPedData(p.id, enum.PedMaxHealth, int64(maxHealth))
}

func (p *IPed) SetCurrentWeapon(currentWeapon weapon.ModelHash) {
	p.currentWeapon = currentWeapon
	w.SetPedData(p.id, enum.PedCurrentWeapon, int64(currentWeapon))
}

func (p *IPed) SetCurrentWeaponByName(model string) {
	modelHash := weapon.ModelHash(utils.Hash(model))
	p.currentWeapon = modelHash
	w.SetPedData(p.id, enum.PedCurrentWeapon, int64(modelHash))
}