package models

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common"
	"github.com/StanZzzz222/RAltGo/common/core/alt/alt_timers"
	"github.com/StanZzzz222/RAltGo/common/core/scheduler"
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"github.com/StanZzzz222/RAltGo/hash_enums/ped_hash"
	"github.com/StanZzzz222/RAltGo/hash_enums/weapon_hash"
	"github.com/StanZzzz222/RAltGo/hash_enums/weather_hash"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enums"
	"math"
	"net"
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
	socialName         string
	socialId           uint64
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
	*BaseObject
	*NetworkData
	*EntityData
}

func (p *IPlayer) NewIPlayer(id uint32, name, ip, authToken, socialName string, socialId, hwIdHash, hwIdExHash uint64, position, rotation *entities.Vector3) *IPlayer {
	ipParse := net.ParseIP(ip)
	return &IPlayer{
		id:            id,
		name:          name,
		gameName:      name,
		ip:            &ipParse,
		authToken:     authToken,
		socialName:    socialName,
		socialId:      socialId,
		hwIdHash:      hwIdHash,
		hwIdExHash:    hwIdExHash,
		currentWeapon: weapon_hash.Fist,
		armour:        0,
		maxHealth:     200,
		BaseObject:    NewBaseObject(position, rotation, hash_enums.DefaultDimension, false, true, true),
		NetworkData:   NewNetworkData(id, enums.Player),
		EntityData:    NewEntityData(id, enums.Player),
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
func (p *IPlayer) GetHwIdHash() uint64                     { return p.hwIdHash }
func (p *IPlayer) GetHwIdExHash() uint64                   { return p.hwIdExHash }
func (p *IPlayer) GetSocialId() uint64                     { return p.socialId }
func (p *IPlayer) GetSocialName() string                   { return p.socialName }
func (p *IPlayer) GetGameName() string                     { return p.gameName }
func (p *IPlayer) GetHealth() uint16 {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerHealth))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U16Val
	}
	return 0
}
func (p *IPlayer) IsEnteringVehicle() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerIsEnteringVehicle))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsDead() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerIsDead))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsInVehicle() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerIsInVehicle))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsAiming() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerIsAiming))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsInCover() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerIsInCover))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsInRagdoll() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerIsInRagdoll))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsShooting() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerIsShooting))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsJumping() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerIsJumping))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsLeavingVehicle() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerIsLeavingVehicle))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) IsInMelle() bool {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerIsInMelle))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (p *IPlayer) Vehicle() *IVehicle {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerVehicle))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return pools.GetVehicle(cDataResult.U32Val)
	}
	return nil
}
func (p *IPlayer) GetPosition() *entities.Vector3 {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerPosition))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.Vector3Val
	}
	return nil
}
func (p *IPlayer) GetRotation() *entities.Vector3 {
	ret, freeDataResultFunc := w.GetData(p.id, enums.Player, uint8(enums.PlayerRotation))
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
	w.SetPlayerMetaModelData(p.id, enums.PlayerSpawn, uint32(model), int64(uint64(math.Float32bits(position.X))|(uint64(math.Float32bits(position.Y))<<32)), uint64(math.Float32bits(position.Z))<<32)
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
	w.SetPlayerData(p.id, enums.PlayerDespawn, int64(0))
}

func (p *IPlayer) ClearBloodDamage() {
	w.SetPlayerData(p.id, enums.PlayerClearBloodDamage, int64(0))
}

func (p *IPlayer) SetGameName(gameName string) {
	p.gameName = gameName
}

func (p *IPlayer) SetEyeColor(eyeColor int16) {
	p.eyeColor = eyeColor
	w.SetPlayerData(p.id, enums.PlayerEyeColor, int64(eyeColor))
}

func (p *IPlayer) SetHairColor(hairColor uint8) {
	p.hairColor = hairColor
	w.SetPlayerData(p.id, enums.PlayerHairColor, int64(hairColor))
}

func (p *IPlayer) SetHairHighlightColor(hairHighlightColor uint8) {
	p.hairHighlightColor = hairHighlightColor
	w.SetPlayerData(p.id, enums.PlayerHairHighlightColor, int64(hairHighlightColor))
}

func (p *IPlayer) SetHealth(health uint16) {
	if health > hash_enums.MaxHealth {
		p.health = health
		w.SetPlayerData(p.id, enums.PlayerHealth, int64(health))
		return
	}
	p.health = health
	w.SetPlayerData(p.id, enums.PlayerHealth, int64(health))
}

func (p *IPlayer) SetPosition(position *entities.Vector3) {
	p.position = position
	w.SetPlayerMetaData(p.id, enums.PlayerPosition, int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32)
}

func (p *IPlayer) SetDateTime(t time.Time) {
	p.time = t.UTC()
	w.SetPlayerData(p.id, enums.PlayerDateTime, t.UTC().Unix())
}

func (p *IPlayer) SetDateTimeUTC8(t time.Time) {
	p.time = t.UTC().Add(time.Hour * 8)
	w.SetPlayerData(p.id, enums.PlayerDateTime, t.UTC().Add(time.Hour*8).Unix())
}

func (p *IPlayer) SetWeather(wather weather_hash.WeatherType) {
	p.weather = wather
	w.SetPlayerData(p.id, enums.PlayerWeather, int64(wather))
}

func (p *IPlayer) SetMaxHealth(maxHealth uint16) {
	if maxHealth > hash_enums.MaxHealth {
		p.maxHealth = maxHealth
		w.SetPlayerData(p.id, enums.PlayerMaxHealth, int64(maxHealth))
		return
	}
	p.maxHealth = maxHealth
	w.SetPlayerData(p.id, enums.PlayerMaxHealth, int64(maxHealth))
}

func (p *IPlayer) SetMaxArmour(maxArmour uint16) {
	p.maxArmour = maxArmour
	w.SetPlayerData(p.id, enums.PlayerMaxArmour, int64(maxArmour))
}

func (p *IPlayer) SetAmmo(weapon string, ammo uint16) {
	w.SetPlayerMetaData(p.id, enums.PlayerAmmo, int64(common.Hash(weapon)), uint64(ammo))
}

func (p *IPlayer) SetMaxAmmo(weapon string, ammo uint16) {
	w.SetPlayerMetaData(p.id, enums.PlayerMaxAmmo, int64(common.Hash(weapon)), uint64(ammo))
}

func (p *IPlayer) SetCurrentWeapon(currentWeapon weapon_hash.ModelHash) {
	p.currentWeapon = currentWeapon
	w.SetPlayerData(p.id, enums.PlayerCurrentWeapon, int64(currentWeapon))
}

func (p *IPlayer) SetCurrentWeaponByName(model string) {
	modelHash := weapon_hash.ModelHash(common.Hash(model))
	if len(modelHash.String()) > 0 {
		p.currentWeapon = modelHash
		w.SetPlayerData(p.id, enums.PlayerCurrentWeapon, int64(modelHash))
	}
}

func (p *IPlayer) SetWeaponAmmo(weapon weapon_hash.ModelHash, ammo uint16) {
	w.SetPlayerMetaData(p.id, enums.PlayerWeaponAmmo, int64(weapon), uint64(ammo))
}

func (p *IPlayer) SetCurrentWeaponAmmo(ammo uint16) {
	if p.currentWeapon > 0 {
		p.SetWeaponAmmo(p.currentWeapon, ammo)
	}
}

func (p *IPlayer) SetDimension(dimension int32) {
	p.dimension = dimension
	w.SetPlayerData(p.id, enums.PlayerDimension, int64(dimension))
}

func (p *IPlayer) SetRotation(rotation *entities.Vector3) {
	p.rotation = rotation
	w.SetPlayerMetaData(p.id, enums.PlayerRotation, int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32)
}

func (p *IPlayer) SetPositionRotation(position, rotation *entities.Vector3) {
	p.position = position
	p.rotation = rotation
	w.SetPlayerMetaData(p.id, enums.PlayerPosition, int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32)
	w.SetPlayerMetaData(p.id, enums.PlayerRotation, int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32)
}

func (p *IPlayer) SetIntoVehicle(vehicle *IVehicle, seat uint8) {
	alt_timers.SetTimeout(time.Millisecond*30, func() {
		w.SetPlayerMetaData(p.id, enums.PlayerInVehicle, int64(vehicle.id), uint64(seat))
	})
}

func (p *IPlayer) SetArmour(armour uint16) {
	p.armour = armour
	w.SetPlayerData(p.id, enums.PlayerArmour, int64(armour))
}

func (p *IPlayer) SetPedModel(model ped_hash.ModelHash) {
	p.model = model
	w.SetPlayerData(p.id, enums.PlayerModel, int64(model))
}

func (p *IPlayer) SetPedModelByName(model string) {
	modelHash := ped_hash.ModelHash(common.Hash(model))
	if len(modelHash.String()) > 0 {
		p.model = modelHash
		w.SetPlayerData(p.id, enums.PlayerModel, int64(modelHash))
	}
}

func (p *IPlayer) SetFrozen(frozen bool) {
	p.frozen = frozen
	value := 0
	if frozen {
		value = 1
	}
	w.SetPlayerData(p.id, enums.PlayerFrozen, int64(value))
}

func (p *IPlayer) SetVisible(visible bool) {
	p.visible = visible
	value := 0
	if visible {
		value = 1
	}
	w.SetPlayerData(p.id, enums.PlayerVisible, int64(value))
}

func (p *IPlayer) SetCollision(collision bool) {
	p.collision = collision
	value := 0
	if collision {
		value = 1
	}
	w.SetPlayerData(p.id, enums.PlayerCollision, int64(value))
}

func (p *IPlayer) SetInvincible(invincible bool) {
	p.invincible = invincible
	value := 0
	if invincible {
		value = 1
	}
	w.SetPlayerData(p.id, enums.PlayerInvincible, int64(value))
}

func (p *IPlayer) ClearProps(compoent uint8) {
	w.SetPlayerData(p.id, enums.PlayerClearProps, int64(compoent))
}

func (p *IPlayer) ClearDecorations() {
	w.SetPlayerData(p.id, enums.PlayerClearDecorations, int64(0))
}

func (p *IPlayer) ClearTasks() {
	w.SetPlayerData(p.id, enums.PlayerClearTasks, int64(0))
}

func (p *IPlayer) Kick(reason string) {
	w.SetServerData(enums.KickPlayer, int64(p.GetId()), reason)
}

func (p *IPlayer) ResetHeadBlendData() {
	w.SetPlayerData(p.id, enums.PlayerResetHeadBlendData, int64(0))
}

func (p *IPlayer) SetHeadBlendData(shapeFirstId, shapeSecondId, shapeThirdId, skinFirstId, skinSecondId, skinThirdId uint32, shapeMix, skinMix, thirdMix float32) {
	w.SetPlayerHeadData(p.id, enums.PlayerHeadBlendData, shapeFirstId, shapeSecondId, shapeThirdId, skinFirstId, skinSecondId, skinThirdId, shapeMix, skinMix, thirdMix)
}

func (p *IPlayer) SetHeadOverlay(overlayId uint32, index uint8, opacity float32) {
	w.SetPlayerHeadData(p.id, enums.PlayerHeadOverlay, overlayId, uint32(index), uint32(0), uint32(0), uint32(0), uint32(0), opacity, float32(0), float32(0))
}

func (p *IPlayer) SetHeadOverlayColor(overlayId, colorType, colorIndex, secondColorIndex uint8) {
	w.SetPlayerHeadData(p.id, enums.PlayerHeadOverlayColor, uint32(overlayId), uint32(colorType), uint32(colorIndex), uint32(secondColorIndex), uint32(0), uint32(0), float32(0), float32(0), float32(0))
}

func (p *IPlayer) SetHeadBlendPaletteColor(id, r, g, b uint8) {
	w.SetPlayerHeadData(p.id, enums.PlayerHeadBlendPaletteColor, uint32(id), uint32(r), uint32(g), uint32(b), uint32(0), uint32(0), float32(0), float32(0), float32(0))
}

func (p *IPlayer) SendBroadcastMessage(message string) {
	p.Emit("chat:message", "", message)
}
