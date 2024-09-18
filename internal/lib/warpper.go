package lib

// #include <stdlib.h>
import "C"
import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/hash_enums/blip_type"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"github.com/StanZzzz222/RAltGo/internal/utils"
	"github.com/StanZzzz222/RAltGo/logger"
	"os"
	"syscall"
	"time"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/5
   File: warpper.go
*/

var dllPath string
var dll *syscall.DLL
var freeProc *syscall.Proc
var mainProc *syscall.Proc
var freePlayerProc *syscall.Proc
var freeVehicleProc *syscall.Proc
var freeBlipProc *syscall.Proc
var freePedProc *syscall.Proc
var freeDataResultProc *syscall.Proc
var setVehicleDataProc *syscall.Proc
var setBlipDataProc *syscall.Proc
var setPlayerDataProc *syscall.Proc
var setPedDataProc *syscall.Proc
var createVehicleProc *syscall.Proc
var createBlipProc *syscall.Proc
var createPedProc *syscall.Proc
var getDataProc *syscall.Proc
var emitProc *syscall.Proc
var taskQueue = utils.NewTaskQueue()

type Warrper struct{}

//export onTick
func onTick() {
	if taskQueue.PopCheck() {
		task := taskQueue.Pop()
		task()
	}
}

func init() {
	path, _ := os.Getwd()
	path = fmt.Sprintf("%v/modules/rs-go-module.dll", path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		logger.LogErrorf(":: Please check if %v exists", path)
		time.Sleep(time.Second * 3)
		os.Exit(-1)
		return
	}
	dllPath = path
	dll = syscall.MustLoadDLL(dllPath)
	mainProc = dll.MustFindProc("main")
	freeProc = dll.MustFindProc("free_c_str")
	freePlayerProc = dll.MustFindProc("free_player")
	freeVehicleProc = dll.MustFindProc("free_vehicle")
	freeBlipProc = dll.MustFindProc("free_blip")
	freePedProc = dll.MustFindProc("free_ped")
	freeDataResultProc = dll.MustFindProc("free_data_result")
	setPedDataProc = dll.MustFindProc("set_ped_data")
	setPlayerDataProc = dll.MustFindProc("set_player_data")
	setVehicleDataProc = dll.MustFindProc("set_vehicle_data")
	setBlipDataProc = dll.MustFindProc("set_blip_data")
	createVehicleProc = dll.MustFindProc("create_vehicle")
	createBlipProc = dll.MustFindProc("create_blip")
	createPedProc = dll.MustFindProc("create_ped")
	getDataProc = dll.MustFindProc("get_data")
	emitProc = dll.MustFindProc("emit")
}

func (w *Warrper) ModuleMain(altVersion, core, resourceName, resourceHandlers, moduleHandlers uintptr) bool {
	ret, _, err := mainProc.Call(altVersion, core, resourceName, resourceHandlers, moduleHandlers)
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.LogErrorf("load mounted failed: %v", err.Error())
		os.Exit(-1)
	}
	return ret != 0
}

func (w *Warrper) SetPedData(id uint32, pedDataType enum.PedDataType, data int64) {
	_, _, err := setPedDataProc.Call(uintptr(id), uintptr(pedDataType), uintptr(data), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *Warrper) Emit(id uint32, eventName string, compressData string) {
	eventNamePtr, freeEventNameCStringFunc := w.GoStringMarshalPtr(eventName)
	compressDataPtr, freecompressDataCStringFunc := w.GoStringMarshalPtr(compressData)
	defer func() {
		freeEventNameCStringFunc()
		freecompressDataCStringFunc()
	}()
	_, _, err := emitProc.Call(uintptr(id), eventNamePtr, compressDataPtr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("emit failed: %v", err.Error())
		return
	}
}

func (w *Warrper) GetData(id uint32, objectType enum.ObjectType, dataType uint8) (uintptr, func()) {
	ret, _, err := getDataProc.Call(uintptr(id), uintptr(objectType), uintptr(dataType))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("get data failed: %v", err.Error())
		return 0, func() {}
	}
	freeDataResultFunc := func() {
		if ret != 0 {
			w.FreeDataResult(ret)
		}
	}
	return ret, freeDataResultFunc
}

func (w *Warrper) SetPedMetaData(id uint32, pedDataType enum.PedDataType, data int64, metaData uint64) {
	_, _, err := setPedDataProc.Call(uintptr(id), uintptr(pedDataType), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *Warrper) SetBlipData(id uint32, blipDataType enum.BlipDataType, data int64) {
	_, _, err := setBlipDataProc.Call(uintptr(id), uintptr(blipDataType), uintptr(data), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set blip data failed: %v", err.Error())
		return
	}
}

func (w *Warrper) SetBlipMetaData(id uint32, blipDataType enum.BlipDataType, data int64, metaData uint64, strData string, r, g, b, a uint8) {
	var freeCStringFunc func()
	var strPtr = uintptr(0)
	if len(strData) > 0 {
		strPtr, freeCStringFunc = w.GoStringMarshalPtr(strData)
		defer freeCStringFunc()
	}
	_, _, err := setBlipDataProc.Call(uintptr(id), uintptr(blipDataType), uintptr(data), uintptr(metaData), strPtr, uintptr(r), uintptr(g), uintptr(b), uintptr(a))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set blip data failed: %v", err.Error())
		return
	}
}

func (w *Warrper) SetVehicleData(id uint32, vehicleDataType enum.VehicleDataType, data int64) {
	_, _, err := setVehicleDataProc.Call(uintptr(id), uintptr(vehicleDataType), uintptr(data), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set vehicle_hash data failed: %v", err.Error())
		return
	}
}

func (w *Warrper) SetVehicleMetaData(id uint32, vehicleDataType enum.VehicleDataType, data int64, metaData uint64, strData string, l, r, t, b uint8) {
	var freeCStringFunc func()
	var strPtr = uintptr(0)
	if len(strData) > 0 {
		strPtr, freeCStringFunc = w.GoStringMarshalPtr(strData)
		defer freeCStringFunc()
	}
	_, _, err := setVehicleDataProc.Call(uintptr(id), uintptr(vehicleDataType), uintptr(data), uintptr(metaData), strPtr, uintptr(l), uintptr(r), uintptr(t), uintptr(b))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set vehicle_hash data failed: %v", err.Error())
		return
	}
}

func (w *Warrper) SetPlayerMetaData(id uint32, playerDataType enum.PlayerDataType, data int64, metaData uint64) {
	_, _, err := setPlayerDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(0), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *Warrper) SetPlayerMetaModelData(id uint32, playerDataType enum.PlayerDataType, model uint32, data int64, metaData uint64) {
	_, _, err := setPlayerDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(model), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *Warrper) SetPlayerData(id uint32, playerDataType enum.PlayerDataType, data int64) {
	_, _, err := setPlayerDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(0), uintptr(data), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *Warrper) CreateVehicle(model uint32, posData, posMetaData, rotData, rotMetaData uint64, numberplate uintptr, primaryColor, secondColor uint8) (uintptr, func()) {
	ret, _, err := createVehicleProc.Call(uintptr(model), uintptr(posData), uintptr(posMetaData), uintptr(rotData), uintptr(rotMetaData), numberplate, uintptr(primaryColor), uintptr(secondColor))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.LogErrorf("create vehicle_hash failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeVehicle(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *Warrper) CreatePed(model uint32, posData, posMetaData, rotData, rotMetaData uint64, streamingDistance uint32) (uintptr, func()) {
	ret, _, err := createPedProc.Call(uintptr(model), uintptr(posData), uintptr(posMetaData), uintptr(rotData), uintptr(rotMetaData), uintptr(streamingDistance))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.LogErrorf("create vehicle_hash failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreePed(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *Warrper) CreateBlip(blipType blip_type.BlipType, spriteId, color uint32, strData string, posData, posMetaData uint64, width, height, radius float32) (uintptr, func()) {
	var freeCStringFunc func()
	var strPtr = uintptr(0)
	if len(strData) > 0 {
		strPtr, freeCStringFunc = w.GoStringMarshalPtr(strData)
		defer freeCStringFunc()
	}
	ret, _, err := createBlipProc.Call(uintptr(blipType), uintptr(spriteId), uintptr(color), strPtr, uintptr(posData), uintptr(posMetaData), uintptr(width), uintptr(height), uintptr(radius))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.LogErrorf("create blip failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeBlip(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *Warrper) PushTask(callback func()) {
	taskQueue.AddTask(callback)
}

func (w *Warrper) Free(ptr uintptr) {
	_, _, err := freeProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("free failed: %v", err.Error())
		return
	}
}

func (w *Warrper) FreePlayer(ptr uintptr) {
	_, _, err := freePlayerProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("free player failed: %v", err.Error())
		return
	}
}

func (w *Warrper) FreeVehicle(ptr uintptr) {
	_, _, err := freeVehicleProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("free vehicle_hash failed: %v", err.Error())
		return
	}
}

func (w *Warrper) FreeBlip(ptr uintptr) {
	_, _, err := freeBlipProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("free blip failed: %v", err.Error())
		return
	}
}

func (w *Warrper) FreePed(ptr uintptr) {
	_, _, err := freePedProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("free ped_hash failed: %v", err.Error())
		return
	}
}

func (w *Warrper) FreeDataResult(ptr uintptr) {
	_, _, err := freeDataResultProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("free data result failed: %v", err.Error())
		return
	}
}

func (w *Warrper) GoStringMarshalPtr(s string) (uintptr, func()) {
	cStr := C.CString(s)
	return uintptr(unsafe.Pointer(cStr)), func() { C.free(unsafe.Pointer(cStr)) }
}

func (w *Warrper) PtrMarshalGoString(ret uintptr) string {
	cStr := (*C.char)(unsafe.Pointer(ret))
	defer w.Free(ret)
	return C.GoString(cStr)
}
