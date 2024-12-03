package models

import (
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"github.com/StanZzzz222/RAltGo/hash_enums/ped_hash"
	"github.com/StanZzzz222/RAltGo/hash_enums/weapon_hash"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enums"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/13
   File: ped_hash.go
*/

type IPed struct {
	id            uint32
	model         ped_hash.ModelHash
	health        uint16
	armour        uint16
	maxHealth     uint16
	currentWeapon weapon_hash.ModelHash
	warpper       *lib.Wrapper
	*BaseObject
	*NetworkData
	*EntityData
}

func (p *IPed) NewIPed(id, model uint32, position, rotation *Vector3) *IPed {
	return &IPed{
		id:            id,
		model:         ped_hash.ModelHash(model),
		currentWeapon: weapon_hash.Fist,
		armour:        0,
		maxHealth:     200,
		warpper:       lib.GetWrapper(),
		BaseObject:    NewBaseObject(position, rotation, hash_enums.DefaultDimension, false, true, true),
		NetworkData:   NewNetworkData(id, enums.Ped),
		EntityData:    NewEntityData(id, enums.Ped),
	}
}

func (p *IPed) GetId() uint32                           { return p.id }
func (p *IPed) GetModel() ped_hash.ModelHash            { return p.model }
func (p *IPed) GetCurrentWeapon() weapon_hash.ModelHash { return p.currentWeapon }
func (p *IPed) GetMaxHealth() uint16                    { return p.maxHealth }
func (p *IPed) GetDimension() int32                     { return p.dimension }
func (p *IPed) GetFrozen() bool                         { return p.frozen }
func (p *IPed) GetCollision() bool                      { return p.collision }
func (p *IPed) GetVisible() bool                        { return p.visible }
func (p *IPed) GetPosition() *Vector3 {
	ret, freeDataResultFunc := p.warpper.GetData(p.id, enums.Ped, uint8(enums.PedPosition))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return (*Vector3)(cDataResult.Vector3Val)
	}
	return nil
}
func (p *IPed) GetRotation() *Vector3 {
	ret, freeDataResultFunc := p.warpper.GetData(p.id, enums.Ped, uint8(enums.PedRotation))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return (*Vector3)(cDataResult.Vector3Val)
	}
	return nil
}

func (p *IPed) SetPosition(position *Vector3) {
	p.position = position
	p.warpper.SetPedMetaData(p.id, enums.PedPosition, int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32)
}

func (p *IPed) SetRotation(rotation *Vector3) {
	p.rotation = rotation
	p.warpper.SetPedMetaData(p.id, enums.PedRotation, int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32)
}

func (p *IPed) SetHealth(health uint16) {
	if health > hash_enums.MaxHealth {
		p.health = health
		p.warpper.SetPedData(p.id, enums.PedMaxHealth, int64(health))
		return
	}
	p.health = health
	p.warpper.SetPedData(p.id, enums.PedHealth, int64(health))
}

func (p *IPed) SetArmour(armour uint16) {
	p.armour = armour
	p.warpper.SetPedData(p.id, enums.PedArmour, int64(armour))
}

func (p *IPed) SetDimension(dimension int32) {
	p.dimension = dimension
	p.warpper.SetPedData(p.id, enums.PedDimension, int64(dimension))
}

func (p *IPed) SetCollision(collision bool) {
	p.collision = collision
	value := 0
	if collision {
		value = 1
	}
	p.warpper.SetPedData(p.id, enums.PedCollision, int64(value))
}

func (p *IPed) SetFrozen(frozen bool) {
	p.frozen = frozen
	value := 0
	if frozen {
		value = 1
	}
	p.warpper.SetPedData(p.id, enums.PedFrozen, int64(value))
}

func (p *IPed) SetVisible(visible bool) {
	p.visible = visible
	value := 0
	if visible {
		value = 1
	}
	p.warpper.SetPedData(p.id, enums.PedVisible, int64(value))
}

func (p *IPed) SetMaxHealth(maxHealth uint16) {
	if maxHealth > hash_enums.MaxHealth {
		p.maxHealth = maxHealth
		p.warpper.SetPedData(p.id, enums.PedMaxHealth, int64(maxHealth))
		return
	}
	p.maxHealth = maxHealth
	p.warpper.SetPedData(p.id, enums.PedMaxHealth, int64(maxHealth))
}

func (p *IPed) SetCurrentWeapon(currentWeapon weapon_hash.ModelHash) {
	p.currentWeapon = currentWeapon
	p.warpper.SetPedData(p.id, enums.PedCurrentWeapon, int64(currentWeapon))
}

func (p *IPed) SetCurrentWeaponByName(model string) {
	modelHash := weapon_hash.ModelHash(hash(model))
	p.currentWeapon = modelHash
	p.warpper.SetPedData(p.id, enums.PedCurrentWeapon, int64(modelHash))
}

func (p *IPed) Destroy() {
	p.warpper.SetPedData(p.id, enums.PedDestroy, int64(0))
	pools.DestroyPed(p)
}
