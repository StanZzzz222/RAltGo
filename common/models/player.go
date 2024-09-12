package models

import (
	"github.com/StanZzzz222/RAltGo/common/utils"
	"github.com/StanZzzz222/RAltGo/enums"
	"github.com/StanZzzz222/RAltGo/enums/ped"
	"github.com/StanZzzz222/RAltGo/enums/weapon"
	"github.com/StanZzzz222/RAltGo/enums/weather"
	"github.com/StanZzzz222/RAltGo/internal/entitys"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"github.com/StanZzzz222/RAltGo/timers"
	"math"
	"net"
	"strings"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/9/5
   File: player.go
*/

type IPlayer struct {
	id            uint32
	name          string
	ip            *net.IP
	authToken     string
	hwIdHash      uint64
	hwIdExHash    uint64
	model         uint32
	health        uint16
	armour        uint16
	weather       uint16
	maxHealth     uint16
	maxArmour     uint16
	currentWeapon string
	invincible    bool
	time          time.Time
	*BaseObject
}

func (p *IPlayer) NewIPlayer(id uint32, name, ip, authToken string, hwIdHash, hwIdExHash uint64, position, rotation *entitys.Vector3) *IPlayer {
	ipParse := net.ParseIP(ip)
	return &IPlayer{
		id:         id,
		name:       name,
		ip:         &ipParse,
		authToken:  authToken,
		hwIdHash:   hwIdHash,
		hwIdExHash: hwIdExHash,
		BaseObject: NewBaseObject(position, rotation, enums.DefaultDimension, false, true, true),
	}
}

func (p *IPlayer) Spawn(model string, position *entitys.Vector3) {
	if position != nil {
		p.model = utils.Hash(model)
		p.position = position
		w.SpawnPlayer(p.id, utils.Hash(model), position.X, position.Y, position.Z)
	}
}

func (p *IPlayer) GetId() uint32        { return p.id }
func (p *IPlayer) GetName() string      { return p.name }
func (p *IPlayer) GetIP() *net.IP       { return p.ip }
func (p *IPlayer) GetModel() uint32     { return p.model }
func (p *IPlayer) GetMaxHealth() uint16 { return p.maxHealth }
func (p *IPlayer) GetMaxArmour() uint16 { return p.maxArmour }
func (p *IPlayer) GetDimension() int32  { return p.dimension }
func (p *IPlayer) GetFrozen() bool      { return p.frozen }
func (p *IPlayer) GetCollision() bool   { return p.collision }
func (p *IPlayer) GetInvincible() bool  { return p.invincible }

// func (p *IPlayer) GetHealth() uint16            { return p.health }
// func (p *IPlayer) GetCurrentWeapon() string     { return p.currentWeapon }
// func (p *IPlayer) GetCurrentWeaponHash() uint32 { return utils.Hash(p.currentWeapon) }
//func (p *IPlayer) GetPosition() *Vector3 { return p.position }
//func (p *IPlayer) GetPositionString() string {
//	return fmt.Sprintf("%v,%v,%v", p.position.X, p.position.Y, p.position.Z)
//}
//func (p *IPlayer) GetRotation() *Vector3 { return p.rotation }
//func (p *IPlayer) GetRotationString() string {
//	return fmt.Sprintf("%v,%v,%v", p.rotation.X, p.rotation.Y, p.rotation.Z)
//}
//func (p *IPlayer) GetPositionRotation() (*Vector3, *Vector3) { return p.rotation, p.rotation }
//func (p *IPlayer) GetPositionRotationString() (string, string) {
//	return fmt.Sprintf("%v,%v,%v", p.rotation.X, p.rotation.Y, p.rotation.Z), fmt.Sprintf("%v,%v,%v", p.rotation.X, p.rotation.Y, p.rotation.Z)
//}

func (p *IPlayer) SetHealth(health uint16) {
	p.health = health
	w.SetPlayerData(p.id, enum.Health, int64(health))
}

func (p *IPlayer) SetPosition(position *entitys.Vector3) {
	p.position = position
	w.SetPlayerMetaData(p.id, enum.Position, int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32)
}

func (p *IPlayer) SetDateTime(t time.Time) {
	p.time = t.UTC()
	w.SetPlayerData(p.id, enum.DateTime, t.UTC().Unix())
}

func (p *IPlayer) SetDateTimeUTC8(t time.Time) {
	p.time = t.UTC().Add(time.Hour * 8)
	w.SetPlayerData(p.id, enum.DateTime, t.UTC().Add(time.Hour*8).Unix())
}

func (p *IPlayer) SetWeather(wather weather.WeatherType) {
	p.weather = uint16(wather)
	w.SetPlayerData(p.id, enum.Weather, int64(wather))
}

func (p *IPlayer) SetMaxHealth(maxHealth uint16) {
	p.maxHealth = maxHealth
	w.SetPlayerData(p.id, enum.MaxHealth, int64(maxHealth))
}

func (p *IPlayer) SetMaxArmour(maxArmour uint16) {
	p.maxArmour = maxArmour
	w.SetPlayerData(p.id, enum.MaxArmour, int64(maxArmour))
}

func (p *IPlayer) SetAmmo(weapon string, ammo uint16) {
	w.SetPlayerMetaData(p.id, enum.Ammo, int64(utils.Hash(weapon)), uint64(ammo))
}

func (p *IPlayer) SetMaxAmmo(weapon string, ammo uint16) {
	w.SetPlayerMetaData(p.id, enum.MaxAmmo, int64(utils.Hash(weapon)), uint64(ammo))
}

func (p *IPlayer) SetCurrentWeapon(weapon string) {
	p.currentWeapon = strings.ToLower(weapon)
	w.SetPlayerData(p.id, enum.CurrentWeapon, int64(utils.Hash(weapon)))
}

func (p *IPlayer) SetCurrentWeaponByHash(weaponHash weapon.ModelHash) {
	p.currentWeapon = strings.ToLower(weaponHash.String())
	w.SetPlayerData(p.id, enum.CurrentWeapon, int64(weaponHash))
}

func (p *IPlayer) SetWeaponAmmo(weapon string, ammo uint16) {
	w.SetPlayerMetaData(p.id, enum.WeaponAmmo, int64(utils.Hash(weapon)), uint64(ammo))
}

func (p *IPlayer) SetCurrentWeaponAmmo(ammo uint16) {
	if len(p.currentWeapon) > 0 {
		p.SetWeaponAmmo(p.currentWeapon, ammo)
	}
}

func (p *IPlayer) SetDimension(dimension int32) {
	if dimension > 0 {
		p.dimension = dimension
		w.SetPlayerData(p.id, enum.Dimension, int64(dimension))
	}
}

func (p *IPlayer) SetRotation(rotation *entitys.Vector3) {
	p.rotation = rotation
	w.SetPlayerMetaData(p.id, enum.Rot, int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32)
}

func (p *IPlayer) SetPositionRotation(position, rotation *entitys.Vector3) {
	p.position = position
	p.rotation = rotation
	w.SetPlayerMetaData(p.id, enum.Position, int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32)
	w.SetPlayerMetaData(p.id, enum.Rot, int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32)
}

func (p *IPlayer) SetIntoVehicle(vehicle *IVehicle, seat uint8) {
	timers.SetTimeout(time.Millisecond*30, func() {
		w.SetPlayerMetaData(p.id, enum.InVehicle, int64(vehicle.id), uint64(seat))
	})
}

func (p *IPlayer) SetArmour(armour uint16) {
	p.armour = armour
	w.SetPlayerData(p.id, enum.Armour, int64(armour))
}

func (p *IPlayer) SetPedModel(model string) {
	p.model = utils.Hash(model)
	w.SetPlayerData(p.id, enum.Model, int64(utils.Hash(model)))
}

func (p *IPlayer) SetPedModelByHash(modelHash ped.ModelHash) {
	p.model = uint32(modelHash)
	w.SetPlayerData(p.id, enum.Model, int64(modelHash))
}

func (p *IPlayer) SetFrozen(frozen bool) {
	p.frozen = frozen
	value := 0
	if frozen {
		value = 1
	}
	w.SetPlayerData(p.id, enum.Frozen, int64(value))
}

func (p *IPlayer) SetVisible(visible bool) {
	p.visible = visible
	value := 0
	if visible {
		value = 1
	}
	w.SetPlayerData(p.id, enum.Visible, int64(value))
}

func (p *IPlayer) SetCollision(collision bool) {
	p.collision = collision
	value := 0
	if collision {
		value = 1
	}
	w.SetPlayerData(p.id, enum.Collision, int64(value))
}

func (p *IPlayer) SetInvincible(invincible bool) {
	p.invincible = invincible
	value := 0
	if invincible {
		value = 1
	}
	w.SetPlayerData(p.id, enum.Invincible, int64(value))
}
