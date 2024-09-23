package models

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/alt/broadcast"
	"github.com/StanZzzz222/RAltGo/common/alt/scheduler"
	"github.com/StanZzzz222/RAltGo/common/alt/timers"
	"github.com/StanZzzz222/RAltGo/common/utils"
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"github.com/StanZzzz222/RAltGo/hash_enums/ped_hash"
	"github.com/StanZzzz222/RAltGo/hash_enums/weapon_hash"
	"github.com/StanZzzz222/RAltGo/hash_enums/weather_hash"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"math"
	"net"
	"sync"
	"time"
)

/*
   Create by zyx
   Date Time: 2024/9/5
   File: player.go
*/

type IPlayer struct {
	id                 uint32
	name               string
	gameName           string
	ip                 *net.IP
	authToken          string
	hwIdHash           uint64
	hwIdExHash         uint64
	model              ped_hash.ModelHash
	health             uint16
	armour             uint16
	weather            weather_hash.WeatherType
	maxHealth          uint16
	maxArmour          uint16
	eyeColor           int16
	hairColor          uint8
	hairHighlightColor uint8
	currentWeapon      weapon_hash.ModelHash
	invincible         bool
	time               time.Time
	datas              *sync.Map
	*BaseObject
}

func (p *IPlayer) NewIPlayer(id uint32, name, ip, authToken string, hwIdHash, hwIdExHash uint64, position, rotation *entities.Vector3) *IPlayer {
	ipParse := net.ParseIP(ip)
	return &IPlayer{
		id:            id,
		name:          name,
		gameName:      name,
		ip:            &ipParse,
		authToken:     authToken,
		hwIdHash:      hwIdHash,
		hwIdExHash:    hwIdExHash,
		currentWeapon: weapon_hash.Fist,
		datas:         &sync.Map{},
		BaseObject:    NewBaseObject(position, rotation, hash_enums.DefaultDimension, false, true, true),
	}
}

func (p *IPlayer) GetId() uint32                           { return p.id }
func (p *IPlayer) GetName() string                         { return p.name }
func (p *IPlayer) GetIP() *net.IP                          { return p.ip }
func (p *IPlayer) GetModel() ped_hash.ModelHash            { return p.model }
func (p *IPlayer) GetCurrentWeapon() weapon_hash.ModelHash { return p.currentWeapon }
func (p *IPlayer) GetMaxHealth() uint16                    { return p.maxHealth }
func (p *IPlayer) GetMaxArmour() uint16                    { return p.maxArmour }
func (p *IPlayer) GetDimension() int32                     { return p.dimension }
func (p *IPlayer) GetFrozen() bool                         { return p.frozen }
func (p *IPlayer) GetCollision() bool                      { return p.collision }
func (p *IPlayer) GetInvincible() bool                     { return p.invincible }
func (p *IPlayer) GetGameName() string                     { return p.gameName }
func (p *IPlayer) GetHealth() uint16 {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.Health))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U16Val
	}
	return 0
}
func (p *IPlayer) IsEnteringVehicle() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.IsEnteringVehicle))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsDead() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.IsDead))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsInVehicle() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.IsInVehicle))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsAiming() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.IsAiming))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsInCover() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.IsInCover))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsInRagdoll() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.IsInRagdoll))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsShooting() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.IsShooting))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsJumping() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.IsJumping))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsLeavingVehicle() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.IsLeavingVehicle))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsInMelle() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.IsInMelle))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) Vehicle() *IVehicle {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.PlayerVehicle))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return pools.GetVehicle(cDataResult.U32Val)
	}
	return nil
}
func (p *IPlayer) GetPosition() *entities.Vector3 {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.Position))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.Vector3Val
	}
	return nil
}
func (p *IPlayer) GetRotation() *entities.Vector3 {
	ret, freeDataResultFunc := w.GetData(p.id, enum.Player, uint8(enum.Rot))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.Vector3Val
	}
	return nil
}
func (p *IPlayer) GetPositionString() string {
	position := p.GetPosition()
	return fmt.Sprintf("%v,%v,%v", position.X, position.Y, position.Z)
}
func (p *IPlayer) GetRotationString() string {
	rotation := p.GetRotation()
	return fmt.Sprintf("%v,%v,%v", rotation.X, rotation.Y, rotation.Z)
}
func (p *IPlayer) GetPositionRotation() (*entities.Vector3, *entities.Vector3) {
	return p.GetPosition(), p.GetRotation()
}
func (p *IPlayer) GetPositionRotationString() (string, string) {
	return p.GetPositionString(), p.GetRotationString()
}

func (p *IPlayer) Spawn(model ped_hash.ModelHash, position *entities.Vector3) {
	p.model = model
	p.position = position
	w.SetPlayerMetaModelData(p.id, enum.Spawn, uint32(model), int64(uint64(math.Float32bits(position.X))|(uint64(math.Float32bits(position.Y))<<32)), uint64(math.Float32bits(position.Z))<<32)
}

func (p *IPlayer) Emit(eventName string, args ...any) {
	s := scheduler.NewScheduler()
	s.AddTask(func() {
		mvalues := NewMValues(args...)
		w.Emit(p.id, eventName, mvalues.Dump())
	})
	s.Run()
}

func (p *IPlayer) Despawn() {
	w.SetPlayerData(p.id, enum.Despawn, int64(0))
}

func (p *IPlayer) ClearBloodDamage() {
	w.SetPlayerData(p.id, enum.ClearBloodDamage, int64(0))
}

func (p *IPlayer) SetGameName(gameName string) {
	p.gameName = gameName
}

func (p *IPlayer) SetEyeColor(eyeColor int16) {
	p.eyeColor = eyeColor
	w.SetPlayerData(p.id, enum.EyeColor, int64(eyeColor))
}

func (p *IPlayer) SetHairColor(hairColor uint8) {
	p.hairColor = hairColor
	w.SetPlayerData(p.id, enum.HairColor, int64(hairColor))
}

func (p *IPlayer) SetHairHighlightColor(hairHighlightColor uint8) {
	p.hairHighlightColor = hairHighlightColor
	w.SetPlayerData(p.id, enum.HairHighlightColor, int64(hairHighlightColor))
}

func (p *IPlayer) SetHealth(health uint16) {
	if health > hash_enums.MaxHealth {
		p.health = health
		w.SetPlayerData(p.id, enum.Health, int64(health))
		return
	}
	p.health = health
	w.SetPlayerData(p.id, enum.Health, int64(health))
}

func (p *IPlayer) SetPosition(position *entities.Vector3) {
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

func (p *IPlayer) SetWeather(wather weather_hash.WeatherType) {
	p.weather = wather
	w.SetPlayerData(p.id, enum.Weather, int64(wather))
}

func (p *IPlayer) SetMaxHealth(maxHealth uint16) {
	if maxHealth > hash_enums.MaxHealth {
		p.maxHealth = maxHealth
		w.SetPlayerData(p.id, enum.MaxHealth, int64(maxHealth))
		return
	}
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

func (p *IPlayer) SetCurrentWeapon(currentWeapon weapon_hash.ModelHash) {
	p.currentWeapon = currentWeapon
	w.SetPlayerData(p.id, enum.CurrentWeapon, int64(currentWeapon))
}

func (p *IPlayer) SetCurrentWeaponByName(model string) {
	modelHash := weapon_hash.ModelHash(utils.Hash(model))
	if len(modelHash.String()) > 0 {
		p.currentWeapon = modelHash
		w.SetPlayerData(p.id, enum.CurrentWeapon, int64(modelHash))
	}
}

func (p *IPlayer) SetWeaponAmmo(weapon weapon_hash.ModelHash, ammo uint16) {
	w.SetPlayerMetaData(p.id, enum.WeaponAmmo, int64(weapon), uint64(ammo))
}

func (p *IPlayer) SetCurrentWeaponAmmo(ammo uint16) {
	if p.currentWeapon > 0 {
		p.SetWeaponAmmo(p.currentWeapon, ammo)
	}
}

func (p *IPlayer) SetDimension(dimension int32) {
	if dimension > 0 {
		p.dimension = dimension
		w.SetPlayerData(p.id, enum.Dimension, int64(dimension))
	}
}

func (p *IPlayer) SetRotation(rotation *entities.Vector3) {
	p.rotation = rotation
	w.SetPlayerMetaData(p.id, enum.Rot, int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32)
}

func (p *IPlayer) SetPositionRotation(position, rotation *entities.Vector3) {
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

func (p *IPlayer) SetPedModel(model ped_hash.ModelHash) {
	p.model = model
	w.SetPlayerData(p.id, enum.Model, int64(model))
}

func (p *IPlayer) SetPedModelByName(model string) {
	modelHash := ped_hash.ModelHash(utils.Hash(model))
	if len(modelHash.String()) > 0 {
		p.model = modelHash
		w.SetPlayerData(p.id, enum.Model, int64(modelHash))
	}
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

func (p *IPlayer) ClearProps(compoent uint8) {
	w.SetPlayerData(p.id, enum.ClearProps, int64(compoent))
}

func (p *IPlayer) ClearDecorations() {
	w.SetPlayerData(p.id, enum.ClearDecorations, int64(0))
}

func (p *IPlayer) ClearTasks() {
	w.SetPlayerData(p.id, enum.ClearTasks, int64(0))
}

func (p *IPlayer) ResetHeadBlendData() {
	w.SetPlayerData(p.id, enum.ResetHeadBlendData, int64(0))
}

func (p *IPlayer) SetHeadBlendData(shapeFirstId, shapeSecondId, shapeThirdId, skinFirstId, skinSecondId, skinThirdId uint32, shapeMix, skinMix, thirdMix float32) {
	w.SetPlayerHeadData(p.id, enum.HeadBlendData, shapeFirstId, shapeSecondId, shapeThirdId, skinFirstId, skinSecondId, skinThirdId, shapeMix, skinMix, thirdMix)
}

func (p *IPlayer) SetHeadOverlay(overlayId uint32, index uint8, opacity float32) {
	w.SetPlayerHeadData(p.id, enum.HeadOverlay, overlayId, uint32(index), uint32(0), uint32(0), uint32(0), uint32(0), opacity, float32(0), float32(0))
}

func (p *IPlayer) SetHeadOverlayColor(overlayId, colorType, colorIndex, secondColorIndex uint8) {
	w.SetPlayerHeadData(p.id, enum.HeadOverlayColor, uint32(overlayId), uint32(colorType), uint32(colorIndex), uint32(secondColorIndex), uint32(0), uint32(0), float32(0), float32(0), float32(0))
}

func (p *IPlayer) SetHeadBlendPaletteColor(id, r, g, b uint8) {
	w.SetPlayerHeadData(p.id, enum.HeadBlendPaletteColor, uint32(id), uint32(r), uint32(g), uint32(b), uint32(0), uint32(0), float32(0), float32(0), float32(0))
}

func (p *IPlayer) SendBroadcastMessage(message string) {
	p.Emit("chat:message", "", message)
}

func (p *IPlayer) SendBroadcastAll(message string) {
	broadcast.SendBroadcast(message)
}

func (p *IPlayer) SetData(key string, value any) {
	p.datas.Store(key, value)
}

func (p *IPlayer) DelData(key string) {
	_, ok := p.datas.Load(key)
	if ok {
		p.datas.Delete(key)
	}
}

func (p *IPlayer) DelAllData() {
	p.datas.Range(func(key, value any) bool {
		p.datas.Delete(key)
		return true
	})
}

func (p *IPlayer) HasData(key string) bool {
	_, ok := p.datas.Load(key)
	if ok {
		return true
	}
	return false
}

func (p *IPlayer) GetData(key string) any {
	value, ok := p.datas.Load(key)
	if ok {
		return value
	}
	return value
}

func (p *IPlayer) GetDatas() []any {
	var datas []any
	p.datas.Range(func(key, value any) bool {
		datas = append(datas, key)
		return true
	})
	return datas
}
