package lib

// #include <stdlib.h>
import "C"
import (
	"github.com/StanZzzz222/RAltGo/hash_enums/blip_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_type"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"github.com/StanZzzz222/RAltGo/internal/lib/syscall_warpper"
	"github.com/StanZzzz222/RAltGo/internal/lib/windows_warpper"
	"github.com/StanZzzz222/RAltGo/internal/utils"
	"github.com/StanZzzz222/RAltGo/logger"
	"runtime"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: lib.go
*/

var warpper *Warpper
var taskQueue = utils.NewTaskQueue()

type Warpper struct {
	windows *windows_warpper.WindowsWarrper
	syscall *syscall_warpper.SyscallWarrper
}

//export onTick
func onTick() {
	defer panicRecover()
	if taskQueue.PopCheck() {
		task := taskQueue.Pop()
		task()
	}
}

func init() {
	if runtime.GOOS == "windows" {
		warpper = &Warpper{
			windows: &windows_warpper.WindowsWarrper{},
			syscall: nil,
		}
		return
	}
	warpper = &Warpper{
		windows: nil,
		syscall: &syscall_warpper.SyscallWarrper{},
	}
}

func GetWarpper() *Warpper {
	return warpper
}

func (w *Warpper) IsWindows() bool {
	return w.windows != nil
}

func (w *Warpper) ModuleMain(altVersion, core, resourceName, resourceHandlers, moduleHandlers uintptr) bool {
	if w.IsWindows() {
		return w.windows.ModuleMain(altVersion, core, resourceName, resourceHandlers, moduleHandlers)
	}
	return w.syscall.ModuleMain(altVersion, core, resourceName, resourceHandlers, moduleHandlers)
}

func (w *Warpper) SetPedData(id uint32, pedDataType enum.PedDataType, data int64) {
	if w.IsWindows() {
		w.windows.SetPedData(id, pedDataType, data)
		return
	}
	w.syscall.SetPedData(id, pedDataType, data)
}

func (w *Warpper) SetColshapeData(id uint32, colshapeDataType enum.ColshapeDataType, data int64, metaData uint64) {
	if w.IsWindows() {
		w.windows.SetColshapeData(id, colshapeDataType, data, metaData)
		return
	}
	w.syscall.SetColshapeData(id, colshapeDataType, data, metaData)
}

func (w *Warpper) Emit(id uint32, eventName, data string) {
	if w.IsWindows() {
		w.windows.Emit(id, eventName, data)
		return
	}
	w.syscall.Emit(id, eventName, data)
}

func (w *Warpper) EmitAllPlayer(eventName, data string) {
	if w.IsWindows() {
		w.windows.EmitAllPlayer(eventName, data)
		return
	}
	w.syscall.EmitAllPlayer(eventName, data)
}

func (w *Warpper) OnClientEvent(eventName string, eventArgsDump string) {
	if w.IsWindows() {
		w.windows.OnClientEvent(eventName, eventArgsDump)
		return
	}
	w.syscall.OnClientEvent(eventName, eventArgsDump)
}

func (w *Warpper) GetData(id uint32, objectType enum.ObjectType, dataType uint8) (uintptr, func()) {
	if w.IsWindows() {
		return w.windows.GetData(id, objectType, dataType)
	}
	return w.syscall.GetData(id, objectType, dataType)
}

func (w *Warpper) GetMetaData(id uint32, objectType enum.ObjectType, dataType uint8, metaData int64) (uintptr, func()) {
	if w.IsWindows() {
		return w.windows.GetMetaData(id, objectType, dataType, metaData)
	}
	return w.syscall.GetMetaData(id, objectType, dataType, metaData)
}

func (w *Warpper) SetPedMetaData(id uint32, pedDataType enum.PedDataType, data int64, metaData uint64) {
	if w.IsWindows() {
		w.windows.SetPedMetaData(id, pedDataType, data, metaData)
		return
	}
	w.syscall.SetPedMetaData(id, pedDataType, data, metaData)
}

func (w *Warpper) SetBlipData(id uint32, blipDataType enum.BlipDataType, data int64) {
	if w.IsWindows() {
		w.windows.SetBlipData(id, blipDataType, data)
		return
	}
	w.syscall.SetBlipData(id, blipDataType, data)
}

func (w *Warpper) SetBlipMetaData(id uint32, blipDataType enum.BlipDataType, data int64, metaData uint64, strData string, r, g, b, a uint8) {
	if w.IsWindows() {
		w.windows.SetBlipMetaData(id, blipDataType, data, metaData, strData, r, g, b, a)
		return
	}
	w.syscall.SetBlipMetaData(id, blipDataType, data, metaData, strData, r, g, b, a)
}

func (w *Warpper) SetVehicleData(id uint32, vehicleDataType enum.VehicleDataType, data int64) {
	if w.IsWindows() {
		w.windows.SetVehicleData(id, vehicleDataType, data)
		return
	}
	w.syscall.SetVehicleData(id, vehicleDataType, data)
}

func (w *Warpper) SetVehicleMetaData(id uint32, vehicleDataType enum.VehicleDataType, data int64, metaData uint64, strData string, l, r, t, b uint8) {
	if w.IsWindows() {
		w.windows.SetVehicleMetaData(id, vehicleDataType, data, metaData, strData, l, r, t, b)
		return
	}
	w.syscall.SetVehicleMetaData(id, vehicleDataType, data, metaData, strData, l, r, t, b)
}

func (w *Warpper) SetPlayerMetaData(id uint32, playerDataType enum.PlayerDataType, data int64, metaData uint64) {
	if w.IsWindows() {
		w.windows.SetPlayerMetaData(id, playerDataType, data, metaData)
		return
	}
	w.syscall.SetPlayerMetaData(id, playerDataType, data, metaData)
}

func (w *Warpper) SetPlayerMetaModelData(id uint32, playerDataType enum.PlayerDataType, model uint32, data int64, metaData uint64) {
	if w.IsWindows() {
		w.windows.SetPlayerMetaModelData(id, playerDataType, model, data, metaData)
		return
	}
	w.syscall.SetPlayerMetaModelData(id, playerDataType, model, data, metaData)
}

func (w *Warpper) SetPlayerHeadData(id uint32, playerDataType enum.PlayerDataType, shape1, shape2, shape3, skin1, skin2, skin3 uint32, shapeMix, skinMix, thirdMix float32) {
	if w.IsWindows() {
		w.windows.SetPlayerHeadData(id, playerDataType, shape1, shape2, shape3, skin1, skin2, skin3, shapeMix, skinMix, thirdMix)
		return
	}
	w.syscall.SetPlayerHeadData(id, playerDataType, shape1, shape2, shape3, skin1, skin2, skin3, shapeMix, skinMix, thirdMix)
}

func (w *Warpper) SetPlayerData(id uint32, playerDataType enum.PlayerDataType, data int64) {
	if w.IsWindows() {
		w.windows.SetPlayerData(id, playerDataType, data)
		return
	}
	w.syscall.SetPlayerData(id, playerDataType, data)
}

func (w *Warpper) CreateVehicle(model uint32, posData, posMetaData, rotData, rotMetaData uint64, numberplate uintptr, primaryColor, secondColor uint8) (uintptr, func()) {
	if w.IsWindows() {
		return w.windows.CreateVehicle(model, posData, posMetaData, rotData, rotMetaData, numberplate, primaryColor, secondColor)
	}
	return w.syscall.CreateVehicle(model, posData, posMetaData, rotData, rotMetaData, numberplate, primaryColor, secondColor)
}

func (w *Warpper) CreatePed(model uint32, posData, posMetaData, rotData, rotMetaData uint64, streamingDistance uint32) (uintptr, func()) {
	if w.IsWindows() {
		return w.windows.CreatePed(model, posData, posMetaData, rotData, rotMetaData, streamingDistance)
	}
	return w.syscall.CreatePed(model, posData, posMetaData, rotData, rotMetaData, streamingDistance)
}

func (w *Warpper) CreateBlip(blipType blip_type.BlipType, spriteId, color uint32, strData string, posData, posMetaData uint64, width, height, radius float32) (uintptr, func()) {
	if w.IsWindows() {
		return w.windows.CreateBlip(blipType, spriteId, color, strData, posData, posMetaData, width, height, radius)
	}
	return w.syscall.CreateBlip(blipType, spriteId, color, strData, posData, posMetaData, width, height, radius)
}

func (w *Warpper) CreateColshape(colshapeType colshape_type.ColshapeType, posData, posMetaData, secondPosData, secondPosMetaData uint64, radius, height float32) (uintptr, func()) {
	if w.IsWindows() {
		return w.windows.CreateColshape(colshapeType, posData, posMetaData, secondPosData, secondPosMetaData, radius, height)
	}
	return w.syscall.CreateColshape(colshapeType, posData, posMetaData, secondPosData, secondPosMetaData, radius, height)
}

func (w *Warpper) CreatePolygonColshape(colshapeType colshape_type.ColshapeType, minZ, maxZ float32, pointsData []byte) (uintptr, func()) {
	if w.IsWindows() {
		return w.windows.CreatePolygonColshape(colshapeType, minZ, maxZ, pointsData)
	}
	return w.syscall.CreatePolygonColshape(colshapeType, minZ, maxZ, pointsData)
}

func (w *Warpper) PushTask(callback func()) {
	taskQueue.AddTask(callback)
}

func (w *Warpper) Free(ptr uintptr) {
	if w.IsWindows() {
		w.windows.Free(ptr)
		return
	}
	w.syscall.Free(ptr)
}

func (w *Warpper) FreePlayer(ptr uintptr) {
	if w.IsWindows() {
		w.windows.FreePlayer(ptr)
		return
	}
	w.syscall.FreePlayer(ptr)
}

func (w *Warpper) FreeVehicle(ptr uintptr) {
	if w.IsWindows() {
		w.windows.FreeVehicle(ptr)
		return
	}
	w.syscall.FreeVehicle(ptr)
}

func (w *Warpper) FreeBlip(ptr uintptr) {
	if w.IsWindows() {
		w.windows.FreeBlip(ptr)
		return
	}
	w.syscall.FreeBlip(ptr)
}

func (w *Warpper) FreePed(ptr uintptr) {
	if w.IsWindows() {
		w.windows.FreePed(ptr)
		return
	}
	w.syscall.FreePed(ptr)
}

func (w *Warpper) FreeColshape(ptr uintptr) {
	if w.IsWindows() {
		w.windows.FreeColshape(ptr)
		return
	}
	w.syscall.FreeColshape(ptr)
}

func (w *Warpper) FreeDataResult(ptr uintptr) {
	if w.IsWindows() {
		w.windows.FreeDataResult(ptr)
		return
	}
	w.syscall.FreeDataResult(ptr)
}

func (w *Warpper) GoStringMarshalPtr(s string) (uintptr, func()) {
	cStr := C.CString(s)
	return uintptr(unsafe.Pointer(cStr)), func() { C.free(unsafe.Pointer(cStr)) }
}

func (w *Warpper) PtrMarshalGoString(ret uintptr) string {
	cStr := (*C.char)(unsafe.Pointer(ret))
	if w.IsWindows() {
		defer w.windows.Free(ret)
	} else {
		defer w.syscall.Free(ret)
	}
	return C.GoString(cStr)
}

func panicRecover() {
	if r := recover(); r != nil {
		var stackBuf [4096]byte
		stackSize := runtime.Stack(stackBuf[:], false)
		logger.LogErrorf("Panic recovered: %v", r)
		logger.LogErrorf("StackTrace: %s", stackBuf[:stackSize])
	}
}
