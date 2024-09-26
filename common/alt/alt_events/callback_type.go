package alt_events

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/colshape_entity_type"
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
	OnCommandErrorCallback          func(player *models.IPlayer, commandName, desc string)
)
