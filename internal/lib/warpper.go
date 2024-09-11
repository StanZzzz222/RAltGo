package lib

import "C"
import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"github.com/StanZzzz222/RAltGo/logger"
	"github.com/bwmarrin/snowflake"
	"math"
	"os"
	"sync"
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
var tasks = &sync.Map{}
var snowflakeNode *snowflake.Node
var freeProc *syscall.Proc
var mainProc *syscall.Proc
var freePlayerProc *syscall.Proc
var freeVehicleProc *syscall.Proc
var spawnPlayerProc *syscall.Proc
var setVehicleDataProc *syscall.Proc
var setPlayerDataProc *syscall.Proc
var createVehicleProc *syscall.Proc

type Warrper struct{}

//export onTick
func onTick() {
	tasks.Range(func(key, value any) bool {
		handler, ok := value.(func())
		if ok {
			tasks.Delete(key)
			handler()
		}
		return true
	})
}

func init() {
	path, _ := os.Getwd()
	path = fmt.Sprintf("%v/modules/rs-go-module.dll", path)
	node, err := snowflake.NewNode(1)
	if err != nil {
		logger.LogErrorf("Snowflake NewNode err: %v", err)
		return
	}
	snowflakeNode = node
	_, err = os.Stat(path)
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
	spawnPlayerProc = dll.MustFindProc("spawn_player")
	setPlayerDataProc = dll.MustFindProc("set_player_data")
	setVehicleDataProc = dll.MustFindProc("set_vehicle_data")
	createVehicleProc = dll.MustFindProc("create_vehicle")
}

func (w *Warrper) ModuleMain(altVersion, core, resourceName, resourceHandlers, moduleHandlers uintptr) bool {
	ret, _, err := mainProc.Call(altVersion, core, resourceName, resourceHandlers, moduleHandlers)
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		logger.LogErrorf("load mounted failed: %v", err.Error())
		os.Exit(-1)
	}
	return ret != 0
}

func (w *Warrper) SpawnPlayer(id uint32, hash uint32, x, y, z float32) {
	tasks.Store(snowflakeNode.Generate().String(), func() {
		_, _, err := spawnPlayerProc.Call(uintptr(id), uintptr(hash), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
		if err != nil && err.Error() != "The operation completed successfully." {
			logger.LogErrorf("spawn player failed: %v", err.Error())
			return
		}
	})
}

func (w *Warrper) SetVehicleData(id uint32, vehicleDataType enum.VehicleDataType, data uint64) {
	tasks.Store(snowflakeNode.Generate().String(), func() {
		_, _, err := setVehicleDataProc.Call(uintptr(id), uintptr(vehicleDataType), uintptr(data), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0), uintptr(0))
		if err != nil && err.Error() != "The operation completed successfully." {
			logger.LogErrorf("set vehicle data failed: %v", err.Error())
			return
		}
	})
}

func (w *Warrper) SetVehicleMetaData(id uint32, vehicleDataType enum.VehicleDataType, data, metaData uint64, strData string, l, r, t, b uint8) {
	tasks.Store(snowflakeNode.Generate().String(), func() {
		_, _, err := setVehicleDataProc.Call(uintptr(id), uintptr(vehicleDataType), uintptr(data), uintptr(metaData), w.GoStringMarshalPtr(strData), uintptr(l), uintptr(r), uintptr(t), uintptr(b))
		if err != nil && err.Error() != "The operation completed successfully." {
			logger.LogErrorf("set vehicle data failed: %v", err.Error())
			return
		}
	})
}

func (w *Warrper) SetPlayerMetaData(id uint32, playerDataType enum.PlayerDataType, data int64, metaData uint64) {
	tasks.Store(snowflakeNode.Generate().String(), func() {
		_, _, err := setPlayerDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(data), uintptr(metaData))
		if err != nil && err.Error() != "The operation completed successfully." {
			logger.LogErrorf("set player data failed: %v", err.Error())
			return
		}
	})
}

func (w *Warrper) SetPlayerData(id uint32, playerDataType enum.PlayerDataType, data int64) {
	tasks.Store(snowflakeNode.Generate().String(), func() {
		_, _, err := setPlayerDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(data), uintptr(0))
		if err != nil && err.Error() != "The operation completed successfully." {
			logger.LogErrorf("set player data failed: %v", err.Error())
			return
		}
	})
}

func (w *Warrper) CreateVehicle(model uint32, posData, posMetaData, rotData, rotMetaData uint64, numberplate uintptr, primaryColor, secondColor uint8) uintptr {
	var ch = make(chan uintptr)
	tasks.Store(snowflakeNode.Generate().String(), func() {
		ret, _, err := createVehicleProc.Call(uintptr(model), uintptr(posData), uintptr(posMetaData), uintptr(rotData), uintptr(rotMetaData), numberplate, uintptr(primaryColor), uintptr(secondColor))
		defer func() {
			if ret != 0 {
				w.FreeVehicle(ret)
			}
		}()
		if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
			logger.LogErrorf("create vehicle failed: %v", err.Error())
			ch <- uintptr(0)
			return
		}
		ch <- ret
	})
	res := <-ch
	close(ch)
	return res
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
		logger.LogErrorf("free vehicle failed: %v", err.Error())
		return
	}
}

func (w *Warrper) GoStringMarshalPtr(s string) uintptr {
	cStr := C.CString(s)
	return uintptr(unsafe.Pointer(&cStr))
}

func (w *Warrper) PtrMarshalGoString(ret uintptr) string {
	cStr := (*C.char)(unsafe.Pointer(ret))
	ptr := uintptr(unsafe.Pointer(&cStr))
	if ptr != 0 {
		defer w.Free(ptr)
	}
	return C.GoString(cStr)
}
