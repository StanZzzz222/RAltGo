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
	OnStartCallback            func()
	OnStopCallback             func()
	OnServerStartedCallback    func()
	OnPlayerConnectCallback    func(player *models.IPlayer)
	OnPlayerDisconnectCallback func(player *models.IPlayer, reason string)
	OnEnterVehicleCallback     func(player *models.IPlayer, vehicle *models.IVehicle, seat uint8)
	OnLeaveVehicleCallback     func(player *models.IPlayer, vehicle *models.IVehicle, seat uint8)
	OnEnterColshapeCallback    func(colshapeEntityType colshape_entity_type.ColshapeEntityType, player *models.IPlayer, vehicle *models.IVehicle, colshape *models.IColshape)
	OnLeaveColshapeCallback    func(colshapeEntityType colshape_entity_type.ColshapeEntityType, player *models.IPlayer, vehicle *models.IVehicle, colshape *models.IColshape)
)
