//go:build !windows

package lib

// #include <stdlib.h>
import "C"
import (
	"github.com/StanZzzz222/RAltGo/hash_enums/blip_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_type"
	"github.com/StanZzzz222/RAltGo/internal/enums"
	"github.com/StanZzzz222/RAltGo/internal/utils"
	"github.com/StanZzzz222/RAltGo/logger"
	"runtime"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: lib_windows.go
*/

type Wrapper struct {
	syscall *syscall_warpper.SyscallWrapper
}

var wrapper *Wrapper
var taskQueue = utils.NewTaskQueue()

//export onTick
func onTick() {
	if taskQueue.PopCheck() {
		defer panicRecover()
		taskQueue.Pop()()
	}
}

func init() {
	wrapper = &Wrapper{
		syscall: &syscall_warpper.SyscallWrapper{},
	}
}

func GetWrapper() *Wrapper {
	return wrapper
}

func (w *Wrapper) ModuleMain(altVersion, core, resourceName, resourceHandlers, moduleHandlers uintptr) bool {
	return w.syscall.ModuleMain(altVersion, core, resourceName, resourceHandlers, moduleHandlers)
}

func (w *Wrapper) SetPedData(id uint32, pedDataType enums.PedDataType, data int64) {
	w.syscall.SetPedData(id, pedDataType, data)
}

func (w *Wrapper) SetVirtualEntityData(id uint32, virtualEntityDataType enums.VirtualEntityDataType, data int64, metaData uint64) {
	w.syscall.SetVirtualEntityData(id, virtualEntityDataType, data, metaData)
}

func (w *Wrapper) SetColshapeData(id uint32, colshapeDataType enums.ColshapeDataType, data int64, metaData uint64) {
	w.syscall.SetColshapeData(id, colshapeDataType, data, metaData)
}

func (w *Wrapper) SetCheckpointData(id uint32, checkpointDataType enums.CheckpointDataType, data int64, metaData uint64, otherData float32, r, g, b, a uint8) {
	w.syscall.SetCheckpointData(id, checkpointDataType, data, metaData, otherData, r, g, b, a)
}

func (w *Wrapper) SetMarkerData(id uint32, markerDataType enums.MarkerDataType, data int64, metaData uint64, r, g, b, a uint8) {
	w.syscall.SetMarkerData(id, markerDataType, data, metaData, r, g, b, a)
}

func (w *Wrapper) SetObjectData(id uint32, objectDataType enums.ObjectDataType, data int64, metaData uint64) {
	w.syscall.SetObjectData(id, objectDataType, data, metaData)
}

func (w *Wrapper) Emit(id uint32, eventName, data string) {
	w.syscall.Emit(id, eventName, data)
}

func (w *Wrapper) EmitAllPlayer(eventName, data string) {
	w.syscall.EmitAllPlayer(eventName, data)
}

func (w *Wrapper) OnClientEvent(eventName string, eventArgsDump string) {
	w.syscall.OnClientEvent(eventName, eventArgsDump)
}

func (w *Wrapper) SetServerData(setType enums.ServerDataType, data int64, strData string) (uintptr, func()) {
	return w.syscall.SetServerData(int32(setType), data, strData)
}

func (w *Wrapper) GetServerData(getType enums.ServerDataType, data uint32) (uintptr, func()) {
	return w.syscall.GetServerData(int32(getType), data)
}

func (w *Wrapper) GetColshapeData(id uint32, dataType enums.ColshapeDataType, entityType enums.ObjectType, data int64, metaData uint64) (uintptr, func()) {
	return w.syscall.GetColshapeData(id, enums.Colshape, dataType, entityType, data, metaData)
}

func (w *Wrapper) GetData(id uint32, objectType enums.ObjectType, dataType uint8) (uintptr, func()) {
	return w.syscall.GetData(id, objectType, dataType)
}

func (w *Wrapper) GetMetaData(id uint32, objectType enums.ObjectType, dataType uint8, metaData int64) (uintptr, func()) {
	return w.syscall.GetMetaData(id, objectType, dataType, metaData)
}

func (w *Wrapper) SetPedMetaData(id uint32, pedDataType enums.PedDataType, data int64, metaData uint64) {
	w.syscall.SetPedMetaData(id, pedDataType, data, metaData)
}

func (w *Wrapper) SetBlipData(id uint32, blipDataType enums.BlipDataType, data int64) {
	w.syscall.SetBlipData(id, blipDataType, data)
}

func (w *Wrapper) SetBlipMetaData(id uint32, blipDataType enums.BlipDataType, data int64, metaData uint64, strData string, r, g, b, a uint8) {
	w.syscall.SetBlipMetaData(id, blipDataType, data, metaData, strData, r, g, b, a)
}

func (w *Wrapper) SetVehicleData(id uint32, vehicleDataType enums.VehicleDataType, data int64) {
	w.syscall.SetVehicleData(id, vehicleDataType, data)
}

func (w *Wrapper) SetVehicleMetaData(id uint32, vehicleDataType enums.VehicleDataType, data int64, metaData uint64, strData string, l, r, t, b uint8) {
	w.syscall.SetVehicleMetaData(id, vehicleDataType, data, metaData, strData, l, r, t, b)
}

func (w *Wrapper) SetPlayerMetaData(id uint32, playerDataType enums.PlayerDataType, data int64, metaData uint64) {
	w.syscall.SetPlayerMetaData(id, playerDataType, data, metaData)
}

func (w *Wrapper) SetPlayerMetaModelData(id uint32, playerDataType enums.PlayerDataType, model uint32, data int64, metaData uint64) {
	w.syscall.SetPlayerMetaModelData(id, playerDataType, model, data, metaData)
}

func (w *Wrapper) SetPlayerHeadData(id uint32, playerDataType enums.PlayerDataType, shape1, shape2, shape3, skin1, skin2, skin3 uint32, shapeMix, skinMix, thirdMix float32) {
	w.syscall.SetPlayerHeadData(id, playerDataType, shape1, shape2, shape3, skin1, skin2, skin3, shapeMix, skinMix, thirdMix)
}

func (w *Wrapper) SetPlayerData(id uint32, playerDataType enums.PlayerDataType, data int64) {
	w.syscall.SetPlayerData(id, playerDataType, data)
}

func (w *Wrapper) GetEntityData(id uint32, dataType enums.ObjectType, networkDataType enums.EntityDataType) (uintptr, func()) {
	ret, freeEntityDataFunc := w.syscall.GetEntityData(id, uint8(dataType), uint8(networkDataType))
	return ret, freeEntityDataFunc
}

func (w *Wrapper) SetEntityData(id uint32, dataType enums.ObjectType, entityDataType enums.EntityDataType, entityType enums.ObjectType, data uint64, metaData uint32, attachData string) {
	w.syscall.SetEntityData(id, uint8(dataType), uint8(entityDataType), uint8(entityType), data, metaData, attachData)
}

func (w *Wrapper) SetNetworkData(id uint32, dataType enums.ObjectType, networkDataType enums.NetworkDataType, keysData, valuesData string) {
	w.syscall.SetNetworkData(id, uint8(dataType), uint8(networkDataType), keysData, valuesData)
}

func (w *Wrapper) SetVoiceChannelData(id uint32, voiceChannelDataType enums.VoiceChannelDataType, data int64) {
	w.syscall.SetVoiceChannelData(id, uint8(voiceChannelDataType), data)
}

func (w *Wrapper) CreateVirtualEntityGroup(maxEntitiesInStream uint32) (uintptr, func()) {
	return w.syscall.CreateVirtualEntityGroup(maxEntitiesInStream)
}

func (w *Wrapper) CreateVoiceChannel(spatial bool, maxDistance float32) (uintptr, func()) {
	var value uint8 = 0
	if spatial {
		value = 1
	}
	return w.syscall.CreateVoiceChannel(value, maxDistance)
}

func (w *Wrapper) CreateVirtualEntity(groupId uint32, posData, posMetaData uint64, streamingDistance uint32) (uintptr, func()) {
	return w.syscall.CreateVirtualEntity(groupId, posData, posMetaData, streamingDistance)
}

func (w *Wrapper) CreateCheckpoint(checkPointType uint8, posData, posMetaData uint64, radius, height float32, r, g, b, a uint8, streamingDistance uint32) (uintptr, func()) {
	return w.syscall.CreateCheckpoint(checkPointType, posData, posMetaData, radius, height, r, g, b, a, streamingDistance)
}

func (w *Wrapper) CreateMarker(markerType uint8, posData, posMetaData uint64, r, g, b, a uint8) (uintptr, func()) {
	return w.syscall.CreateMarker(markerType, posData, posMetaData, r, g, b, a)
}

func (w *Wrapper) CreateObject(model uint32, posData, posMetaData, rotData, rotMetaData uint64) (uintptr, func()) {
	return w.syscall.CreateObject(model, posData, posMetaData, rotData, rotMetaData)
}

func (w *Wrapper) CreateVehicle(model uint32, posData, posMetaData, rotData, rotMetaData uint64, numberplate uintptr, primaryColor, secondColor uint8) (uintptr, func()) {
	return w.syscall.CreateVehicle(model, posData, posMetaData, rotData, rotMetaData, numberplate, primaryColor, secondColor)
}

func (w *Wrapper) CreatePed(model uint32, posData, posMetaData, rotData, rotMetaData uint64, streamingDistance uint32) (uintptr, func()) {
	return w.syscall.CreatePed(model, posData, posMetaData, rotData, rotMetaData, streamingDistance)
}

func (w *Wrapper) CreateBlip(blipType blip_type.BlipType, spriteId, color uint32, strData string, posData, posMetaData uint64, width, height, radius float32) (uintptr, func()) {
	return w.syscall.CreateBlip(blipType, spriteId, color, strData, posData, posMetaData, width, height, radius)
}

func (w *Wrapper) CreateColshape(colshapeType colshape_type.ColshapeType, posData, posMetaData, secondPosData, secondPosMetaData uint64, radius, height float32) (uintptr, func()) {
	return w.syscall.CreateColshape(colshapeType, posData, posMetaData, secondPosData, secondPosMetaData, radius, height)
}

func (w *Wrapper) CreatePolygonColshape(colshapeType colshape_type.ColshapeType, minZ, maxZ float32, pointsData []byte) (uintptr, func()) {
	return w.syscall.CreatePolygonColshape(colshapeType, minZ, maxZ, pointsData)
}

func (w *Wrapper) PushTask(callback func()) {
	taskQueue.AddTask(callback)
}

func (w *Wrapper) Free(ptr uintptr) {
	w.syscall.Free(ptr)
}

func (w *Wrapper) FreePlayer(ptr uintptr) {
	w.syscall.FreePlayer(ptr)
}

func (w *Wrapper) FreeVehicle(ptr uintptr) {
	w.syscall.FreeVehicle(ptr)
}

func (w *Wrapper) FreeBlip(ptr uintptr) {
	w.syscall.FreeBlip(ptr)
}

func (w *Wrapper) FreePed(ptr uintptr) {
	w.syscall.FreePed(ptr)
}

func (w *Wrapper) FreeColshape(ptr uintptr) {
	w.syscall.FreeColshape(ptr)
}

func (w *Wrapper) FreeDataResult(ptr uintptr) {
	w.syscall.FreeDataResult(ptr)
}

func (w *Wrapper) GoStringMarshalPtr(s string) (uintptr, func()) {
	cStr := C.CString(s)
	return uintptr(unsafe.Pointer(cStr)), func() { C.free(unsafe.Pointer(cStr)) }
}

func (w *Wrapper) PtrMarshalGoString(ret uintptr) string {
	cStr := (*C.char)(unsafe.Pointer(ret))
	defer w.syscall.Free(ret)
	return C.GoString(cStr)
}

func panicRecover() {
	if r := recover(); r != nil {
		var stackBuf [4096]byte
		stackSize := runtime.Stack(stackBuf[:], false)
		logger.Logger().LogErrorf("Panic recovered: %v", r)
		logger.Logger().LogErrorf("StackTrace: %s", stackBuf[:stackSize])
	}
}
