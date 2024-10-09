package models

import (
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"github.com/StanZzzz222/RAltGo/hash_enums/radio_station_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_door_state"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_door_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_hash"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_head_light_color_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_light_id_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_light_state_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_lock_state_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/vehicle_mod_type"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enums"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/5
   File: vehicle_hash.go
*/

type IVehicle struct {
	id                  uint32
	model               vehicle_hash.ModelHash
	primaryColor        uint8
	secondColor         uint8
	interiorColor       uint8
	windowTint          uint8
	numberplate         string
	engineOn            bool
	neonActive          bool
	driftMode           bool
	disableTowing       bool
	boatAnchorActive    bool
	customTires         bool
	hybridExtraActive   bool
	hybridExtraState    uint8
	numberplateIndex    uint32
	specialDarkness     uint8
	manualEngineControl bool
	scriptMaxSpeed      int
	lockState           vehicle_lock_state_type.VehicleLockState
	lightState          vehicle_light_state_type.VehicleLightState
	headLightColor      vehicle_head_light_color_type.VehicleHeadLightColorType
	dirtLevel           uint8
	bodyHealth          uint32
	lightsMultiplier    float32
	wheelColor          uint8
	tireSmokeColor      *entities.Rgba
	primaryColorRgb     *entities.Rgba
	secondColorRgb      *entities.Rgba
	neonColor           *entities.Rgba
	*BaseObject
	*NetworkData
	*EntityData
}

func (v *IVehicle) NewIVehicle(id, model uint32, primaryColor, secondColor uint8, position, rotation *entities.Vector3) *IVehicle {
	return &IVehicle{
		id:               id,
		model:            vehicle_hash.ModelHash(model),
		primaryColor:     primaryColor,
		secondColor:      secondColor,
		headLightColor:   vehicle_head_light_color_type.Default,
		bodyHealth:       1000,
		lightsMultiplier: 1,
		wheelColor:       0,
		interiorColor:    0,
		lockState:        vehicle_lock_state_type.VehicleLockNone,
		lightState:       vehicle_light_state_type.VehicleLightOff,
		neonColor:        &entities.Rgba{R: 0, G: 0, B: 0, A: 0},
		BaseObject:       NewBaseObject(position, rotation, hash_enums.DefaultDimension, false, true, true),
		NetworkData:      NewNetworkData(id, enums.Vehicle),
		EntityData:       NewEntityData(id, enums.Vehicle),
	}
}

func (v *IVehicle) GetId() uint32                      { return v.id }
func (v *IVehicle) GetModel() vehicle_hash.ModelHash   { return v.model }
func (v *IVehicle) GetPrimaryColor() uint8             { return v.primaryColor }
func (v *IVehicle) GetSecondColor() uint8              { return v.secondColor }
func (v *IVehicle) GetInteriorColor() uint8            { return v.interiorColor }
func (v *IVehicle) GetWindowTint() uint8               { return v.windowTint }
func (v *IVehicle) GetNumberplate() string             { return v.numberplate }
func (v *IVehicle) GetNeonActive() bool                { return v.neonActive }
func (v *IVehicle) GetNeonColor() *entities.Rgba       { return v.neonColor }
func (v *IVehicle) GetDimension() int32                { return v.dimension }
func (v *IVehicle) GetFrozen() bool                    { return v.frozen }
func (v *IVehicle) GetBoatAnchorActive() bool          { return v.boatAnchorActive }
func (v *IVehicle) GetCustomTires() bool               { return v.customTires }
func (v *IVehicle) GetVisible() bool                   { return v.visible }
func (v *IVehicle) GetCollision() bool                 { return v.collision }
func (v *IVehicle) GetDriftMode() bool                 { return v.driftMode }
func (v *IVehicle) GetDisableTowing() bool             { return v.disableTowing }
func (v *IVehicle) GetHybridExtraActive() bool         { return v.hybridExtraActive }
func (v *IVehicle) GetHybridExtraState() uint8         { return v.hybridExtraState }
func (v *IVehicle) GetManualEngineControl() bool       { return v.manualEngineControl }
func (v *IVehicle) GetNumberplateIndex() uint32        { return v.numberplateIndex }
func (v *IVehicle) GetScriptMaxSpeed() int             { return v.scriptMaxSpeed }
func (v *IVehicle) GetSpecialDarkness() uint8          { return v.specialDarkness }
func (v *IVehicle) GetTireSmokeColor() *entities.Rgba  { return v.tireSmokeColor }
func (v *IVehicle) GetPrimaryRgbColor() *entities.Rgba { return v.primaryColorRgb }
func (v *IVehicle) GetSecondRgbColor() *entities.Rgba  { return v.secondColorRgb }
func (v *IVehicle) GetHeadLightColor() vehicle_head_light_color_type.VehicleHeadLightColorType {
	return v.headLightColor
}
func (v *IVehicle) GetBodyAdditionalHealth() bool {
	ret, freeDataResultFunc := w.GetData(v.id, enums.Vehicle, uint8(enums.VehicleBodyAdditionalHealth))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (v *IVehicle) GetSirenActive() bool {
	ret, freeDataResultFunc := w.GetData(v.id, enums.Vehicle, uint8(enums.VehicleSirenActive))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (v *IVehicle) DoesWheelHasTire(wheelId uint8) bool {
	ret, freeDataResultFunc := w.GetMetaData(v.id, enums.Vehicle, uint8(enums.VehicleWheelHasTire), int64(wheelId))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (v *IVehicle) IsWheelOnFire(wheelId uint8) bool {
	ret, freeDataResultFunc := w.GetMetaData(v.id, enums.Vehicle, uint8(enums.VehicleWheelOnFire), int64(wheelId))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (v *IVehicle) IsWheelBurst(wheelId uint8) bool {
	ret, freeDataResultFunc := w.GetMetaData(v.id, enums.Vehicle, uint8(enums.VehicleWheelBurst), int64(wheelId))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (v *IVehicle) IsWheelDetached(wheelId uint8) bool {
	ret, freeDataResultFunc := w.GetMetaData(v.id, enums.Vehicle, uint8(enums.VehicleWheelDetached), int64(wheelId))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (v *IVehicle) GetWeaponCapacity(index uint8) bool {
	ret, freeDataResultFunc := w.GetMetaData(v.id, enums.Vehicle, uint8(enums.VehicleWeaponCapacity), int64(index))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (v *IVehicle) IsWindowOpened(windowId uint8) bool {
	ret, freeDataResultFunc := w.GetMetaData(v.id, enums.Vehicle, uint8(enums.VehicleWindowOpened), int64(windowId))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (v *IVehicle) IsWindowDamaged(windowId uint8) bool {
	ret, freeDataResultFunc := w.GetMetaData(v.id, enums.Vehicle, uint8(enums.VehicleWindowDamaged), int64(windowId))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (v *IVehicle) IsSpecialLightDamaged(lightId uint8) bool {
	ret, freeDataResultFunc := w.GetMetaData(v.id, enums.Vehicle, uint8(enums.VehicleSpecialLightDamaged), int64(lightId))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (v *IVehicle) GetRoofClosed() bool {
	ret, freeDataResultFunc := w.GetData(v.id, enums.Vehicle, uint8(enums.VehicleRoofClosed))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (v *IVehicle) GetRoofLivery() uint8 {
	ret, freeDataResultFunc := w.GetData(v.id, enums.Vehicle, uint8(enums.VehicleRoofLivery))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U8Val
	}
	return 0
}
func (v *IVehicle) GetLivery() uint8 {
	ret, freeDataResultFunc := w.GetData(v.id, enums.Vehicle, uint8(enums.VehicleLivery))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U8Val
	}
	return 0
}
func (v *IVehicle) GetPosition() *entities.Vector3 {
	ret, freeDataResultFunc := w.GetData(v.id, enums.Vehicle, uint8(enums.VehiclePosition))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.Vector3Val
	}
	return nil
}
func (v *IVehicle) GetRotation() *entities.Vector3 {
	ret, freeDataResultFunc := w.GetData(v.id, enums.Vehicle, uint8(enums.VehicleRot))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.Vector3Val
	}
	return nil
}
func (v *IVehicle) GetLightState() vehicle_light_state_type.VehicleLightState {
	ret, freeDataResultFunc := w.GetData(v.id, enums.Vehicle, uint8(enums.VehicleLightState))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return vehicle_light_state_type.VehicleLightState(cDataResult.U8Val)
	}
	return vehicle_light_state_type.VehicleLightNormal
}
func (v *IVehicle) GetBodyHealth() uint32 {
	ret, freeDataResultFunc := w.GetData(v.id, enums.Vehicle, uint8(enums.VehicleBodyHealth))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U32Val
	}
	return 1000
}
func (v *IVehicle) GetRadioStation() uint32 {
	ret, freeDataResultFunc := w.GetData(v.id, enums.Vehicle, uint8(enums.VehicleRadioStation))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U32Val
	}
	return 0
}
func (v *IVehicle) GetDashboardColor() uint8 {
	ret, freeDataResultFunc := w.GetData(v.id, enums.Vehicle, uint8(enums.VehicleDashboardColor))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U8Val
	}
	return 0
}
func (v *IVehicle) IsLightDamaged(lightId vehicle_light_id_type.VehicleLightType) bool {
	ret, freeDataResultFunc := w.GetMetaData(v.id, enums.Vehicle, uint8(enums.VehicleLightDamaged), int64(lightId))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (v *IVehicle) GetDriver() *IPlayer {
	ret, freeDataResultFunc := w.GetMetaData(v.id, enums.Vehicle, uint8(enums.VehicleDriver), int64(0))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return pools.GetPlayer(cDataResult.U32Val)
	}
	return nil
}
func (v *IVehicle) GetEngineOn() bool {
	ret, freeDataResultFunc := w.GetMetaData(v.id, enums.Vehicle, uint8(enums.VehicleEngineOn), int64(0))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.BoolVal
	}
	return false
}
func (v *IVehicle) GetLockState() vehicle_lock_state_type.VehicleLockState {
	ret, freeDataResultFunc := w.GetMetaData(v.id, enums.Vehicle, uint8(enums.VehicleLockState), int64(0))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return vehicle_lock_state_type.VehicleLockState(cDataResult.U8Val)
	}
	return vehicle_lock_state_type.VehicleLockNone
}

func (v *IVehicle) SetPrimaryColor(primaryColor uint8) {
	v.primaryColor = primaryColor
	w.SetVehicleData(v.id, enums.VehiclePrimaryColor, int64(primaryColor))
}

func (v *IVehicle) SetSecondColor(secondColor uint8) {
	v.secondColor = secondColor
	w.SetVehicleData(v.id, enums.VehicleSecondColor, int64(secondColor))
}

func (v *IVehicle) SetPosition(position *entities.Vector3) {
	v.position = position
	w.SetVehicleMetaData(v.id, enums.VehiclePosition, int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32, "", 0, 0, 0, 0)
}

func (v *IVehicle) SetRotation(rotation *entities.Vector3) {
	v.rotation = rotation
	w.SetVehicleMetaData(v.id, enums.VehicleRot, int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32, "", 0, 0, 0, 0)
}

func (v *IVehicle) SetPositionRotation(position, rotation *entities.Vector3) {
	v.position = position
	v.rotation = rotation
	w.SetVehicleMetaData(v.id, enums.VehiclePosition, int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32, "", 0, 0, 0, 0)
	w.SetVehicleMetaData(v.id, enums.VehicleRot, int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32, "", 0, 0, 0, 0)
}

func (v *IVehicle) SetVisible(visible bool) {
	v.visible = visible
	value := 0
	if visible {
		value = 1
	}
	w.SetVehicleData(v.id, enums.VehicleVisible, int64(value))
}

func (v *IVehicle) SetDimension(dimension int32) {
	v.dimension = dimension
	w.SetVehicleData(v.id, enums.VehicleDimension, int64(dimension))
}

func (v *IVehicle) SetCollision(collision bool) {
	v.collision = collision
	value := 0
	if collision {
		value = 1
	}
	w.SetVehicleData(v.id, enums.VehicleCollision, int64(value))
}

func (v *IVehicle) SetFrozen(frozen bool) {
	v.frozen = frozen
	value := 0
	if frozen {
		value = 1
	}
	w.SetVehicleData(v.id, enums.VehicleFrozen, int64(value))
}

func (v *IVehicle) SetDoorState(door vehicle_door_type.VehicleDoorType, state vehicle_door_state.VehicleDoorState) {
	w.SetVehicleMetaData(v.id, enums.VehicleDoorState, int64(door), uint64(state), "", uint8(0), uint8(0), uint8(0), uint8(0))
}

func (v *IVehicle) SetEngineOn(engineOn bool) {
	v.engineOn = engineOn
	value := 0
	if engineOn {
		value = 1
	}
	w.SetVehicleData(v.id, enums.VehicleEngineOn, int64(value))
}

func (v *IVehicle) ToggleEngine() {
	v.engineOn = !v.engineOn
	value := 0
	if v.engineOn {
		value = 1
	}
	w.SetVehicleData(v.id, enums.VehicleEngineOn, int64(value))
}

func (v *IVehicle) SetLockState(lockState vehicle_lock_state_type.VehicleLockState) {
	v.lockState = lockState
	w.SetVehicleData(v.id, enums.VehicleLockState, int64(lockState))
}

func (v *IVehicle) SetLightState(lightState vehicle_light_state_type.VehicleLightState) {
	v.lightState = lightState
	w.SetVehicleData(v.id, enums.VehicleLightState, int64(lightState))
}

func (v *IVehicle) SetHeadLightColor(headLightColor vehicle_head_light_color_type.VehicleHeadLightColorType) {
	v.headLightColor = headLightColor
	v.SetModKit(1)
	v.SetMod(vehicle_mod_type.Xenon, 1)
	w.SetVehicleData(v.id, enums.VehicleHeadLightColor, int64(headLightColor))
}

func (v *IVehicle) SetDriftMode(driftMode bool) {
	v.driftMode = driftMode
	value := 0
	if driftMode {
		value = 1
	}
	w.SetVehicleData(v.id, enums.VehicleDriftMode, int64(value))
}

func (v *IVehicle) SetDisableTowing(disableTowing bool) {
	v.disableTowing = disableTowing
	value := 0
	if disableTowing {
		value = 1
	}
	w.SetVehicleData(v.id, enums.VehicleDisableTowing, int64(value))
}

func (v *IVehicle) SetDirtLevel(dirtLevel uint8) {
	if dirtLevel < 0 {
		dirtLevel = 0
	}
	if dirtLevel > 15 {
		dirtLevel = 15
	}
	v.dirtLevel = dirtLevel
	w.SetVehicleData(v.id, enums.VehicleDirtLevel, int64(dirtLevel))
}

func (v *IVehicle) SetBodyHealth(bodyHealth uint32) {
	if bodyHealth > 1000 {
		bodyHealth = 1000
	}
	if bodyHealth < 0 {
		bodyHealth = 0
	}
	v.bodyHealth = bodyHealth
	w.SetVehicleData(v.id, enums.VehicleBodyHealth, int64(bodyHealth))
}

func (v *IVehicle) SetEngineHealth(engineHealth int32) {
	w.SetVehicleData(v.id, enums.VehicleEngineHealth, int64(engineHealth))
}

func (v *IVehicle) SetLightsMultiplier(lightsMultiplier float32) {
	v.lightsMultiplier = lightsMultiplier
	w.SetVehicleData(v.id, enums.VehicleLightsMultiplier, int64(math.Float64bits(float64(lightsMultiplier))))
}

func (v *IVehicle) SetWheelColor(wheelColor uint8) {
	v.wheelColor = wheelColor
	w.SetVehicleData(v.id, enums.VehicleWheelColor, int64(wheelColor))
}

func (v *IVehicle) SetMod(modType vehicle_mod_type.VehicleModType, id uint8) {
	w.SetVehicleMetaData(v.id, enums.VehicleMod, int64(modType), uint64(id), "", uint8(0), uint8(0), uint8(0), uint8(0))
}

func (v *IVehicle) SetModKit(id uint8) {
	w.SetVehicleData(v.id, enums.VehicleModKit, int64(id))
}

func (v *IVehicle) SetRearWheels(variation uint8) {
	w.SetVehicleData(v.id, enums.VehicleRearWheels, int64(variation))
}

func (v *IVehicle) SetNeonColor(neonColor *entities.Rgba) {
	v.neonColor = neonColor
	w.SetVehicleMetaData(v.id, enums.VehicleNeonColor, int64(0), uint64(0), "", neonColor.R, neonColor.G, neonColor.B, neonColor.A)
}

func (v *IVehicle) SetNeonActive(neonActive bool) {
	v.neonActive = neonActive
	value := 0
	if neonActive {
		value = 1
	}
	w.SetVehicleMetaData(v.id, enums.VehicleNeonActive, int64(0), uint64(0), "", uint8(value), uint8(value), uint8(value), uint8(value))
}

func (v *IVehicle) SetNumberPlate(numberplate string) {
	v.numberplate = numberplate
	w.SetVehicleMetaData(v.id, enums.VehicleNumberPlate, int64(0), uint64(0), numberplate, uint8(0), uint8(0), uint8(0), uint8(0))
}

func (v *IVehicle) SetInteriorColor(color uint8) {
	v.interiorColor = color
	w.SetVehicleData(v.id, enums.VehicleInteriorColor, int64(color))
}

func (v *IVehicle) SetBoatAnchorActive(boatAnchorActive bool) {
	v.boatAnchorActive = boatAnchorActive
	value := 0
	if boatAnchorActive {
		value = 1
	}
	w.SetVehicleData(v.id, enums.VehicleBoatAnchorActive, int64(value))
}

func (v *IVehicle) SetCustomTires(customTires bool) {
	v.customTires = customTires
	value := 0
	if customTires {
		value = 1
	}
	w.SetVehicleData(v.id, enums.VehicleCustomTires, int64(value))
}

func (v *IVehicle) SetLightDamaged(lightId vehicle_light_id_type.VehicleLightType, damaged bool) {
	value := 0
	if damaged {
		value = 1
	}
	w.SetVehicleMetaData(v.id, enums.VehicleLightDamaged, int64(lightId), uint64(value), "", uint8(0), uint8(0), uint8(0), uint8(0))
}

func (v *IVehicle) SetRadioStation(radioStation radio_station_type.RadioStation) {
	w.SetVehicleData(v.id, enums.VehicleRadioStation, int64(radioStation))
}

func (v *IVehicle) SetDashboardColor(color uint8) {
	w.SetVehicleData(v.id, enums.VehicleDashboardColor, int64(color))
}

func (v *IVehicle) SetWindowTint(windowTint uint8) {
	v.windowTint = windowTint
	w.SetVehicleData(v.id, enums.VehicleWindowTint, int64(windowTint))
}

func (v *IVehicle) Repair() {
	w.SetVehicleData(v.id, enums.VehicleRepair, int64(0))
}

func (v *IVehicle) SetHybridExtraActive(hybridExtraActive bool) {
	v.hybridExtraActive = hybridExtraActive
	value := 0
	if hybridExtraActive {
		value = 1
	}
	w.SetVehicleData(v.id, enums.VehicleHybridExtraActive, int64(value))
}

func (v *IVehicle) SetHybridExtraState(state uint8) {
	v.hybridExtraState = state
	w.SetVehicleData(v.id, enums.VehicleHybridExtraState, int64(state))
}

func (v *IVehicle) SetManualEngineControl(manualEngineControl bool) {
	v.manualEngineControl = manualEngineControl
	value := 0
	if manualEngineControl {
		value = 1
	}
	w.SetVehicleData(v.id, enums.VehicleManualEngineControl, int64(value))
}

func (v *IVehicle) SetLivery(livery uint8) {
	w.SetVehicleData(v.id, enums.VehicleLivery, int64(livery))
}

func (v *IVehicle) SetNumberplateIndex(index uint32) {
	v.numberplateIndex = index
	w.SetVehicleData(v.id, enums.VehicleNumberplateIndex, int64(index))
}

func (v *IVehicle) SetRoofClosed(roofClosed bool) {
	value := 0
	if roofClosed {
		value = 1
	}
	w.SetVehicleData(v.id, enums.VehicleRoofClosed, int64(value))
}

func (v *IVehicle) SetRoofLivery(roofLivery uint8) {
	w.SetVehicleData(v.id, enums.VehicleRoofLivery, int64(roofLivery))
}

func (v *IVehicle) SetScriptMaxSpeed(speed int) {
	v.scriptMaxSpeed = speed
	w.SetVehicleData(v.id, enums.VehicleScriptMaxSpeed, int64(speed))
}

func (v *IVehicle) SetSirenActive(sirenActive bool) {
	value := 0
	if sirenActive {
		value = 1
	}
	w.SetVehicleData(v.id, enums.VehicleSirenActive, int64(value))
}

func (v *IVehicle) SetSpecialDarkness(specialDarkness uint8) {
	v.specialDarkness = specialDarkness
	w.SetVehicleData(v.id, enums.VehicleSpecialDarkness, int64(specialDarkness))
}

func (v *IVehicle) SetSpecialLightDamaged(specialLightId uint8, specialDarkness bool) {
	value := 0
	if specialDarkness {
		value = 1
	}
	w.SetVehicleMetaData(v.id, enums.VehicleSpecialLightDamaged, int64(specialLightId), uint64(value), "", uint8(0), uint8(0), uint8(0), uint8(0))
}

func (v *IVehicle) SetTireSmokeColor(tireSmokeColor *entities.Rgba) {
	v.tireSmokeColor = tireSmokeColor
	w.SetVehicleMetaData(v.id, enums.VehicleTireSmokeColor, int64(0), uint64(0), "", tireSmokeColor.R, tireSmokeColor.G, tireSmokeColor.B, tireSmokeColor.A)
}

func (v *IVehicle) SetWeaponCapacity(index uint8, state int32) {
	w.SetVehicleMetaData(v.id, enums.VehicleWeaponCapacity, int64(index), uint64(state), "", 0, 0, 0, 0)
}

func (v *IVehicle) SetWheelBurst(wheelId uint8, state bool) {
	value := 0
	if state {
		value = 1
	}
	w.SetVehicleMetaData(v.id, enums.VehicleWheelBurst, int64(wheelId), uint64(value), "", 0, 0, 0, 0)
}

func (v *IVehicle) SetWheelDetached(wheelId uint8, state bool) {
	value := 0
	if state {
		value = 1
	}
	w.SetVehicleMetaData(v.id, enums.VehicleWheelDetached, int64(wheelId), uint64(value), "", 0, 0, 0, 0)
}

func (v *IVehicle) SetWheelFixed(wheelId uint8) {
	w.SetVehicleData(v.id, enums.VehicleWheelFixed, int64(wheelId))
}

func (v *IVehicle) SetWheelHasTire(wheelId uint8, state bool) {
	value := 0
	if state {
		value = 1
	}
	w.SetVehicleMetaData(v.id, enums.VehicleWheelHasTire, int64(wheelId), uint64(value), "", 0, 0, 0, 0)
}

func (v *IVehicle) SetWheelOnFire(wheelId uint8, state bool) {
	value := 0
	if state {
		value = 1
	}
	w.SetVehicleMetaData(v.id, enums.VehicleWheelOnFire, int64(wheelId), uint64(value), "", 0, 0, 0, 0)
}

func (v *IVehicle) SetWindowOpened(windowId uint8, opened bool) {
	value := 0
	if opened {
		value = 1
	}
	w.SetVehicleMetaData(v.id, enums.VehicleWindowOpened, int64(windowId), uint64(value), "", 0, 0, 0, 0)
}

func (v *IVehicle) SetWindowDamaged(windowId uint8, damaged bool) {
	value := 0
	if damaged {
		value = 1
	}
	w.SetVehicleMetaData(v.id, enums.VehicleWindowOpened, int64(windowId), uint64(value), "", 0, 0, 0, 0)
}

func (v *IVehicle) SetPrimaryColorRgb(color *entities.Rgba) {
	v.primaryColorRgb = color
	w.SetVehicleMetaData(v.id, enums.VehiclePrimaryColorRgb, int64(0), uint64(0), "", color.R, color.G, color.B, color.A)
}

func (v *IVehicle) SetSecondColorRgb(color *entities.Rgba) {
	v.secondColorRgb = color
	w.SetVehicleMetaData(v.id, enums.VehicleSecondColorRgb, int64(0), uint64(0), "", color.R, color.G, color.B, color.A)
}

func (v *IVehicle) SetBodyAdditionalHealth(health uint32) {
	w.SetVehicleData(v.id, enums.VehicleBodyAdditionalHealth, int64(health))
}

func (v *IVehicle) SetWheels(wheelType, variation uint8) {
	w.SetVehicleMetaData(v.id, enums.VehicleWheels, int64(wheelType), uint64(variation), "", 0, 0, 0, 0)
}

func (v *IVehicle) Destroy() {
	w.SetVehicleData(v.id, enums.VehicleDestroy, int64(0))
	pools.DestroyVehicle(v)
}
