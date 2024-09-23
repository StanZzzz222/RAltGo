package windows_warpper

// #include <stdlib.h>
import "C"
import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/hash_enums/blip_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_type"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"github.com/StanZzzz222/RAltGo/logger"
	"golang.org/x/sys/windows"
	"math"
	"os"
	"runtime"
	"time"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/5
   File: warpper.go
*/

var dllPath string
var dll *windows.DLL
var freeProc *windows.Proc
var mainProc *windows.Proc
var freePlayerProc *windows.Proc
var freeVehicleProc *windows.Proc
var freeBlipProc *windows.Proc
var freePedProc *windows.Proc
var freeColshapeProc *windows.Proc
var freeDataResultProc *windows.Proc
var setColshapeData *windows.Proc
var setVehicleDataProc *windows.Proc
var setBlipDataProc *windows.Proc
var setPlayerDataProc *windows.Proc
var setPayerHeadDataProc *windows.Proc
var setPedDataProc *windows.Proc
var createVehicleProc *windows.Proc
var createBlipProc *windows.Proc
var createPedProc *windows.Proc
var createColshapeProc *windows.Proc
var getDataProc *windows.Proc
var emitProc *windows.Proc
var emitAllPlayerProc *windows.Proc
var onClientEventProc *windows.Proc

type WindowsWarrper struct{}

func init() {
	if runtime.GOOS == "windows" {
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
		dll = windows.MustLoadDLL(dllPath)
		mainProc = dll.MustFindProc("main")
		freeProc = dll.MustFindProc("free_c_str")
		freePlayerProc = dll.MustFindProc("free_player")
		freeVehicleProc = dll.MustFindProc("free_vehicle")
		freeBlipProc = dll.MustFindProc("free_blip")
		freePedProc = dll.MustFindProc("free_ped")
		freeColshapeProc = dll.MustFindProc("free_colshape")
		freeDataResultProc = dll.MustFindProc("free_data_result")
		setPedDataProc = dll.MustFindProc("set_ped_data")
		setPlayerDataProc = dll.MustFindProc("set_player_data")
		setPayerHeadDataProc = dll.MustFindProc("set_player_head_data")
		setVehicleDataProc = dll.MustFindProc("set_vehicle_data")
		setBlipDataProc = dll.MustFindProc("set_blip_data")
		setColshapeData = dll.MustFindProc("set_colshape_data")
		createVehicleProc = dll.MustFindProc("create_vehicle")
		createBlipProc = dll.MustFindProc("create_blip")
		createPedProc = dll.MustFindProc("create_ped")
		createColshapeProc = dll.MustFindProc("create_colshape")
		getDataProc = dll.MustFindProc("get_data")
		emitProc = dll.MustFindProc("emit")
		emitAllPlayerProc = dll.MustFindProc("emit_all")
		onClientEventProc = dll.MustFindProc("on_client_event")
	}
}

func (w *WindowsWarrper) ModuleMain(altVersion, core, resourceName, resourceHandlers, moduleHandlers uintptr) bool {
	ret, _, err := mainProc.Call(altVersion, core, resourceName, resourceHandlers, moduleHandlers)
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.LogErrorf("load mounted failed: %v", err.Error())
		os.Exit(-1)
	}
	return ret != 0
}

func (w *WindowsWarrper) SetPedData(id uint32, pedDataType enum.PedDataType, data int64) {
	_, _, err := setPedDataProc.Call(uintptr(id), uintptr(pedDataType), uintptr(data), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) SetColshapeData(id uint32, colshapeDataType enum.ColshapeDataType, data int64, metaData uint64) {
	_, _, err := setColshapeData.Call(uintptr(id), uintptr(colshapeDataType), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set colshape data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) Emit(id uint32, eventName, data string) {
	eventNamePtr, freeEventNameCStringFunc := w.GoStringMarshalPtr(eventName)
	dataPtr, freeDataCStringFunc := w.GoStringMarshalPtr(data)
	defer func() {
		freeEventNameCStringFunc()
		freeDataCStringFunc()
	}()
	_, _, err := emitProc.Call(uintptr(id), eventNamePtr, dataPtr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("emit failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) EmitAllPlayer(eventName, data string) {
	eventNamePtr, freeEventNameCStringFunc := w.GoStringMarshalPtr(eventName)
	dataPtr, freeDataCStringFunc := w.GoStringMarshalPtr(data)
	defer func() {
		freeEventNameCStringFunc()
		freeDataCStringFunc()
	}()
	_, _, err := emitAllPlayerProc.Call(eventNamePtr, dataPtr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("emit all failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) OnClientEvent(eventName string, eventArgsDump string) {
	eventNamePtr, freeEventNameCStringFunc := w.GoStringMarshalPtr(eventName)
	argsDumpDataPtr, freeArgsDumpDataCStringFunc := w.GoStringMarshalPtr(eventArgsDump)
	defer func() {
		freeEventNameCStringFunc()
		freeArgsDumpDataCStringFunc()
	}()
	_, _, err := onClientEventProc.Call(eventNamePtr, argsDumpDataPtr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("on client event failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) GetData(id uint32, objectType enum.ObjectType, dataType uint8) (uintptr, func()) {
	ret, _, err := getDataProc.Call(uintptr(id), uintptr(objectType), uintptr(dataType), uintptr(0))
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

func (w *WindowsWarrper) GetMetaData(id uint32, objectType enum.ObjectType, dataType uint8, metaData int64) (uintptr, func()) {
	ret, _, err := getDataProc.Call(uintptr(id), uintptr(objectType), uintptr(dataType), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("get meta data failed: %v", err.Error())
		return 0, func() {}
	}
	freeDataResultFunc := func() {
		if ret != 0 {
			w.FreeDataResult(ret)
		}
	}
	return ret, freeDataResultFunc
}

func (w *WindowsWarrper) SetPedMetaData(id uint32, pedDataType enum.PedDataType, data int64, metaData uint64) {
	_, _, err := setPedDataProc.Call(uintptr(id), uintptr(pedDataType), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) SetBlipData(id uint32, blipDataType enum.BlipDataType, data int64) {
	_, _, err := setBlipDataProc.Call(uintptr(id), uintptr(blipDataType), uintptr(data), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set blip data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) SetBlipMetaData(id uint32, blipDataType enum.BlipDataType, data int64, metaData uint64, strData string, r, g, b, a uint8) {
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

func (w *WindowsWarrper) SetVehicleData(id uint32, vehicleDataType enum.VehicleDataType, data int64) {
	_, _, err := setVehicleDataProc.Call(uintptr(id), uintptr(vehicleDataType), uintptr(data), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set vehicle_hash data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) SetVehicleMetaData(id uint32, vehicleDataType enum.VehicleDataType, data int64, metaData uint64, strData string, l, r, t, b uint8) {
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

func (w *WindowsWarrper) SetPlayerMetaData(id uint32, playerDataType enum.PlayerDataType, data int64, metaData uint64) {
	_, _, err := setPlayerDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(0), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) SetPlayerMetaModelData(id uint32, playerDataType enum.PlayerDataType, model uint32, data int64, metaData uint64) {
	_, _, err := setPlayerDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(model), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) SetPlayerData(id uint32, playerDataType enum.PlayerDataType, data int64) {
	_, _, err := setPlayerDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(0), uintptr(data), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) SetPlayerHeadData(id uint32, playerDataType enum.PlayerDataType, shape1, shape2, shape3, skin1, skin2, skin3 uint32, shapeMix, skinMix, thirdMix float32) {
	_, _, err := setPayerHeadDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(shape1), uintptr(shape2), uintptr(shape3), uintptr(skin1), uintptr(skin2), uintptr(skin3), uintptr(math.Float32bits(shapeMix)), uintptr(math.Float32bits(skinMix)), uintptr(math.Float32bits(thirdMix)))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("set player head data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) CreateVehicle(model uint32, posData, posMetaData, rotData, rotMetaData uint64, numberplate uintptr, primaryColor, secondColor uint8) (uintptr, func()) {
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

func (w *WindowsWarrper) CreatePed(model uint32, posData, posMetaData, rotData, rotMetaData uint64, streamingDistance uint32) (uintptr, func()) {
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

func (w *WindowsWarrper) CreateBlip(blipType blip_type.BlipType, spriteId, color uint32, strData string, posData, posMetaData uint64, width, height, radius float32) (uintptr, func()) {
	var freeCStringFunc func()
	var strPtr = uintptr(0)
	if len(strData) > 0 {
		strPtr, freeCStringFunc = w.GoStringMarshalPtr(strData)
		defer freeCStringFunc()
	}
	ret, _, err := createBlipProc.Call(uintptr(blipType), uintptr(spriteId), uintptr(color), strPtr, uintptr(posData), uintptr(posMetaData), uintptr(width), uintptr(math.Float32bits(height)), uintptr(math.Float32bits(radius)))
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

func (w *WindowsWarrper) CreateColshape(colshapeType colshape_type.ColshapeType, posData, posMetaData, secondPosData, secondPosMetaData uint64, radius, height float32) (uintptr, func()) {
	ret, _, err := createColshapeProc.Call(uintptr(colshapeType), uintptr(posData), uintptr(posMetaData), uintptr(secondPosData), uintptr(secondPosMetaData), uintptr(math.Float32bits(radius)), uintptr(math.Float32bits(height)))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.LogErrorf("create colshape failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeColshape(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *WindowsWarrper) CreatePolygonColshape(colshapeType colshape_type.ColshapeType, minZ, maxZ float32, pointsData []byte) (uintptr, func()) {
	var freeCStringFunc func()
	var strPtr = uintptr(0)
	var strData = string(pointsData)
	if len(strData) > 0 {
		strPtr, freeCStringFunc = w.GoStringMarshalPtr(strData)
		defer freeCStringFunc()
	}
	ret, _, err := createColshapeProc.Call(uintptr(colshapeType), uintptr(math.Float32bits(minZ)), uintptr(math.Float32bits(maxZ)), strPtr)
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.LogErrorf("create polygon colshape failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeColshape(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *WindowsWarrper) Free(ptr uintptr) {
	_, _, err := freeProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("free failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) FreePlayer(ptr uintptr) {
	_, _, err := freePlayerProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("free player failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) FreeVehicle(ptr uintptr) {
	_, _, err := freeVehicleProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("free vehicle_hash failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) FreeBlip(ptr uintptr) {
	_, _, err := freeBlipProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("free blip failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) FreePed(ptr uintptr) {
	_, _, err := freePedProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("free ped failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) FreeColshape(ptr uintptr) {
	_, _, err := freeColshapeProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("free colshape failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) FreeDataResult(ptr uintptr) {
	_, _, err := freeDataResultProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.LogErrorf("free data result failed: %v", err.Error())
		return
	}
}

func (w *WindowsWarrper) GoStringMarshalPtr(s string) (uintptr, func()) {
	cStr := C.CString(s)
	return uintptr(unsafe.Pointer(cStr)), func() { C.free(unsafe.Pointer(cStr)) }
}

func (w *WindowsWarrper) PtrMarshalGoString(ret uintptr) string {
	cStr := (*C.char)(unsafe.Pointer(ret))
	defer w.Free(ret)
	return C.GoString(cStr)
}
