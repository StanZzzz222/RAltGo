package alt_events

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_entity_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/denied_reason_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/weapon_hash"
)

/*
   Create by zyx
   Date Time: 2024/9/19
   File: callback_type.go
*/

type (
	OnStartCallback                 func()
	OnStopCallback                  func()
	OnServerStartedCallback         func()
	OnPlayerConnectCallback         func(player *models.IPlayer)
	OnPlayerSpawnCallback           func(player *models.IPlayer)
	OnPlayerDisconnectCallback      func(player *models.IPlayer, reason string)
	OnPlayerInteriorChangeCallback  func(player *models.IPlayer, oldInterior, newInterior uint32)
	OnPlayerDimensionChangeCallback func(player *models.IPlayer, oldDimension, newDimension int32)
	OnEnteringVehicleCallback       func(player *models.IPlayer, vehicle *models.IVehicle, seat uint8)
	OnConsoleCommandCallback        func(name string, args []string)
	OnNetOwnerChangeCallback        func(entity any, oldNetOwner *models.IPlayer, newNetOwner *models.IPlayer)
	OnChangeVehicleSeatCallback     func(player *models.IPlayer, vehicle *models.IVehicle, oldSeat, newSeat uint8)
	OnEnterVehicleCallback          func(player *models.IPlayer, vehicle *models.IVehicle, seat uint8)
	OnLeaveVehicleCallback          func(player *models.IPlayer, vehicle *models.IVehicle, seat uint8)
	OnEnterColshapeCallback         func(colshapeEntityType colshape_entity_type.ColshapeEntityType, player *models.IPlayer, vehicle *models.IVehicle, colshape *models.IColshape)
	OnLeaveColshapeCallback         func(colshapeEntityType colshape_entity_type.ColshapeEntityType, player *models.IPlayer, vehicle *models.IVehicle, colshape *models.IColshape)
	OnChatMessageCallback           func(player *models.IPlayer, message string)
	OnCommandErrorCallback          func(player *models.IPlayer, existCommand bool, commandName, desc string)
	OnPlayerDeathCallback           func(player *models.IPlayer, killer any, weaponHash weapon_hash.ModelHash)
	OnPlayerDamageCallback          func(player *models.IPlayer, attacker any, healthDamage, armourDamage uint16)
	OnPlayerWeaponChangeCallback    func(player *models.IPlayer, oldWeaponHash weapon_hash.ModelHash, newWeaponHash weapon_hash.ModelHash)
	OnPlayerConnectDeniedCallback   func(reason denied_reason_type.DeniedReason, name, ip string, passwordHash uint64, isDebug bool, branch string, majorVersion, minorVersion uint16, cdnUrl string, discordId int64)
	OnPlayerHealCallback            func(player *models.IPlayer, oldHealth, newHealth, oldArmour, newArmour uint16)
	OnPlayerRequestControlCallback  func(player *models.IPlayer, entity any)
	OnVehicleAttachCallback         func(vehicle *models.IVehicle, attached *models.IVehicle)
	OnVehicleDetachCallback         func(vehicle *models.IVehicle, detached *models.IVehicle)
	OnVehicleDestroyCallback        func(vehicle *models.IVehicle)
)
