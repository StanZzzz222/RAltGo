package mounted

import "C"
import (
	"github.com/StanZzzz222/RAltGo/common/command"
	"github.com/StanZzzz222/RAltGo/common/core/alt/alt_events"
	"github.com/StanZzzz222/RAltGo/common/core/scheduler"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_entity_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/denied_reason_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/weapon_hash"
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
		player = player.NewIPlayer(cPlayer.ID, cPlayer.Name, cPlayer.IP, cPlayer.AuthToken, cPlayer.SocialName, cPlayer.SocialID, cPlayer.HWIDHash, cPlayer.HWIDExHash, (*models.Vector3)(cPlayer.Position), (*models.Vector3)(cPlayer.Rotation))
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
		if cPlayer != nil {
			w.FreePlayer(cPtr)
		}
		if cVehicle != nil {
			w.FreeVehicle(cvPtr)
		}
		if cColshape != nil {
			w.FreeColshape(ccPtr)
		}
	}()
	switch colshapeEntityType {
	case colshape_entity_type.Player:
		if cPlayer != nil && cColshape != nil {
			alt_events.Triggers().TriggerOnEnterColshape(colshapeEntityType, models.GetPools().GetPlayer(cPlayer.ID), nil, models.GetPools().GetColshape(cColshape.ID))
		}
		break
	case colshape_entity_type.Vehicle:
		if cVehicle != nil && cColshape != nil {
			alt_events.Triggers().TriggerOnEnterColshape(colshapeEntityType, models.GetPools().GetPlayer(cPlayer.ID), models.GetPools().GetVehicle(cVehicle.ID), models.GetPools().GetColshape(cColshape.ID))
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
			alt_events.Triggers().TriggerOnLeaveColshape(colshapeEntityType, models.GetPools().GetPlayer(cPlayer.ID), models.GetPools().GetVehicle(cVehicle.ID), models.GetPools().GetColshape(cColshape.ID))
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

//export onPlayerDeath
func onPlayerDeath(cPtr uintptr, deathType, objectType, objectId uint8, weaponHash uint32) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		defer w.FreePlayer(cPtr)
		switch deathType {
		case 0:
			var killer any
			switch enums.ObjectType(objectType) {
			case enums.Player:
				killer = models.GetPools().GetPlayer(uint32(objectId))
				break
			case enums.Vehicle:
				killer = models.GetPools().GetVehicle(uint32(objectId))
				break
			case enums.Ped:
				killer = models.GetPools().GetPed(uint32(objectId))
				break
			case enums.Object:
				killer = models.GetPools().GetObject(uint32(objectId))
				break
			default:
				killer = nil
				break
			}
			alt_events.Triggers().TriggerOnPlayerDeath(models.GetPools().GetPlayer(cPlayer.ID), killer, weapon_hash.ModelHash(weaponHash))
			break
		case 1:
			alt_events.Triggers().TriggerOnPlayerDeath(models.GetPools().GetPlayer(cPlayer.ID), nil, weapon_hash.ModelHash(weaponHash))
			break
		}
	}
}

//export onPlayerDamage
func onPlayerDamage(cPtr uintptr, deathType, objectType, objectId uint8, healthDamage, armourDamage uint16) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		defer w.FreePlayer(cPtr)
		switch deathType {
		case 0:
			var attacker any
			switch enums.ObjectType(objectType) {
			case enums.Player:
				attacker = models.GetPools().GetPlayer(uint32(objectId))
				break
			case enums.Vehicle:
				attacker = models.GetPools().GetVehicle(uint32(objectId))
				break
			case enums.Ped:
				attacker = models.GetPools().GetPed(uint32(objectId))
				break
			case enums.Object:
				attacker = models.GetPools().GetObject(uint32(objectId))
				break
			default:
				attacker = nil
				break
			}
			alt_events.Triggers().TriggerOnPlayerDamage(models.GetPools().GetPlayer(cPlayer.ID), attacker, healthDamage, armourDamage)
			break
		case 1:
			alt_events.Triggers().TriggerOnPlayerDamage(models.GetPools().GetPlayer(cPlayer.ID), nil, healthDamage, armourDamage)
			break
		}
	}
}

//export onPlayerWeaponChange
func onPlayerWeaponChange(cPtr uintptr, oldWeaponHash, newWeaponHash uint32) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		defer w.FreePlayer(cPtr)
		alt_events.Triggers().TriggerOnPlayerWeaponChange(models.GetPools().GetPlayer(cPlayer.ID), oldWeaponHash, newWeaponHash)
	}
}

//export onPlayerConnectDenied
func onPlayerConnectDenied(cReason uint8, cNamePtr, cIpPtr uintptr, passwordHash uint64, isDebug bool, cBranchPtr uintptr, majorVersion, minorVersion uint16, cCdnUrlPtr uintptr, discordId int64) {
	defer panicRecover()
	var w = lib.GetWarpper()
	alt_events.Triggers().TriggerOnPlayerConnectDenied(denied_reason_type.DeniedReason(cReason), w.PtrMarshalGoString(cNamePtr), w.PtrMarshalGoString(cIpPtr),
		passwordHash, isDebug, w.PtrMarshalGoString(cBranchPtr), majorVersion, minorVersion, w.PtrMarshalGoString(cCdnUrlPtr), discordId)
}

//export onPlayerHeal
func onPlayerHeal(cPtr uintptr, oldHealth, newHealth, oldArmour, newArmour uint16) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		defer w.FreePlayer(cPtr)
		alt_events.Triggers().TriggerOnPlayerHeal(models.GetPools().GetPlayer(cPlayer.ID), oldHealth, newHealth, oldArmour, newArmour)
	}
}

//export onPlayerRequestControl
func onPlayerRequestControl(cPtr uintptr, objectType, objectId uint8) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cPlayer != nil {
		defer w.FreePlayer(cPtr)
		var entity any
		switch enums.ObjectType(objectType) {
		case enums.Player:
			entity = models.GetPools().GetPlayer(uint32(objectId))
			break
		case enums.Vehicle:
			entity = models.GetPools().GetVehicle(uint32(objectId))
			break
		case enums.Ped:
			entity = models.GetPools().GetPed(uint32(objectId))
			break
		case enums.Object:
			entity = models.GetPools().GetObject(uint32(objectId))
			break
		default:
			entity = nil
			break
		}
		alt_events.Triggers().TriggerOnPlayerRequestControl(models.GetPools().GetPlayer(cPlayer.ID), entity)
	}
}

//export onVehicleAttach
func onVehicleAttach(cvPtr, caPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	var cAttached = entities.ConvertCVehicle(caPtr)
	if cVehicle != nil && cAttached != nil {
		defer func() {
			w.FreeVehicle(cvPtr)
			w.FreeVehicle(caPtr)
		}()
		alt_events.Triggers().TriggerOnVehicleAttach(models.GetPools().GetVehicle(cVehicle.ID), models.GetPools().GetVehicle(cAttached.ID))
	}
}

//export onVehicleDetach
func onVehicleDetach(cvPtr, cdPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	var cDetached = entities.ConvertCVehicle(cdPtr)
	if cVehicle != nil && cDetached != nil {
		defer func() {
			w.FreeVehicle(cvPtr)
			w.FreeVehicle(cdPtr)
		}()
		alt_events.Triggers().TriggerOnVehicleDetach(models.GetPools().GetVehicle(cVehicle.ID), models.GetPools().GetVehicle(cDetached.ID))
	}
}

//export onVehicleDestroy
func onVehicleDestroy(cvPtr uintptr) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	if cVehicle != nil {
		defer w.FreeVehicle(cvPtr)
		alt_events.Triggers().TriggerOnVehicleDestroy(models.GetPools().GetVehicle(cVehicle.ID))
	}
}

//export onVehicleDamage
func onVehicleDamage(cvPtr uintptr, objectType, objectId uint8, bodyHealthDamage, bodyAdditionalHealthDamage, engineHealthDamage, petrolTankHealthDamage, weapon uint32) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	if cVehicle != nil {
		defer w.FreeVehicle(cvPtr)
		var entity any
		switch enums.ObjectType(objectType) {
		case enums.Player:
			entity = models.GetPools().GetPlayer(uint32(objectId))
			break
		case enums.Vehicle:
			entity = models.GetPools().GetVehicle(uint32(objectId))
			break
		case enums.Ped:
			entity = models.GetPools().GetPed(uint32(objectId))
			break
		case enums.Object:
			entity = models.GetPools().GetObject(uint32(objectId))
			break
		default:
			entity = nil
			break
		}
		alt_events.Triggers().TriggerOnVehicleDamage(models.GetPools().GetVehicle(cVehicle.ID), entity, bodyHealthDamage, bodyAdditionalHealthDamage, engineHealthDamage, petrolTankHealthDamage, weapon_hash.ModelHash(weapon))
	}
}

//export onVehicleHorn
func onVehicleHorn(cvPtr, cPtr uintptr, state bool) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	var cPlayer = entities.ConvertCPlayer(cPtr)
	if cVehicle != nil && cPlayer != nil {
		defer w.FreeVehicle(cvPtr)
		defer w.FreePlayer(cPtr)
		alt_events.Triggers().TriggerOnVehicleHorn(models.GetPools().GetVehicle(cVehicle.ID), models.GetPools().GetPlayer(cPlayer.ID), state)
	}
}

//export onVehicleSiren
func onVehicleSiren(cvPtr uintptr, state bool) {
	defer panicRecover()
	var w = lib.GetWarpper()
	var cVehicle = entities.ConvertCVehicle(cvPtr)
	if cVehicle != nil {
		defer w.FreeVehicle(cvPtr)
		alt_events.Triggers().TriggerOnVehicleSiren(models.GetPools().GetVehicle(cVehicle.ID), state)
	}
}

func panicRecover() {
	if r := recover(); r != nil {
		var stackBuf [4096]byte
		stackSize := runtime.Stack(stackBuf[:], false)
		logger.Logger().LogErrorf("Panic recovered: %v", r)
		logger.Logger().LogErrorf("StackTrace: %s", stackBuf[:stackSize])
	}
}
