package models

import (
	"gamemode/enums/ped"
	"gamemode/enums/weapon"
	"gamemode/internal/enum"
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
	dimension     int32
	currentWeapon string
	frozen        bool
	invincible    bool
	collision     bool
	position      *Vector3
	rotation      *Vector3
	time          time.Time
}

func (p *IPlayer) NewIPlayer(id uint32, name, ip, authToken string, hwIdHash, hwIdExHash uint64, position, rotation *Vector3) *IPlayer {
	ipParse := net.ParseIP(ip)
	return &IPlayer{
		id:         id,
		name:       name,
		ip:         &ipParse,
		authToken:  authToken,
		hwIdHash:   hwIdHash,
		hwIdExHash: hwIdExHash,
		position:   position,
		rotation:   rotation,
	}
}

func (p *IPlayer) Spawn(model string, position *Vector3) {
	if position != nil {
		p.model = hash(model)
		p.position = position
		w.SpawnPlayer(p.id, hash(model), position.X, position.Y, position.Z)
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
	w.SetPlayerData(p.id, enum.Health, uint64(health))
}

func (p *IPlayer) SetPosition(position *Vector3) {
	p.position = position
	w.SetPlayerMetaData(p.id, enum.Positon, uint64(math.Float32bits(position.X))|(uint64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32)
}

func (p *IPlayer) SetDateTime(t time.Time) {
	p.time = t.UTC()
	w.SetPlayerData(p.id, enum.DateTime, uint64(t.UTC().Unix()))
}

func (p *IPlayer) SetDateTimeUTC8(t time.Time) {
	p.time = t.UTC().Add(time.Hour * 8)
	w.SetPlayerData(p.id, enum.DateTime, uint64(t.UTC().Add(time.Hour*8).Unix()))
}

func (p *IPlayer) SetWeather(wather uint16) {
	p.weather = wather
	w.SetPlayerData(p.id, enum.Weather, uint64(wather))
}

func (p *IPlayer) SetMaxHealth(maxHealth uint16) {
	p.maxHealth = maxHealth
	w.SetPlayerData(p.id, enum.MaxHealth, uint64(maxHealth))
}

func (p *IPlayer) SetMaxArmour(maxArmour uint16) {
	p.maxArmour = maxArmour
	w.SetPlayerData(p.id, enum.MaxArmour, uint64(maxArmour))
}

func (p *IPlayer) SetAmmo(weapon string, ammo uint16) {
	w.SetPlayerMetaData(p.id, enum.Ammo, uint64(hash(weapon)), uint64(ammo))
}

func (p *IPlayer) SetMaxAmmo(weapon string, ammo uint16) {
	w.SetPlayerMetaData(p.id, enum.MaxAmmo, uint64(hash(weapon)), uint64(ammo))
}

func (p *IPlayer) SetCurrentWeapon(weapon string) {
	p.currentWeapon = strings.ToLower(weapon)
	w.SetPlayerData(p.id, enum.CurrentWeapon, uint64(hash(weapon)))
}

func (p *IPlayer) SetCurrentWeaponByHash(weaponHash weapon.ModelHash) {
	p.currentWeapon = strings.ToLower(weaponHash.String())
	w.SetPlayerData(p.id, enum.CurrentWeapon, uint64(weaponHash))
}

func (p *IPlayer) SetWeaponAmmo(weapon string, ammo uint16) {
	w.SetPlayerMetaData(p.id, enum.WeaponAmmo, uint64(hash(weapon)), uint64(ammo))
}

func (p *IPlayer) SetCurrentWeaponAmmo(ammo uint16) {
	if len(p.currentWeapon) > 0 {
		p.SetWeaponAmmo(p.currentWeapon, ammo)
	}
}

func (p *IPlayer) SetDimension(dimension int32) {
	if dimension > 0 {
		p.dimension = dimension
		w.SetPlayerData(p.id, enum.Dimension, uint64(dimension))
	}
}

func (p *IPlayer) SetRotation(rotation *Vector3) {
	p.rotation = rotation
	w.SetPlayerMetaData(p.id, enum.Rot, uint64(math.Float32bits(rotation.X))|(uint64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32)
}

func (p *IPlayer) SetPositionRotation(position, rotation *Vector3) {
	p.position = position
	p.rotation = rotation
	w.SetPlayerMetaData(p.id, enum.Positon, uint64(math.Float32bits(position.X))|(uint64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32)
	w.SetPlayerMetaData(p.id, enum.Rot, uint64(math.Float32bits(rotation.X))|(uint64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32)
}

func (p *IPlayer) SetArmour(armour uint16) {
	p.armour = armour
	w.SetPlayerData(p.id, enum.Armour, uint64(armour))
}

func (p *IPlayer) SetPedModel(model string) {
	p.model = hash(model)
	w.SetPlayerData(p.id, enum.Model, uint64(hash(model)))
}

func (p *IPlayer) SetPedModelByHash(modelHash ped.ModelHash) {
	p.model = uint32(modelHash)
	w.SetPlayerData(p.id, enum.Model, uint64(modelHash))
}

func (p *IPlayer) SetFrozen(frozen bool) {
	p.frozen = frozen
	value := 0
	if frozen {
		value = 1
	}
	w.SetPlayerData(p.id, enum.Frozen, uint64(value))
}

func (p *IPlayer) SetCollision(collision bool) {
	p.collision = collision
	value := 0
	if collision {
		value = 1
	}
	w.SetPlayerData(p.id, enum.Collision, uint64(value))
}

func (p *IPlayer) SetInvincible(invincible bool) {
	p.invincible = invincible
	value := 0
	if invincible {
		value = 1
	}
	w.SetPlayerData(p.id, enum.Invincible, uint64(value))
}
