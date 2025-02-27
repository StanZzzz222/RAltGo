//go:build windows

package windows_wrapper

// #include <stdlib.h>
import "C"
import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/hash_enums/blip_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_type"
	"github.com/StanZzzz222/RAltGo/internal/enums"
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
   File: wrapper.go
*/

var dllPath string
var dll *windows.DLL
var emitProc *windows.Proc
var emitAllPlayerProc *windows.Proc
var onClientEventProc *windows.Proc
var freeProc *windows.Proc
var mainProc *windows.Proc
var freePlayerProc *windows.Proc
var freeVehicleProc *windows.Proc
var freeBlipProc *windows.Proc
var freePedProc *windows.Proc
var freeColshapeProc *windows.Proc
var freeCheckpointProc *windows.Proc
var freeMarkerProc *windows.Proc
var freeObjectProc *windows.Proc
var freeVirtualEntityGroupProc *windows.Proc
var freeVirtualEntityProc *windows.Proc
var freeVoiceChannelProc *windows.Proc
var freeDataResultProc *windows.Proc
var setVirtualEntityDataProc *windows.Proc
var setColshapeDataProc *windows.Proc
var setCheckpointDataProc *windows.Proc
var setMarkerDataProc *windows.Proc
var setObjectDataProc *windows.Proc
var setVehicleDataProc *windows.Proc
var setNetworkDataProc *windows.Proc
var setBlipDataProc *windows.Proc
var setPlayerDataProc *windows.Proc
var setPlayerHeadDataProc *windows.Proc
var setPedDataProc *windows.Proc
var setServerDataProc *windows.Proc
var setEntityDataProc *windows.Proc
var setVoiceChannelDataProc *windows.Proc
var createVirtualEntityGroupProc *windows.Proc
var createVirtualEntityProc *windows.Proc
var createObjectProc *windows.Proc
var createMarkerProc *windows.Proc
var createCheckpointProc *windows.Proc
var createVehicleProc *windows.Proc
var createBlipProc *windows.Proc
var createPedProc *windows.Proc
var createColshapeProc *windows.Proc
var createPolygonColshapeProc *windows.Proc
var createVoiceChannelProc *windows.Proc
var getEntityDataProc *windows.Proc
var getServerDataProc *windows.Proc
var getColshapeDataProc *windows.Proc
var getDataProc *windows.Proc

type WindowsWrapper struct{}

func init() {
	if runtime.GOOS == "windows" {
		path, _ := os.Getwd()
		path = fmt.Sprintf("%v/modules/rs-go-module.dll", path)
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			logger.Logger().LogErrorf(":: Please check if %v exists", path)
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
		freeCheckpointProc = dll.MustFindProc("free_checkpoint")
		freeMarkerProc = dll.MustFindProc("free_marker")
		freeObjectProc = dll.MustFindProc("free_object")
		freeVirtualEntityGroupProc = dll.MustFindProc("free_virtual_entity_group")
		freeVirtualEntityProc = dll.MustFindProc("free_virtual_entity")
		freeVoiceChannelProc = dll.MustFindProc("free_voice_channel")
		freeDataResultProc = dll.MustFindProc("free_data_result")
		setNetworkDataProc = dll.MustFindProc("set_network_data")
		setVirtualEntityDataProc = dll.MustFindProc("set_virtual_entity_data")
		setPedDataProc = dll.MustFindProc("set_ped_data")
		setPlayerDataProc = dll.MustFindProc("set_player_data")
		setPlayerHeadDataProc = dll.MustFindProc("set_player_head_data")
		setVehicleDataProc = dll.MustFindProc("set_vehicle_data")
		setBlipDataProc = dll.MustFindProc("set_blip_data")
		setColshapeDataProc = dll.MustFindProc("set_colshape_data")
		setCheckpointDataProc = dll.MustFindProc("set_checkpoint_data")
		setMarkerDataProc = dll.MustFindProc("set_marker_data")
		setObjectDataProc = dll.MustFindProc("set_object_data")
		setVoiceChannelDataProc = dll.MustFindProc("set_voice_channel_data")
		createMarkerProc = dll.MustFindProc("create_marker")
		createVirtualEntityGroupProc = dll.MustFindProc("create_virtual_entity_group")
		createVirtualEntityProc = dll.MustFindProc("create_virtual_entity")
		createObjectProc = dll.MustFindProc("create_object")
		createCheckpointProc = dll.MustFindProc("create_checkpoint")
		createVehicleProc = dll.MustFindProc("create_vehicle")
		createBlipProc = dll.MustFindProc("create_blip")
		createPedProc = dll.MustFindProc("create_ped")
		createColshapeProc = dll.MustFindProc("create_colshape")
		createPolygonColshapeProc = dll.MustFindProc("create_polygon_colshape")
		createVoiceChannelProc = dll.MustFindProc("create_voice_channel")
		getEntityDataProc = dll.MustFindProc("get_entity_data")
		setServerDataProc = dll.MustFindProc("set_server_data")
		setEntityDataProc = dll.MustFindProc("set_entity_data")
		getColshapeDataProc = dll.MustFindProc("get_colshape_data")
		getDataProc = dll.MustFindProc("get_data")
		getServerDataProc = dll.MustFindProc("get_server_data")
		emitProc = dll.MustFindProc("emit")
		emitAllPlayerProc = dll.MustFindProc("emit_all")
		onClientEventProc = dll.MustFindProc("on_client_event")
	}
}

func (w *WindowsWrapper) ModuleMain(altVersion, core, resourceName, resourceHandlers, moduleHandlers uintptr) bool {
	ret, _, err := mainProc.Call(altVersion, core, resourceName, resourceHandlers, moduleHandlers)
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.Logger().LogErrorf("load mounted failed: %v", err.Error())
		os.Exit(-1)
	}
	return ret != 0
}

func (w *WindowsWrapper) SetVirtualEntityData(id uint32, virtualEntityDataType enums.VirtualEntityDataType, data int64, metaData uint64) {
	_, _, err := setVirtualEntityDataProc.Call(uintptr(id), uintptr(virtualEntityDataType), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set virtual entity data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetPedData(id uint32, pedDataType enums.PedDataType, data int64) {
	_, _, err := setPedDataProc.Call(uintptr(id), uintptr(pedDataType), uintptr(data), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetCheckpointData(id uint32, checkpointDataType enums.CheckpointDataType, data int64, metaData uint64, otherData float32, r, g, b, a uint8) {
	_, _, err := setCheckpointDataProc.Call(uintptr(id), uintptr(checkpointDataType), uintptr(data), uintptr(metaData), uintptr(math.Float32bits(otherData)), uintptr(r), uintptr(g), uintptr(b), uintptr(a))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set checkpoint data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetMarkerData(id uint32, markerDataType enums.MarkerDataType, data int64, metaData uint64, r, g, b, a uint8) {
	_, _, err := setMarkerDataProc.Call(uintptr(id), uintptr(markerDataType), uintptr(data), uintptr(metaData), uintptr(r), uintptr(g), uintptr(b), uintptr(a))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set marker data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetObjectData(id uint32, objectDataType enums.ObjectDataType, data int64, metaData uint64) {
	_, _, err := setObjectDataProc.Call(uintptr(id), uintptr(objectDataType), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set object data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetColshapeData(id uint32, colshapeDataType enums.ColshapeDataType, data int64, metaData uint64) {
	_, _, err := setColshapeDataProc.Call(uintptr(id), uintptr(colshapeDataType), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set colshape data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) Emit(id uint32, eventName, data string) {
	eventNamePtr, freeEventNameCStringFunc := w.GoStringMarshalPtr(eventName)
	dataPtr, freeDataCStringFunc := w.GoStringMarshalPtr(data)
	defer func() {
		freeEventNameCStringFunc()
		freeDataCStringFunc()
	}()
	_, _, err := emitProc.Call(uintptr(id), eventNamePtr, dataPtr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("emit failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) EmitAllPlayer(eventName, data string) {
	eventNamePtr, freeEventNameCStringFunc := w.GoStringMarshalPtr(eventName)
	dataPtr, freeDataCStringFunc := w.GoStringMarshalPtr(data)
	defer func() {
		freeEventNameCStringFunc()
		freeDataCStringFunc()
	}()
	_, _, err := emitAllPlayerProc.Call(eventNamePtr, dataPtr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("emit all failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) OnClientEvent(eventName string, eventArgsDump string) {
	eventNamePtr, freeEventNameCStringFunc := w.GoStringMarshalPtr(eventName)
	argsDumpDataPtr, freeArgsDumpDataCStringFunc := w.GoStringMarshalPtr(eventArgsDump)
	defer func() {
		freeEventNameCStringFunc()
		freeArgsDumpDataCStringFunc()
	}()
	_, _, err := onClientEventProc.Call(eventNamePtr, argsDumpDataPtr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("on client event failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetServerData(setType int32, data int64, strData string) (uintptr, func()) {
	var strDataPtr uintptr
	var freeDataCStringFunc func()
	if len(strData) >= 0 {
		strDataPtr, freeDataCStringFunc = w.GoStringMarshalPtr(strData)
		defer freeDataCStringFunc()
	}
	ret, _, err := setServerDataProc.Call(uintptr(setType), uintptr(data), strDataPtr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set server data failed: %v", err.Error())
		return 0, func() {}
	}
	freeDataResultFunc := func() {
		if ret != 0 {
			w.FreeDataResult(ret)
		}
	}
	return ret, freeDataResultFunc
}

func (w *WindowsWrapper) GetServerData(getType int32, data uint32) (uintptr, func()) {
	ret, _, err := getServerDataProc.Call(uintptr(getType), uintptr(data))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("get server data failed: %v", err.Error())
		return 0, func() {}
	}
	freeDataResultFunc := func() {
		if ret != 0 {
			w.FreeDataResult(ret)
		}
	}
	return ret, freeDataResultFunc
}

func (w *WindowsWrapper) GetColshapeData(id uint32, objectType enums.ObjectType, dataType enums.ColshapeDataType, entityType enums.ObjectType, data int64, metaData uint64) (uintptr, func()) {
	ret, _, err := getColshapeDataProc.Call(uintptr(id), uintptr(objectType), uintptr(dataType), uintptr(entityType), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("get colshape data failed: %v", err.Error())
		return 0, func() {}
	}
	freeDataResultFunc := func() {
		if ret != 0 {
			w.FreeDataResult(ret)
		}
	}
	return ret, freeDataResultFunc
}

func (w *WindowsWrapper) GetData(id uint32, objectType enums.ObjectType, dataType uint8) (uintptr, func()) {
	ret, _, err := getDataProc.Call(uintptr(id), uintptr(objectType), uintptr(dataType), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("get data failed: %v", err.Error())
		return 0, func() {}
	}
	freeDataResultFunc := func() {
		if ret != 0 {
			w.FreeDataResult(ret)
		}
	}
	return ret, freeDataResultFunc
}

func (w *WindowsWrapper) GetMetaData(id uint32, objectType enums.ObjectType, dataType uint8, metaData int64) (uintptr, func()) {
	ret, _, err := getDataProc.Call(uintptr(id), uintptr(objectType), uintptr(dataType), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("get meta data failed: %v", err.Error())
		return 0, func() {}
	}
	freeDataResultFunc := func() {
		if ret != 0 {
			w.FreeDataResult(ret)
		}
	}
	return ret, freeDataResultFunc
}

func (w *WindowsWrapper) SetPedMetaData(id uint32, pedDataType enums.PedDataType, data int64, metaData uint64) {
	_, _, err := setPedDataProc.Call(uintptr(id), uintptr(pedDataType), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetBlipData(id uint32, blipDataType enums.BlipDataType, data int64) {
	_, _, err := setBlipDataProc.Call(uintptr(id), uintptr(blipDataType), uintptr(data), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set blip data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetBlipMetaData(id uint32, blipDataType enums.BlipDataType, data int64, metaData uint64, strData string, r, g, b, a uint8) {
	var freeCStringFunc func()
	var strPtr = uintptr(0)
	if len(strData) > 0 {
		strPtr, freeCStringFunc = w.GoStringMarshalPtr(strData)
		defer freeCStringFunc()
	}
	_, _, err := setBlipDataProc.Call(uintptr(id), uintptr(blipDataType), uintptr(data), uintptr(metaData), strPtr, uintptr(r), uintptr(g), uintptr(b), uintptr(a))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set blip data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetVehicleData(id uint32, vehicleDataType enums.VehicleDataType, data int64) {
	_, _, err := setVehicleDataProc.Call(uintptr(id), uintptr(vehicleDataType), uintptr(data), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set vehicle_hash data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetVehicleMetaData(id uint32, vehicleDataType enums.VehicleDataType, data int64, metaData uint64, strData string, l, r, t, b uint8) {
	var freeCStringFunc func()
	var strPtr = uintptr(0)
	if len(strData) > 0 {
		strPtr, freeCStringFunc = w.GoStringMarshalPtr(strData)
		defer freeCStringFunc()
	}
	_, _, err := setVehicleDataProc.Call(uintptr(id), uintptr(vehicleDataType), uintptr(data), uintptr(metaData), strPtr, uintptr(l), uintptr(r), uintptr(t), uintptr(b))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set vehicle_hash data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetPlayerMetaData(id uint32, playerDataType enums.PlayerDataType, data int64, metaData uint64) {
	_, _, err := setPlayerDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(0), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetPlayerMetaModelData(id uint32, playerDataType enums.PlayerDataType, model uint32, data int64, metaData uint64) {
	_, _, err := setPlayerDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(model), uintptr(data), uintptr(metaData))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) GetEntityData(id uint32, dataType, networkDataType uint8) (uintptr, func()) {
	ret, _, err := getEntityDataProc.Call(uintptr(id), uintptr(dataType), uintptr(networkDataType))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("get entity data failed: %v", err.Error())
		return 0, func() {}
	}
	freeDataResultFunc := func() {
		if ret != 0 {
			w.FreeDataResult(ret)
		}
	}
	return ret, freeDataResultFunc
}

func (w *WindowsWrapper) SetEntityData(id uint32, dataType, entityDataType, entityType uint8, data uint64, metaData uint32, attachData string) {
	var freeCStringFunc func()
	var strPtr = uintptr(0)
	if len(attachData) > 0 {
		strPtr, freeCStringFunc = w.GoStringMarshalPtr(attachData)
		defer freeCStringFunc()
	}
	_, _, err := setEntityDataProc.Call(uintptr(id), uintptr(dataType), uintptr(entityDataType), uintptr(entityType), uintptr(data), uintptr(metaData), strPtr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set entity data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetVoiceChannelData(id uint32, voiceChannelDataType uint8, data int64) {
	_, _, err := setVoiceChannelDataProc.Call(uintptr(id), uintptr(voiceChannelDataType), uintptr(data))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set voice channel data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetNetworkData(id uint32, dataType, networkDataType uint8, keysData, valuesData string) {
	keysDataPtr, freeDataCStringFunc := w.GoStringMarshalPtr(keysData)
	valuesDataPtr, freeValuesDataCStringFunc := w.GoStringMarshalPtr(valuesData)
	defer func() {
		freeValuesDataCStringFunc()
		freeDataCStringFunc()
	}()
	_, _, err := setNetworkDataProc.Call(uintptr(id), uintptr(dataType), uintptr(networkDataType), keysDataPtr, valuesDataPtr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set network data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetPlayerData(id uint32, playerDataType enums.PlayerDataType, data int64) {
	_, _, err := setPlayerDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(0), uintptr(data), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set player data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) SetPlayerHeadData(id uint32, playerDataType enums.PlayerDataType, shape1, shape2, shape3, skin1, skin2, skin3 uint32, shapeMix, skinMix, thirdMix float32) {
	_, _, err := setPlayerHeadDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(shape1), uintptr(shape2), uintptr(shape3), uintptr(skin1), uintptr(skin2), uintptr(skin3), uintptr(math.Float32bits(shapeMix)), uintptr(math.Float32bits(skinMix)), uintptr(math.Float32bits(thirdMix)))
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("set player head data failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) CreateVoiceChannel(spatial uint8, maxDistance float32) (uintptr, func()) {
	ret, _, err := createVoiceChannelProc.Call(uintptr(spatial), uintptr(math.Float32bits(maxDistance)))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.Logger().LogErrorf("create voice channel failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeVoiceChannel(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *WindowsWrapper) CreateVirtualEntityGroup(maxEntitiesInStream uint32) (uintptr, func()) {
	ret, _, err := createVirtualEntityGroupProc.Call(uintptr(maxEntitiesInStream))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.Logger().LogErrorf("create virtual entity group failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeVirtualEntityGroup(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *WindowsWrapper) CreateVirtualEntity(groupId uint32, posData, metaData uint64, streamingDistance uint32) (uintptr, func()) {
	ret, _, err := createVirtualEntityProc.Call(uintptr(groupId), uintptr(posData), uintptr(metaData), uintptr(streamingDistance))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.Logger().LogErrorf("create virtual entity group failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeVirtualEntityGroup(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *WindowsWrapper) CreateCheckpoint(checkPointType uint8, posData, posMetaData uint64, radius, height float32, r, g, b, a uint8, streamingDistance uint32) (uintptr, func()) {
	ret, _, err := createCheckpointProc.Call(uintptr(checkPointType), uintptr(posData), uintptr(posMetaData), uintptr(math.Float32bits(radius)), uintptr(math.Float32bits(height)), uintptr(r), uintptr(g), uintptr(b), uintptr(a), uintptr(streamingDistance))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.Logger().LogErrorf("create checkpoint failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeCheckpoint(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *WindowsWrapper) CreateMarker(markerType uint8, posData, posMetaData uint64, r, g, b, a uint8) (uintptr, func()) {
	ret, _, err := createMarkerProc.Call(uintptr(markerType), uintptr(posData), uintptr(posMetaData), uintptr(r), uintptr(g), uintptr(b), uintptr(a))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.Logger().LogErrorf("create marker failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeMarker(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *WindowsWrapper) CreateObject(model uint32, posData, posMetaData, rotData, rotMetaData uint64) (uintptr, func()) {
	ret, _, err := createObjectProc.Call(uintptr(model), uintptr(posData), uintptr(posMetaData), uintptr(rotData), uintptr(rotMetaData))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.Logger().LogErrorf("create object failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeObject(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *WindowsWrapper) CreateVehicle(model uint32, posData, posMetaData, rotData, rotMetaData uint64, numberplate uintptr, primaryColor, secondColor uint8) (uintptr, func()) {
	ret, _, err := createVehicleProc.Call(uintptr(model), uintptr(posData), uintptr(posMetaData), uintptr(rotData), uintptr(rotMetaData), numberplate, uintptr(primaryColor), uintptr(secondColor))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.Logger().LogErrorf("create vehicle failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeVehicle(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *WindowsWrapper) CreatePed(model uint32, posData, posMetaData, rotData, rotMetaData uint64, streamingDistance uint32) (uintptr, func()) {
	ret, _, err := createPedProc.Call(uintptr(model), uintptr(posData), uintptr(posMetaData), uintptr(rotData), uintptr(rotMetaData), uintptr(streamingDistance))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.Logger().LogErrorf("create vehicle_hash failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreePed(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *WindowsWrapper) CreateBlip(blipType blip_type.BlipType, spriteId, color uint32, strData string, posData, posMetaData uint64, width, height, radius float32) (uintptr, func()) {
	var freeCStringFunc func()
	var strPtr = uintptr(0)
	if len(strData) > 0 {
		strPtr, freeCStringFunc = w.GoStringMarshalPtr(strData)
		defer freeCStringFunc()
	}
	ret, _, err := createBlipProc.Call(uintptr(blipType), uintptr(spriteId), uintptr(color), strPtr, uintptr(posData), uintptr(posMetaData), uintptr(width), uintptr(math.Float32bits(height)), uintptr(math.Float32bits(radius)))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.Logger().LogErrorf("create blip failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeBlip(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *WindowsWrapper) CreateColshape(colshapeType colshape_type.ColshapeType, posData, posMetaData, secondPosData, secondPosMetaData uint64, radius, height float32) (uintptr, func()) {
	ret, _, err := createColshapeProc.Call(uintptr(colshapeType), uintptr(posData), uintptr(posMetaData), uintptr(secondPosData), uintptr(secondPosMetaData), uintptr(math.Float32bits(radius)), uintptr(math.Float32bits(height)))
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.Logger().LogErrorf("create colshape failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeColshape(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *WindowsWrapper) CreatePolygonColshape(colshapeType colshape_type.ColshapeType, minZ, maxZ float32, pointsData []byte) (uintptr, func()) {
	var freeCStringFunc func()
	var strPtr = uintptr(0)
	var strData = string(pointsData)
	if len(strData) > 0 {
		strPtr, freeCStringFunc = w.GoStringMarshalPtr(strData)
		defer freeCStringFunc()
	}
	ret, _, err := createPolygonColshapeProc.Call(uintptr(colshapeType), uintptr(math.Float32bits(minZ)), uintptr(math.Float32bits(maxZ)), strPtr)
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.Logger().LogErrorf("create polygon colshape failed: %v", err.Error())
		return 0, func() {}
	}
	freePtrFunc := func() {
		if ret != 0 {
			w.FreeColshape(ret)
		}
	}
	return ret, freePtrFunc
}

func (w *WindowsWrapper) Free(ptr uintptr) {
	_, _, err := freeProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("free failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) FreePlayer(ptr uintptr) {
	_, _, err := freePlayerProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("free player failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) FreeVehicle(ptr uintptr) {
	_, _, err := freeVehicleProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("free vehicle_hash failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) FreeBlip(ptr uintptr) {
	_, _, err := freeBlipProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("free blip failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) FreePed(ptr uintptr) {
	_, _, err := freePedProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("free ped failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) FreeColshape(ptr uintptr) {
	_, _, err := freeColshapeProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("free colshape failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) FreeCheckpoint(ptr uintptr) {
	_, _, err := freeCheckpointProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("free checkpoint failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) FreeMarker(ptr uintptr) {
	_, _, err := freeMarkerProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("free marker failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) FreeObject(ptr uintptr) {
	_, _, err := freeObjectProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("free object failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) FreeVoiceChannel(ptr uintptr) {
	_, _, err := freeVoiceChannelProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("free voice channel failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) FreeVirtualEntityGroup(ptr uintptr) {
	_, _, err := freeVirtualEntityGroupProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("free virtual entity group failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) FreeVirtualEntity(ptr uintptr) {
	_, _, err := freeVirtualEntityProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("free virtual entity failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) FreeDataResult(ptr uintptr) {
	_, _, err := freeDataResultProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		logger.Logger().LogErrorf("free data result failed: %v", err.Error())
		return
	}
}

func (w *WindowsWrapper) GoStringMarshalPtr(s string) (uintptr, func()) {
	cStr := C.CString(s)
	return uintptr(unsafe.Pointer(cStr)), func() { C.free(unsafe.Pointer(cStr)) }
}

func (w *WindowsWrapper) PtrMarshalGoString(ret uintptr) string {
	cStr := (*C.char)(unsafe.Pointer(ret))
	defer w.Free(ret)
	return C.GoString(cStr)
}
