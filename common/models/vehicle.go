package models

import (
	"github.com/StanZzzz222/RAltGo/common/utils"
	"github.com/StanZzzz222/RAltGo/enums"
	"github.com/StanZzzz222/RAltGo/internal/entitys"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/5
   File: vehicle.go
*/

type IVehicle struct {
	id             uint32
	model          uint32
	primaryColor   uint8
	secondColor    uint8
	numberplate    string
	engineOn       bool
	neonActive     bool
	lockState      uint8
	lightState     uint8
	headLightColor uint8
	neonColor      *utils.Rgba
	*BaseObject
}

func (v *IVehicle) GetId() uint32             { return v.id }
func (v *IVehicle) GetModel() uint32          { return v.model }
func (v *IVehicle) GetPrimaryColor() uint8    { return v.primaryColor }
func (v *IVehicle) GetSecondColor() uint8     { return v.secondColor }
func (v *IVehicle) GetNumberplate() string    { return v.numberplate }
func (v *IVehicle) GetEngineOn() bool         { return v.engineOn }
func (v *IVehicle) GetNeonActive() bool       { return v.neonActive }
func (v *IVehicle) GetLockState() uint8       { return v.lockState }
func (v *IVehicle) GetLightState() uint8      { return v.lightState }
func (v *IVehicle) GetHeadLightColor() uint8  { return v.headLightColor }
func (v *IVehicle) GetNeonColor() *utils.Rgba { return v.neonColor }
func (v *IVehicle) GetDimension() int32       { return v.dimension }
func (v *IVehicle) GetFrozen() bool           { return v.frozen }
func (v *IVehicle) GetVisible() bool          { return v.visible }
func (v *IVehicle) GetCollision() bool        { return v.collision }

func (v *IVehicle) NewIVehicle(id, model uint32, primaryColor, secondColor uint8, position, rotation *entitys.Vector3) *IVehicle {
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
	w.SetVehicleData(v.id, enum.PrimaryColor, uint64(primaryColor))
}

func (v *IVehicle) SetSecondColor(secondColor uint8) {
	v.secondColor = secondColor
	w.SetVehicleData(v.id, enum.SecondColor, uint64(secondColor))
}

func (v *IVehicle) SetPosition(position *entitys.Vector3) {
	v.position = position
	w.SetVehicleMetaData(v.id, enum.VehiclePosition, uint64(int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32)), uint64(math.Float32bits(position.Z))<<32, "", 0, 0, 0, 0)
}

func (v *IVehicle) SetRotation(rotation *entitys.Vector3) {
	v.rotation = rotation
	w.SetVehicleMetaData(v.id, enum.VehicleRot, uint64(int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32)), uint64(math.Float32bits(rotation.Z))<<32, "", 0, 0, 0, 0)
}

func (v *IVehicle) SetPositionRotation(position, rotation *entitys.Vector3) {
	v.position = position
	v.rotation = rotation
	w.SetVehicleMetaData(v.id, enum.VehiclePosition, uint64(int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32)), uint64(math.Float32bits(position.Z))<<32, "", 0, 0, 0, 0)
	w.SetVehicleMetaData(v.id, enum.VehicleRot, uint64(int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32)), uint64(math.Float32bits(rotation.Z))<<32, "", 0, 0, 0, 0)
}

func (v *IVehicle) SetVisible(visible bool) {
	v.visible = visible
	value := 0
	if visible {
		value = 1
	}
	w.SetVehicleData(v.id, enum.VehicleVisible, uint64(value))
}

func (v *IVehicle) SetDimension(dimension int32) {
	v.dimension = dimension
	w.SetVehicleData(v.id, enum.VehicleDimension, uint64(dimension))
}

func (v *IVehicle) SetCollision(collision bool) {
	v.collision = collision
	value := 0
	if collision {
		value = 1
	}
	w.SetVehicleData(v.id, enum.VehicleCollision, uint64(value))
}

func (v *IVehicle) SetFrozen(frozen bool) {
	v.frozen = frozen
	value := 0
	if frozen {
		value = 1
	}
	w.SetVehicleData(v.id, enum.VehicleFrozen, uint64(value))
}

func (v *IVehicle) SetEngineOn(engineOn bool) {
	v.engineOn = engineOn
	value := 0
	if engineOn {
		value = 1
	}
	w.SetVehicleData(v.id, enum.EngineOn, uint64(value))
}

func (v *IVehicle) ToggleEngine() {
	v.engineOn = !v.engineOn
	value := 0
	if v.engineOn {
		value = 1
	}
	w.SetVehicleData(v.id, enum.EngineOn, uint64(value))
}

func (v *IVehicle) SetLockState(lockState uint8) {
	v.lockState = lockState
	w.SetVehicleData(v.id, enum.LockState, uint64(lockState))
}

func (v *IVehicle) SetLightState(lightState uint8) {
	v.lightState = lightState
	w.SetVehicleData(v.id, enum.LightState, uint64(lightState))
}

func (v *IVehicle) SetHeadLightColor(headLightColor uint8) {
	v.headLightColor = headLightColor
	w.SetVehicleData(v.id, enum.HeadLightColor, uint64(headLightColor))
}

func (v *IVehicle) SetNeonColor(neonColor *utils.Rgba) {
	v.neonColor = neonColor
	w.SetVehicleMetaData(v.id, enum.NeonColor, uint64(0), uint64(0), "", neonColor.R, neonColor.G, neonColor.B, neonColor.A)
}

func (v *IVehicle) SetNeonActive(neonActive bool) {
	v.neonActive = neonActive
	value := 0
	if neonActive {
		value = 1
	}
	w.SetVehicleData(v.id, enum.NeonActive, uint64(value))
}

func (v *IVehicle) SetNumberPlate(numberplate string) {
	v.numberplate = numberplate
	w.SetVehicleMetaData(v.id, enum.NumberPlate, uint64(0), uint64(0), numberplate, uint8(0), uint8(0), uint8(0), uint8(0))
}
