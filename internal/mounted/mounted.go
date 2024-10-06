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
	logger.Logger().LogInfo(":: Go module Initialize mounting...")
	return w.ModuleMain(uintptr(cAltvVersion), uintptr(core), uintptr(cResourceName), uintptr(cResourceHandlers), uintptr(cModuleHandlers))
}

//export onStart
func onStart() {
	defer panicRecover()
	s := scheduler.NewScheduler()
	s.AddTask(func() {
		alt_events.Events().OnClientEvent("chat:message", func(player *models.IPlayer, message string) {
			if message[0] == '/' {
				s = scheduler.NewScheduler()
				args := strings.Split(message, " ")
				groups := command.GetCommandGroups()
				for _, group := range groups {
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
				}
				s.Run()
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
		defer func() {
			w.FreePlayer(cPtr)
			pools := models.GetPools()
			pools.PutPlayer(player)
		}()
		player = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.SocialName, cPlayer.SocialID, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
		alt_events.Triggers().TriggerOnPlayerConnect(player)
	}
}

//export onPlayerDisconnect
func onPlayerDisconnect(cPtr, cReasonPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var player = &models.IPlayer{}
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		p := models.GetPools().GetPlayer(cPlayer.ID)
		reason := w.PtrMarshalGoString(cReasonPtr)
		defer func() {
			w.FreePlayer(cPtr)
			pools := models.GetPools()
			pools.DestroyPlayer(player)
		}()
		if p == nil {
			p = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.SocialName, cPlayer.SocialID, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
		}
		alt_events.Triggers().TriggerOnPlayerDisconnect(p, reason)
	}
}

//export onEnterVehicle
func onEnterVehicle(cPtr, cvPtr uintptr, seat uint8) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var player = &models.IPlayer{}
	var veh = &models.IVehicle{}
	var cPlayer = entities.ConvertCPlayer(cPtr)
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	if cPlayer != nil && cVehicle != nil {
		p := models.GetPools().GetPlayer(cPlayer.ID)
		v := models.GetPools().GetVehicle(cVehicle.ID)
		defer func() {
			w.FreePlayer(cPtr)
			w.FreeVehicle(cvPtr)
		}()
		if p == nil {
			p = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.SocialName, cPlayer.SocialID, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
			models.GetPools().PutPlayer(p)
		}
		if v == nil {
			v = veh.NewIVehicle(cVehicle.ID, cVehicle.Model, cVehicle.PrimaryColor, cVehicle.SecondColor, cVehicle.Position, cVehicle.Rotation)
			models.GetPools().PutVehicle(veh)
		}
		alt_events.Triggers().TriggerOnEnterVehicle(p, v, seat)
	}
}

//export onEnteringVehicle
func onEnteringVehicle(cPtr, cvPtr uintptr, seat uint8) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var player = &models.IPlayer{}
	var veh = &models.IVehicle{}
	var cPlayer = entities.ConvertCPlayer(cPtr)
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	if cPlayer != nil && cVehicle != nil {
		p := models.GetPools().GetPlayer(cPlayer.ID)
		v := models.GetPools().GetVehicle(cVehicle.ID)
		defer func() {
			w.FreePlayer(cPtr)
			w.FreeVehicle(cvPtr)
		}()
		if p == nil {
			p = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.SocialName, cPlayer.SocialID, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
			models.GetPools().PutPlayer(p)
		}
		if v == nil {
			v = veh.NewIVehicle(cVehicle.ID, cVehicle.Model, cVehicle.PrimaryColor, cVehicle.SecondColor, cVehicle.Position, cVehicle.Rotation)
			models.GetPools().PutVehicle(veh)
		}
		alt_events.Triggers().TriggerOnEnteringVehicle(p, v, seat)
	}
}

//export onConsoleCommand
func onConsoleCommand(cNamePtr, cArgsPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var args []string
	sName := w.PtrMarshalGoString(cNamePtr)
	sArgs := w.PtrMarshalGoString(cArgsPtr)
	_ = json.Unmarshal([]byte(sArgs), &args)
	alt_events.Triggers().TriggerOnConsoleCommand(sName, args)
}

//export onNetOwnerChange
func onNetOwnerChange(objectType uint8, entityId, oldNetOwnerId, newNetOwnerId uint32) {
	defer panicRecover()
	pools := models.GetPools()
	switch enums.ObjectType(objectType) {
	case enums.Player:
		entity, oldNetOwner, newNetOwner := pools.GetPlayer(entityId), pools.GetPlayer(oldNetOwnerId), pools.GetPlayer(newNetOwnerId)
		alt_events.Triggers().TriggerOnNetOwnerChange(entity, oldNetOwner, newNetOwner)
		break
	case enums.Vehicle:
		entity, oldNetOwner, newNetOwner := pools.GetVehicle(entityId), pools.GetPlayer(oldNetOwnerId), pools.GetPlayer(newNetOwnerId)
		alt_events.Triggers().TriggerOnNetOwnerChange(entity, oldNetOwner, newNetOwner)
		break
	case enums.Ped:
		entity, oldNetOwner, newNetOwner := pools.GetPed(entityId), pools.GetPlayer(oldNetOwnerId), pools.GetPlayer(newNetOwnerId)
		alt_events.Triggers().TriggerOnNetOwnerChange(entity, oldNetOwner, newNetOwner)
		break
	case enums.Object:
		entity, oldNetOwner, newNetOwner := pools.GetObject(entityId), pools.GetPlayer(oldNetOwnerId), pools.GetPlayer(newNetOwnerId)
		alt_events.Triggers().TriggerOnNetOwnerChange(entity, oldNetOwner, newNetOwner)
		break
	default:
		break
	}
}

//export onPlayerSpawn
func onPlayerSpawn(cPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var player = &models.IPlayer{}
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		p := models.GetPools().GetPlayer(cPlayer.ID)
		defer w.FreePlayer(cPtr)
		if p == nil {
			p = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.SocialName, cPlayer.SocialID, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
			models.GetPools().PutPlayer(p)
		}
		alt_events.Triggers().TriggerOnPlayerSpawn(p)
	}
}

//export onInteriorChange
func onInteriorChange(cPtr uintptr, oldInterior, newInterior uint32) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var player = &models.IPlayer{}
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		p := models.GetPools().GetPlayer(cPlayer.ID)
		defer w.FreePlayer(cPtr)
		if p == nil {
			p = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.SocialName, cPlayer.SocialID, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
			models.GetPools().PutPlayer(p)
		}
		alt_events.Triggers().TriggerOnPlayerInteriorChange(p, oldInterior, newInterior)
	}
}

//export onPlayerDimensionChange
func onPlayerDimensionChange(cPtr uintptr, oldDimension, newDimension int32) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var player = &models.IPlayer{}
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		p := models.GetPools().GetPlayer(cPlayer.ID)
		defer w.FreePlayer(cPtr)
		if p == nil {
			p = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.SocialName, cPlayer.SocialID, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
			models.GetPools().PutPlayer(p)
		}
		alt_events.Triggers().TriggerOnPlayerDimensionChange(p, oldDimension, newDimension)
	}
}

//export onChangeVehicleSeat
func onChangeVehicleSeat(cPtr, cvPtr uintptr, oldSeat, newSeat uint8) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var player = &models.IPlayer{}
	var veh = &models.IVehicle{}
	var cPlayer = entities.ConvertCPlayer(cPtr)
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	if cPlayer != nil && cVehicle != nil {
		p := models.GetPools().GetPlayer(cPlayer.ID)
		v := models.GetPools().GetVehicle(cVehicle.ID)
		defer func() {
			w.FreePlayer(cPtr)
			w.FreeVehicle(cvPtr)
		}()
		if p == nil {
			p = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.SocialName, cPlayer.SocialID, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
			models.GetPools().PutPlayer(p)
		}
		if v == nil {
			v = veh.NewIVehicle(cVehicle.ID, cVehicle.Model, cVehicle.PrimaryColor, cVehicle.SecondColor, cVehicle.Position, cVehicle.Rotation)
			models.GetPools().PutVehicle(veh)
		}
		alt_events.Triggers().TriggerOnChangeVehicleSeat(p, v, oldSeat, newSeat)
	}
}

//export onLeaveVehicle
func onLeaveVehicle(cPtr, cvPtr uintptr, seat uint8) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var player = &models.IPlayer{}
	var veh = &models.IVehicle{}
	var cPlayer = entities.ConvertCPlayer(cPtr)
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	if cPlayer != nil && cVehicle != nil {
		p := models.GetPools().GetPlayer(cPlayer.ID)
		v := models.GetPools().GetVehicle(cVehicle.ID)
		defer func() {
			w.FreePlayer(cPtr)
			w.FreeVehicle(cvPtr)
		}()
		if p == nil {
			p = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.SocialName, cPlayer.SocialID, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
			models.GetPools().PutPlayer(p)
		}
		if v == nil {
			v = veh.NewIVehicle(cVehicle.ID, cVehicle.Model, cVehicle.PrimaryColor, cVehicle.SecondColor, cVehicle.Position, cVehicle.Rotation)
			models.GetPools().PutVehicle(veh)
		}
		alt_events.Triggers().TriggerOnLeaveVehicle(p, v, seat)
	}
}

//export onEnterColshape
func onEnterColshape(cType uint8, cPtr, cvPtr, ccPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var player = &models.IPlayer{}
	var veh = &models.IVehicle{}
	var colshape = &models.IColshape{}
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
			p := models.GetPools().GetPlayer(cPlayer.ID)
			c := models.GetPools().GetColshape(cColshape.ID)
			if p == nil {
				p = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.SocialName, cPlayer.SocialID, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
				models.GetPools().PutPlayer(p)
			}
			if c == nil {
				c = colshape.NewIColshape(cColshape.ID, cColshape.ColshapeType, cColshape.Position)
				models.GetPools().PutColshape(c)
			}
			alt_events.Triggers().TriggerOnEnterColshape(colshapeEntityType, p, nil, c)
		}
		break
	case colshape_entity_type.Vehicle:
		if cVehicle != nil && cColshape != nil {
			v := models.GetPools().GetVehicle(cVehicle.ID)
			c := models.GetPools().GetColshape(cColshape.ID)
			if v == nil {
				v = veh.NewIVehicle(cVehicle.ID, cVehicle.Model, cVehicle.PrimaryColor, cVehicle.SecondColor, cVehicle.Position, cVehicle.Rotation)
				models.GetPools().PutVehicle(veh)
			}
			if c == nil {
				c = colshape.NewIColshape(cColshape.ID, cColshape.ColshapeType, cColshape.Position)
				models.GetPools().PutColshape(c)
			}
			alt_events.Triggers().TriggerOnEnterColshape(colshapeEntityType, nil, v, c)
		}
		break
	}
}

//export onLeaveColshape
func onLeaveColshape(cType uint8, cPtr, cvPtr, ccPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var player = &models.IPlayer{}
	var veh = &models.IVehicle{}
	var colshape = &models.IColshape{}
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
			p := models.GetPools().GetPlayer(cPlayer.ID)
			c := models.GetPools().GetColshape(cColshape.ID)
			if p == nil {
				p = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.SocialName, cPlayer.SocialID, cPlayer.HWIDHash, cPlayer.HWIDExHash, cPlayer.Position, cPlayer.Rotation)
				models.GetPools().PutPlayer(p)
			}
			if c == nil {
				c = colshape.NewIColshape(cColshape.ID, cColshape.ColshapeType, cColshape.Position)
				models.GetPools().PutColshape(c)
			}
			alt_events.Triggers().TriggerOnLeaveColshape(colshapeEntityType, p, nil, c)
		}
		break
	case colshape_entity_type.Vehicle:
		if cVehicle != nil && cColshape != nil {
			v := models.GetPools().GetVehicle(cVehicle.ID)
			c := models.GetPools().GetColshape(cColshape.ID)
			if v == nil {
				v = veh.NewIVehicle(cVehicle.ID, cVehicle.Model, cVehicle.PrimaryColor, cVehicle.SecondColor, cVehicle.Position, cVehicle.Rotation)
				models.GetPools().PutVehicle(veh)
			}
			if c == nil {
				c = colshape.NewIColshape(cColshape.ID, cColshape.ColshapeType, cColshape.Position)
				models.GetPools().PutColshape(c)
			}
			alt_events.Triggers().TriggerOnLeaveColshape(colshapeEntityType, nil, v, c)
		}
		break
	}
}

//export onClientEvent
func onClientEvent(cPlayerId uint32, cEventNamePtr, cEventArgsPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	p := models.GetPools().GetPlayer(cPlayerId)
	eventName := w.PtrMarshalGoString(cEventNamePtr)
	eventArgs := w.PtrMarshalGoString(cEventArgsPtr)
	alt_events.Triggers().TriggerOnClientEvent(p, eventName, eventArgs)
}

func panicRecover() {
	if r := recover(); r != nil {
		var stackBuf [4096]byte
		stackSize := runtime.Stack(stackBuf[:], false)
		logger.Logger().LogErrorf("Panic recovered: %v", r)
		logger.Logger().LogErrorf("StackTrace: %s", stackBuf[:stackSize])
	}
}
