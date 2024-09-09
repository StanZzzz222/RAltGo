package lib

import "C"
import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"github.com/bwmarrin/snowflake"
	"math"
	"os"
	"sync"
	"syscall"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/5
   File: warpper.go
*/

var dll *syscall.DLL
var freeProc *syscall.Proc
var mainProc *syscall.Proc
var freePlayerProc *syscall.Proc
var freeVehicleProc *syscall.Proc
var spawnPlayerProc *syscall.Proc
var setPlayerDataProc *syscall.Proc
var createVehicleProc *syscall.Proc
var dllPath string
var tasks = &sync.Map{}
var snowflakeNode *snowflake.Node

type Warrper struct{}

//export onTick
func onTick() {
	tasks.Range(func(key, value any) bool {
		handler, ok := value.(func())
		if ok {
			handler()
			tasks.Delete(key)
		}
		return true
	})
}

func (w *Warrper) InitDLL(path string) {
	dllPath = path
	dll = syscall.MustLoadDLL(dllPath)
	mainProc = dll.MustFindProc("init")
	freeProc = dll.MustFindProc("free_c_str")
	freePlayerProc = dll.MustFindProc("free_player")
	freeVehicleProc = dll.MustFindProc("free_vehicle")
	spawnPlayerProc = dll.MustFindProc("spawn_player")
	setPlayerDataProc = dll.MustFindProc("set_player_data")
	createVehicleProc = dll.MustFindProc("create_vehicle")
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	snowflakeNode = node
}

func (w *Warrper) ModuleMain(altVersion, core, resourceName, resourceHandlers, moduleHandlers uintptr) bool {
	ret, _, err := mainProc.Call(altVersion, core, resourceName, resourceHandlers, moduleHandlers)
	if err != nil && err.Error() != "The operation completed successfully." && err.Error() != "The system could not find the environment option that was entered." {
		_ = fmt.Errorf("load init failed: %v", err.Error())
		os.Exit(-1)
	}
	return ret != 0
}

func (w *Warrper) SpawnPlayer(id uint32, hash uint32, x, y, z float32) {
	tasks.Store(snowflakeNode.Generate().String(), func() {
		_, _, err := spawnPlayerProc.Call(uintptr(id), uintptr(hash), uintptr(math.Float32bits(x)), uintptr(math.Float32bits(y)), uintptr(math.Float32bits(z)))
		if err != nil && err.Error() != "The operation completed successfully." {
			_ = fmt.Errorf("spawn player failed: %v", err.Error())
			return
		}
	})
}

func (w *Warrper) SetPlayerMetaData(id uint32, playerDataType enum.PlayerDataType, data uint64, metaData uint64) {
	tasks.Store(snowflakeNode.Generate().String(), func() {
		_, _, err := setPlayerDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(data), uintptr(metaData))
		if err != nil && err.Error() != "The operation completed successfully." {
			_ = fmt.Errorf("set player data failed: %v", err.Error())
			return
		}
	})
}

func (w *Warrper) SetPlayerData(id uint32, playerDataType enum.PlayerDataType, data uint64) {
	tasks.Store(snowflakeNode.Generate().String(), func() {
		_, _, err := setPlayerDataProc.Call(uintptr(id), uintptr(playerDataType), uintptr(data), uintptr(0))
		if err != nil && err.Error() != "The operation completed successfully." {
			_ = fmt.Errorf("set player data failed: %v", err.Error())
			return
		}
	})
}

func (w *Warrper) CreateVehicle(model uint32, posData, posMetaData, rotData, rotMetaData uint64, numberplate uintptr, primaryColor, secondColor uint8) uintptr {
	var ch = make(chan uintptr)
	tasks.Store(snowflakeNode.Generate().String(), func() {
		ret, _, err := createVehicleProc.Call(uintptr(model), uintptr(posData), uintptr(posMetaData), uintptr(rotData), uintptr(rotMetaData), numberplate, uintptr(primaryColor), uintptr(secondColor))
		if err != nil && err.Error() != "The operation completed successfully." {
			_ = fmt.Errorf("set player data failed: %v", err.Error())
			ch <- uintptr(0)
			return
		}
		ch <- ret
	})
	res := <-ch
	return res
}

func (w *Warrper) Free(ptr uintptr) {
	_, _, err := freeProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		_ = fmt.Errorf("free failed: %v", err.Error())
		return
	}
}

func (w *Warrper) FreePlayer(ptr uintptr) {
	_, _, err := freePlayerProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		_ = fmt.Errorf("free player failed: %v", err.Error())
		return
	}
}

func (w *Warrper) FreeVehicle(ptr uintptr) {
	_, _, err := freePlayerProc.Call(ptr)
	if err != nil && err.Error() != "The operation completed successfully." {
		_ = fmt.Errorf("free player failed: %v", err.Error())
		return
	}
}

func (w *Warrper) GoStringMarshalPtr(s string) uintptr {
	cStr := C.CString(s)
	return uintptr(unsafe.Pointer(cStr))
}

func (w *Warrper) PtrMarshalGoString(ret uintptr) string {
	cStr := (*C.char)(unsafe.Pointer(ret))
	return C.GoString(cStr)
}
