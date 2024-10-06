package mounted

import "C"
import (
	"github.com/StanZzzz222/RAltGo/common/command"
	"github.com/StanZzzz222/RAltGo/common/core/alt/alt_events"
	"github.com/StanZzzz222/RAltGo/common/core/scheduler"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_entity_type"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enums"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"github.com/StanZzzz222/RAltGo/logger"
	"github.com/goccy/go-json"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/9
   File: mounted.go
*/

func Mounted() {}

//export onModuleInit
func onModuleInit(cAltvVersion, core, cResourceName, cResourceHandlers, cModuleHandlers unsafe.Pointer) bool {
	defer panicRecover()
	var w = lib.GetWarpper()
	return w.ModuleMain(uintptr(cAltvVersion), uintptr(core), uintptr(cResourceName), uintptr(cResourceHandlers), uintptr(cModuleHandlers))
}

//export onStart
func onStart() {
	defer panicRecover()
	s := scheduler.NewScheduler()
	s.AddTask(func() {
		alt_events.Events().OnClientEvent("chat:message", func(player *models.IPlayer, message string) {
			if message[0] == '/' {
				exist := false
				s = scheduler.NewScheduler()
				args := strings.Split(message, " ")
				groups := command.GetCommandGroups()
				for _, group := range groups {
					if _, ok := group.GetCommand(args[0]); ok {
						exist = true
						s.AddTask(func() {
							var params []any
							for _, param := range args[1:] {
								if strings.Contains(param, ".") {
									if value, err := strconv.ParseFloat(param, 64); err == nil {
										params = append(params, value)
										continue
									}
									params = append(params, param)
									continue
								}
								if value, err := strconv.ParseInt(param, 10, 64); err == nil {
									params = append(params, value)
									continue
								}
								if value, err := strconv.ParseBool(param); err == nil {
									params = append(params, value)
									continue
								}
								params = append(params, param)
							}
							group.TriggerCommand(args[0], player, params...)
						})
						break
					}
				}
				s.Run()
				if !exist {
					alt_events.Triggers().TriggerOnCommandError(player, false, args[0], "")
				}
				return
			}
			alt_events.Triggers().TriggerOnChatMessage(player, message)
		})
	})
	s.Run()
	alt_events.Triggers().TriggerOnStart()
}

//export onServerStarted
func onServerStarted() {
	defer panicRecover()
	alt_events.Triggers().TriggerOnServerStarted()
}

//export onStop
func onStop() {
	defer panicRecover()
	alt_events.Triggers().TriggerOnStop()
}

//export onPlayerConnect
func onPlayerConnect(cPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var player = &models.IPlayer{}
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		defer w.FreePlayer(cPtr)
		player = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.SocialName, cPlayer.SocialID, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
		models.GetPools().PutPlayer(player)
		alt_events.Triggers().TriggerOnPlayerConnect(player)
	}
}

//export onPlayerDisconnect
func onPlayerDisconnect(cPtr, cReasonPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		defer w.FreePlayer(cPtr)
		p := models.GetPools().GetPlayer(cPlayer.ID)
		reason := w.PtrMarshalGoString(cReasonPtr)
		alt_events.Triggers().TriggerOnPlayerDisconnect(p, reason)
		models.GetPools().DestroyPlayer(p)
	}
}

//export onEnterVehicle
func onEnterVehicle(cPtr, cvPtr uintptr, seat uint8) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cPlayer = entities.ConvertCPlayer(cPtr)
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	if cPlayer != nil && cVehicle != nil {
		defer func() {
			w.FreePlayer(cPtr)
			w.FreeVehicle(cvPtr)
		}()
		alt_events.Triggers().TriggerOnEnterVehicle(models.GetPools().GetPlayer(cPlayer.ID), models.GetPools().GetVehicle(cVehicle.ID), seat)
	}
}

//export onEnteringVehicle
func onEnteringVehicle(cPtr, cvPtr uintptr, seat uint8) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cPlayer = entities.ConvertCPlayer(cPtr)
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	if cPlayer != nil && cVehicle != nil {
		defer func() {
			w.FreePlayer(cPtr)
			w.FreeVehicle(cvPtr)
		}()
		alt_events.Triggers().TriggerOnEnteringVehicle(models.GetPools().GetPlayer(cPlayer.ID), models.GetPools().GetVehicle(cVehicle.ID), seat)
	}
}

//export onConsoleCommand
func onConsoleCommand(cNamePtr, cArgsPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var args []string
	sArgs := w.PtrMarshalGoString(cArgsPtr)
	_ = json.Unmarshal([]byte(sArgs), &args)
	alt_events.Triggers().TriggerOnConsoleCommand(w.PtrMarshalGoString(cNamePtr), args)
}

//export onNetOwnerChange
func onNetOwnerChange(objectType uint8, entityId, oldNetOwnerId, newNetOwnerId uint32) {
	defer panicRecover()
	pools := models.GetPools()
	switch enums.ObjectType(objectType) {
	case enums.Player:
		alt_events.Triggers().TriggerOnNetOwnerChange(pools.GetPlayer(entityId), pools.GetPlayer(oldNetOwnerId), pools.GetPlayer(newNetOwnerId))
		break
	case enums.Vehicle:
		alt_events.Triggers().TriggerOnNetOwnerChange(pools.GetVehicle(entityId), pools.GetPlayer(oldNetOwnerId), pools.GetPlayer(newNetOwnerId))
		break
	case enums.Ped:
		alt_events.Triggers().TriggerOnNetOwnerChange(pools.GetPed(entityId), pools.GetPlayer(oldNetOwnerId), pools.GetPlayer(newNetOwnerId))
		break
	case enums.Object:
		alt_events.Triggers().TriggerOnNetOwnerChange(pools.GetObject(entityId), pools.GetPlayer(oldNetOwnerId), pools.GetPlayer(newNetOwnerId))
		break
	default:
		break
	}
}

//export onPlayerSpawn
func onPlayerSpawn(cPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		defer w.FreePlayer(cPtr)
		alt_events.Triggers().TriggerOnPlayerSpawn(models.GetPools().GetPlayer(cPlayer.ID))
	}
}

//export onInteriorChange
func onInteriorChange(cPtr uintptr, oldInterior, newInterior uint32) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		defer w.FreePlayer(cPtr)
		alt_events.Triggers().TriggerOnPlayerInteriorChange(models.GetPools().GetPlayer(cPlayer.ID), oldInterior, newInterior)
	}
}

//export onPlayerDimensionChange
func onPlayerDimensionChange(cPtr uintptr, oldDimension, newDimension int32) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		defer w.FreePlayer(cPtr)
		alt_events.Triggers().TriggerOnPlayerDimensionChange(models.GetPools().GetPlayer(cPlayer.ID), oldDimension, newDimension)
	}
}

//export onChangeVehicleSeat
func onChangeVehicleSeat(cPtr, cvPtr uintptr, oldSeat, newSeat uint8) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cPlayer = entities.ConvertCPlayer(cPtr)
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	if cPlayer != nil && cVehicle != nil {
		defer func() {
			w.FreePlayer(cPtr)
			w.FreeVehicle(cvPtr)
		}()
		alt_events.Triggers().TriggerOnChangeVehicleSeat(models.GetPools().GetPlayer(cPlayer.ID), models.GetPools().GetVehicle(cVehicle.ID), oldSeat, newSeat)
	}
}

//export onLeaveVehicle
func onLeaveVehicle(cPtr, cvPtr uintptr, seat uint8) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cPlayer = entities.ConvertCPlayer(cPtr)
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	if cPlayer != nil && cVehicle != nil {
		defer func() {
			w.FreePlayer(cPtr)
			w.FreeVehicle(cvPtr)
		}()
		alt_events.Triggers().TriggerOnLeaveVehicle(models.GetPools().GetPlayer(cPlayer.ID), models.GetPools().GetVehicle(cVehicle.ID), seat)
	}
}

//export onEnterColshape
func onEnterColshape(cType uint8, cPtr, cvPtr, ccPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var colshapeEntityType = colshape_entity_type.ColshapeEntityType(cType)
	var cPlayer = entities.ConvertCPlayer(cPtr)
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	var cColshape = entities.ConvertCColshape(ccPtr)
	defer func() {
		if cPtr != 0 {
			w.FreePlayer(cPtr)
		}
		if cvPtr != 0 {
			w.FreeVehicle(cvPtr)
		}
		w.FreeColshape(ccPtr)
	}()
	switch colshapeEntityType {
	case colshape_entity_type.Player:
		if cPlayer != nil && cColshape != nil {
			alt_events.Triggers().TriggerOnEnterColshape(colshapeEntityType, models.GetPools().GetPlayer(cPlayer.ID), nil, models.GetPools().GetColshape(cColshape.ID))
		}
		break
	case colshape_entity_type.Vehicle:
		if cVehicle != nil && cColshape != nil {
			alt_events.Triggers().TriggerOnEnterColshape(colshapeEntityType, nil, models.GetPools().GetVehicle(cVehicle.ID), models.GetPools().GetColshape(cColshape.ID))
		}
		break
	}
}

//export onLeaveColshape
func onLeaveColshape(cType uint8, cPtr, cvPtr, ccPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var colshapeEntityType = colshape_entity_type.ColshapeEntityType(cType)
	var cPlayer = entities.ConvertCPlayer(cPtr)
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	var cColshape = entities.ConvertCColshape(ccPtr)
	defer func() {
		if cPtr != 0 {
			w.FreePlayer(cPtr)
		}
		if cvPtr != 0 {
			w.FreeVehicle(cvPtr)
		}
		w.FreeColshape(ccPtr)
	}()
	switch colshapeEntityType {
	case colshape_entity_type.Player:
		if cPlayer != nil && cColshape != nil {
			alt_events.Triggers().TriggerOnLeaveColshape(colshapeEntityType, models.GetPools().GetPlayer(cPlayer.ID), nil, models.GetPools().GetColshape(cColshape.ID))
		}
		break
	case colshape_entity_type.Vehicle:
		if cVehicle != nil && cColshape != nil {
			alt_events.Triggers().TriggerOnLeaveColshape(colshapeEntityType, nil, models.GetPools().GetVehicle(cVehicle.ID), models.GetPools().GetColshape(cColshape.ID))
		}
		break
	}
}

//export onClientEvent
func onClientEvent(cPlayerId uint32, cEventNamePtr, cEventArgsPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	alt_events.Triggers().TriggerOnClientEvent(models.GetPools().GetPlayer(cPlayerId), w.PtrMarshalGoString(cEventNamePtr), w.PtrMarshalGoString(cEventArgsPtr))
}

func panicRecover() {
	if r := recover(); r != nil {
		var stackBuf [4096]byte
		stackSize := runtime.Stack(stackBuf[:], false)
		logger.Logger().LogErrorf("Panic recovered: %v", r)
		logger.Logger().LogErrorf("StackTrace: %s", stackBuf[:stackSize])
	}
}
