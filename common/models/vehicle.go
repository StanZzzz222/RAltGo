package models

import (
	"github.com/StanZzzz222/RAltGo/enums"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/5
   File: vehicle.go
*/

type IVehicle struct {
	id               uint32
	model            uint32
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
	engineHealth     int32
	lightsMultiplier float32
	wheelColor       uint8
	neonColor        *entities.Rgba
	*BaseObject
}

func (v *IVehicle) GetId() uint32                { return v.id }
func (v *IVehicle) GetModel() uint32             { return v.model }
func (v *IVehicle) GetPrimaryColor() uint8       { return v.primaryColor }
func (v *IVehicle) GetSecondColor() uint8        { return v.secondColor }
func (v *IVehicle) GetNumberplate() string       { return v.numberplate }
func (v *IVehicle) GetEngineOn() bool            { return v.engineOn }
func (v *IVehicle) GetNeonActive() bool          { return v.neonActive }
func (v *IVehicle) GetLockState() uint8          { return v.lockState }
func (v *IVehicle) GetLightState() uint8         { return v.lightState }
func (v *IVehicle) GetHeadLightColor() uint8     { return v.headLightColor }
func (v *IVehicle) GetNeonColor() *entities.Rgba { return v.neonColor }
func (v *IVehicle) GetDimension() int32          { return v.dimension }
func (v *IVehicle) GetFrozen() bool              { return v.frozen }
func (v *IVehicle) GetVisible() bool             { return v.visible }
func (v *IVehicle) GetCollision() bool           { return v.collision }
func (v *IVehicle) GetDriftMode() bool           { return v.driftMode }
func (v *IVehicle) GetDisableTowing() bool       { return v.disableTowing }

func (v *IVehicle) NewIVehicle(id, model uint32, primaryColor, secondColor uint8, position, rotation *entities.Vector3) *IVehicle {
	return &IVehicle{
		id:           id,
		model:        model,
		primaryColor: primaryColor,
		secondColor:  secondColor,
		BaseObject:   NewBaseObject(position, rotation, enums.DefaultDimension, false, true, true),
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
	v.engineHealth = engineHealth
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
