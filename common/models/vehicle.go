package models

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_hash"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"math"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/5
   File: vehicle_hash.go
*/

type IVehicle struct {
	id               uint32
	model            vehicle_hash.ModelHash
	primaryColor     uint8
	secondColor      uint8
	numberplate      string
	engineOn         bool
	neonActive       bool
	driftMode        bool
	disableTowing    bool
	lockState        uint8
	lightState       uint8
	headLightColor   uint8
	dirtLevel        uint8
	bodyHealth       uint32
	lightsMultiplier float32
	wheelColor       uint8
	neonColor        *entities.Rgba
	datas            *sync.Map
	*BaseObject
}

func (v *IVehicle) GetId() uint32                    { return v.id }
func (v *IVehicle) GetModel() vehicle_hash.ModelHash { return v.model }
func (v *IVehicle) GetPrimaryColor() uint8           { return v.primaryColor }
func (v *IVehicle) GetSecondColor() uint8            { return v.secondColor }
func (v *IVehicle) GetNumberplate() string           { return v.numberplate }
func (v *IVehicle) GetEngineOn() bool                { return v.engineOn }
func (v *IVehicle) GetNeonActive() bool              { return v.neonActive }
func (v *IVehicle) GetLockState() uint8              { return v.lockState }
func (v *IVehicle) GetLightState() uint8             { return v.lightState }
func (v *IVehicle) GetHeadLightColor() uint8         { return v.headLightColor }
func (v *IVehicle) GetNeonColor() *entities.Rgba     { return v.neonColor }
func (v *IVehicle) GetDimension() int32              { return v.dimension }
func (v *IVehicle) GetFrozen() bool                  { return v.frozen }
func (v *IVehicle) GetVisible() bool                 { return v.visible }
func (v *IVehicle) GetCollision() bool               { return v.collision }
func (v *IVehicle) GetDriftMode() bool               { return v.driftMode }
func (v *IVehicle) GetDisableTowing() bool           { return v.disableTowing }
func (v *IVehicle) GetPosition() *entities.Vector3 {
	ret, freeDataResultFunc := w.GetData(v.id, enum.Vehicle, uint8(enum.VehiclePosition))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.Vector3Val
	}
	return nil
}
func (v *IVehicle) GetRotation() *entities.Vector3 {
	ret, freeDataResultFunc := w.GetData(v.id, enum.Vehicle, uint8(enum.VehicleRot))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.Vector3Val
	}
	return nil
}
func (v *IVehicle) GetPositionString() string {
	position := v.GetPosition()
	return fmt.Sprintf("%v,%v,%v", position.X, position.Y, position.Z)
}
func (v *IVehicle) GetRotationString() string {
	rotation := v.GetRotation()
	return fmt.Sprintf("%v,%v,%v", rotation.X, rotation.Y, rotation.Z)
}

func (v *IVehicle) NewIVehicle(id, model uint32, primaryColor, secondColor uint8, position, rotation *entities.Vector3) *IVehicle {
	return &IVehicle{
		id:           id,
		model:        vehicle_hash.ModelHash(model),
		primaryColor: primaryColor,
		secondColor:  secondColor,
		datas:        &sync.Map{},
		BaseObject:   NewBaseObject(position, rotation, hash_enums.DefaultDimension, false, true, true),
	}
}

func (v *IVehicle) SetPrimaryColor(primaryColor uint8) {
	v.primaryColor = primaryColor
	w.SetVehicleData(v.id, enum.PrimaryColor, int64(primaryColor))
}

func (v *IVehicle) SetSecondColor(secondColor uint8) {
	v.secondColor = secondColor
	w.SetVehicleData(v.id, enum.SecondColor, int64(secondColor))
}

func (v *IVehicle) SetPosition(position *entities.Vector3) {
	v.position = position
	w.SetVehicleMetaData(v.id, enum.VehiclePosition, int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32, "", 0, 0, 0, 0)
}

func (v *IVehicle) SetRotation(rotation *entities.Vector3) {
	v.rotation = rotation
	w.SetVehicleMetaData(v.id, enum.VehicleRot, int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32, "", 0, 0, 0, 0)
}

func (v *IVehicle) SetPositionRotation(position, rotation *entities.Vector3) {
	v.position = position
	v.rotation = rotation
	w.SetVehicleMetaData(v.id, enum.VehiclePosition, int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32, "", 0, 0, 0, 0)
	w.SetVehicleMetaData(v.id, enum.VehicleRot, int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32, "", 0, 0, 0, 0)
}

func (v *IVehicle) SetVisible(visible bool) {
	v.visible = visible
	value := 0
	if visible {
		value = 1
	}
	w.SetVehicleData(v.id, enum.VehicleVisible, int64(value))
}

func (v *IVehicle) SetDimension(dimension int32) {
	v.dimension = dimension
	w.SetVehicleData(v.id, enum.VehicleDimension, int64(dimension))
}

func (v *IVehicle) SetCollision(collision bool) {
	v.collision = collision
	value := 0
	if collision {
		value = 1
	}
	w.SetVehicleData(v.id, enum.VehicleCollision, int64(value))
}

func (v *IVehicle) SetFrozen(frozen bool) {
	v.frozen = frozen
	value := 0
	if frozen {
		value = 1
	}
	w.SetVehicleData(v.id, enum.VehicleFrozen, int64(value))
}

func (v *IVehicle) SetDoorState(doorId uint8, state uint8) {
	w.SetVehicleMetaData(v.id, enum.DoorState, int64(doorId), uint64(state), "", uint8(0), uint8(0), uint8(0), uint8(0))
}

func (v *IVehicle) SetEngineOn(engineOn bool) {
	v.engineOn = engineOn
	value := 0
	if engineOn {
		value = 1
	}
	w.SetVehicleData(v.id, enum.EngineOn, int64(value))
}

func (v *IVehicle) ToggleEngine() {
	v.engineOn = !v.engineOn
	value := 0
	if v.engineOn {
		value = 1
	}
	w.SetVehicleData(v.id, enum.EngineOn, int64(value))
}

func (v *IVehicle) SetLockState(lockState uint8) {
	v.lockState = lockState
	w.SetVehicleData(v.id, enum.LockState, int64(lockState))
}

func (v *IVehicle) SetLightState(lightState uint8) {
	v.lightState = lightState
	w.SetVehicleData(v.id, enum.LightState, int64(lightState))
}

func (v *IVehicle) SetHeadLightColor(headLightColor uint8) {
	v.headLightColor = headLightColor
	w.SetVehicleData(v.id, enum.HeadLightColor, int64(headLightColor))
}

func (v *IVehicle) SetDriftMode(driftMode bool) {
	v.driftMode = driftMode
	value := 0
	if driftMode {
		value = 1
	}
	w.SetVehicleData(v.id, enum.DriftMode, int64(value))
}

func (v *IVehicle) SetDisableTowing(disableTowing bool) {
	v.disableTowing = disableTowing
	value := 0
	if disableTowing {
		value = 1
	}
	w.SetVehicleData(v.id, enum.DisableTowing, int64(value))
}

func (v *IVehicle) SetDirtLevel(dirtLevel uint8) {
	v.dirtLevel = dirtLevel
	w.SetVehicleData(v.id, enum.DirtLevel, int64(dirtLevel))
}

func (v *IVehicle) SetBodyHealth(bodyHealth uint32) {
	v.bodyHealth = bodyHealth
	w.SetVehicleData(v.id, enum.BodyHealth, int64(bodyHealth))
}

func (v *IVehicle) SetEngineHealth(engineHealth int32) {
	w.SetVehicleData(v.id, enum.EngineHealth, int64(engineHealth))
}

func (v *IVehicle) SetLightsMultiplier(lightsMultiplier float32) {
	v.lightsMultiplier = lightsMultiplier
	w.SetVehicleData(v.id, enum.LightsMultiplier, int64(math.Float64bits(float64(lightsMultiplier))))
}

func (v *IVehicle) SetWheelColor(wheelColor uint8) {
	v.wheelColor = wheelColor
	w.SetVehicleData(v.id, enum.WheelColor, int64(wheelColor))
}

func (v *IVehicle) SetMod(categray uint8, id uint8) {
	w.SetVehicleMetaData(v.id, enum.Mod, int64(categray), uint64(id), "", uint8(0), uint8(0), uint8(0), uint8(0))
}

func (v *IVehicle) SetModKit(id uint8) {
	w.SetVehicleData(v.id, enum.ModKit, int64(id))
}

func (v *IVehicle) SetRearWheels(variation uint8) {
	w.SetVehicleData(v.id, enum.RearWheels, int64(variation))
}

func (v *IVehicle) SetNeonColor(neonColor *entities.Rgba) {
	v.neonColor = neonColor
	w.SetVehicleMetaData(v.id, enum.NeonColor, int64(0), uint64(0), "", neonColor.R, neonColor.G, neonColor.B, neonColor.A)
}

func (v *IVehicle) SetNeonActive(neonActive bool) {
	v.neonActive = neonActive
	value := 0
	if neonActive {
		value = 1
	}
	w.SetVehicleMetaData(v.id, enum.NeonActive, int64(0), uint64(0), "", uint8(value), uint8(value), uint8(value), uint8(value))
}

func (v *IVehicle) SetNumberPlate(numberplate string) {
	v.numberplate = numberplate
	w.SetVehicleMetaData(v.id, enum.NumberPlate, int64(0), uint64(0), numberplate, uint8(0), uint8(0), uint8(0), uint8(0))
}

func (v *IVehicle) Repair() {
	w.SetVehicleData(v.id, enum.Repair, int64(0))
}

func (v *IVehicle) Destroy() {
	w.SetVehicleData(v.id, enum.VehicleDestroy, int64(0))
	pools.DestroyVehicle(v)
}

func (v *IVehicle) SetData(key string, value any) {
	v.datas.Store(key, value)
}

func (v *IVehicle) DelData(key string) {
	_, ok := v.datas.Load(key)
	if ok {
		v.datas.Delete(key)
	}
}

func (v *IVehicle) DelAllData() {
	v.datas.Range(func(key, value any) bool {
		v.datas.Delete(key)
		return true
	})
}

func (v *IVehicle) HasData(key string) bool {
	_, ok := v.datas.Load(key)
	if ok {
		return true
	}
	return false
}

func (v *IVehicle) GetData(key string) any {
	value, ok := v.datas.Load(key)
	if ok {
		return value
	}
	return value
}

func (v *IVehicle) GetDatas() []any {
	var datas []any
	v.datas.Range(func(key, value any) bool {
		datas = append(datas, key)
		return true
	})
	return datas
}
