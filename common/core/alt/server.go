package alt

import (
	"github.com/BurntSushi/toml"
	"github.com/StanZzzz222/RAltGo/common"
	"github.com/StanZzzz222/RAltGo/common/core/alt/alt_events"
	"github.com/StanZzzz222/RAltGo/common/core/pools"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums/ammo_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/closest_entities_order_type"
	"github.com/StanZzzz222/RAltGo/hash_enums/weapon_hash"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enums"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"github.com/StanZzzz222/RAltGo/logger"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: server.go
*/

type ServerConfig struct {
	Modules                    []string       `toml:"modules"`
	Resources                  []string       `toml:"resources"`
	Name                       string         `toml:"name"`
	Host                       string         `toml:"host"`
	Port                       int32          `toml:"port"`
	Players                    int32          `toml:"players"`
	Password                   string         `toml:"password"`
	Announce                   bool           `toml:"announce"`
	Token                      string         `toml:"token"`
	Gamemode                   string         `toml:"gamemode"`
	Website                    string         `toml:"website"`
	Language                   string         `toml:"language"`
	Description                string         `toml:"description"`
	Debug                      bool           `toml:"debug"`
	StreamingDistance          float64        `toml:"streaming_distance"`
	MigrationDistance          float64        `toml:"migration_distance"`
	Timeout                    float64        `toml:"timeout"`
	AnnounceRetryErrorDelay    int32          `toml:"announce_retry_error_delay"`
	AnnounceRetryErrorAttempts int32          `toml:"announce_retry_error_attempts"`
	DuplicatePlayers           int32          `toml:"duplicate_players"`
	MapBoundsMinX              float64        `toml:"map_bounds_min_x"`
	MapBoundsMinY              float64        `toml:"map_bounds_min_y"`
	MapBoundsMaxX              float64        `toml:"map_bounds_max_x"`
	MapBoundsMaxY              float64        `toml:"map_bounds_max_y"`
	MapCellAreaSize            float64        `toml:"map_cell_area_size"`
	ColShapeTickRate           int32          `toml:"col_shape_tick_rate"`
	LogStreams                 []string       `toml:"log_streams"`
	EntityWorkerCount          int32          `toml:"entity_worker_count"`
	Tags                       []string       `toml:"tags"`
	ConnectionQueue            bool           `toml:"connection_queue"`
	UseEarlyAuth               bool           `toml:"use_early_auth"`
	EarlyAuthURL               string         `toml:"early_auth_url"`
	UseCDN                     bool           `toml:"use_cdn"`
	CDNURL                     string         `toml:"cdn_url"`
	SendPlayerNames            bool           `toml:"send_player_names"`
	SpawnAfterConnect          bool           `toml:"spawn_after_connect"`
	WorldProfiler              map[string]any `toml:"world_profiler"`
	JsModule                   map[string]any `toml:"js_module"`
	CsharpModule               map[string]any `toml:"csharp_module"`
	Voice                      map[string]any `toml:"voice"`
}

func GetServerConfig() *ServerConfig {
	var serverConfig *ServerConfig
	if _, err := toml.DecodeFile("./server.toml", &serverConfig); err != nil {
		logger.Logger().LogErrorf("Read server.toml falied, %v", err.Error())
		return nil
	}
	return serverConfig
}

func SendBroadcast(message string) {
	alt_events.EmitAllPlayer("chat:message", "", message)
}

func KickPlayer(player *models.IPlayer, reason string) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.KickPlayer, int64(player.GetId()), reason)
}

func StopServer() {
	var w = lib.GetWrapper()
	w.SetServerData(enums.StopServer, int64(0), "")
}

func ToggleWorldProfiler(toggleWorldProfiler bool) {
	var w = lib.GetWrapper()
	value := 0
	if toggleWorldProfiler {
		value = 1
	}
	w.SetServerData(enums.ToggleWorldProfiler, int64(value), "")
}

func SetStreamingDistance(streamingDistance uint32) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerStreamingDistance, int64(streamingDistance), "")
}

func SetMigrationDistance(migrationDistance uint32) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerMigrationDistance, int64(migrationDistance), "")
}

func SetPassword(password string) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerPassword, 0, password)
}

func SetColShapeTickRate(colshapeTickRate uint32) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerColShapeTickRate, int64(colshapeTickRate), "")
}

func SetMigrationTickRate(migrationTickRate uint32) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerMigrationTickRate, int64(migrationTickRate), "")
}

func SetMaxStreamingObjects(maxStreamingObjects uint16) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerMaxStreamingObjects, int64(maxStreamingObjects), "")
}

func SetMaxStreamingPeds(maxStreamingPeds uint16) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerMaxStreamingPeds, int64(maxStreamingPeds), "")
}

func SetMaxStreamingVehicles(maxStreamingVehicles uint16) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerMaxStreamingVehicles, int64(maxStreamingVehicles), "")
}

func SetMigrationThreadCount(migrationThreadCount uint16) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerMigrationThreadCount, int64(migrationThreadCount), "")
}

func SetStreamerThreadCount(streamerThreadCount uint16) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerStreamerThreadCount, int64(streamerThreadCount), "")
}

func SetStreamingTickRate(streamingTickRate uint16) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerStreamingTickRate, int64(streamingTickRate), "")
}

func SetSyncReceiveThreadCount(syncReceiveThreadCount uint16) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerSyncReceiveThreadCount, int64(syncReceiveThreadCount), "")
}

func SetSyncSendThreadCount(syncSendThreadCount uint16) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerSyncSendThreadCount, int64(syncSendThreadCount), "")
}

func SetVoiceExternal(host string, port uint16) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerVoiceExternal, int64(port), host)
}

func SetVoiceExternalPublic(host string, port uint16) {
	var w = lib.GetWrapper()
	w.SetServerData(enums.ServerVoiceExternalPublic, int64(port), host)
}

func GetNetTime() uint32 {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.NetTime, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U32Val
	}
	return 0
}

func GetStreamingDistance() uint32 {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerStreamingDistance, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U32Val
	}
	return 0
}

func GetMigrationDistance() uint32 {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerMigrationDistance, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U32Val
	}
	return 0
}

func GetColShapeTickRate() uint32 {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerColShapeTickRate, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U32Val
	}
	return 0
}

func GetMigrationTickRate() uint32 {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerMigrationTickRate, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U32Val
	}
	return 0
}

func GetMaxStreamingObjects() uint16 {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerMaxStreamingObjects, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U16Val
	}
	return 0
}

func GetMaxStreamingPeds() uint16 {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerMaxStreamingPeds, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U16Val
	}
	return 0
}

func GetMaxStreamingVehicles() uint16 {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerMaxStreamingVehicles, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U16Val
	}
	return 0
}

func GetMigrationThreadCount() uint8 {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerMigrationThreadCount, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U8Val
	}
	return 0
}

func GetStreamerThreadCount() uint8 {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerStreamerThreadCount, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U8Val
	}
	return 0
}

func GetStreamingTickRate() uint32 {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerStreamingTickRate, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U32Val
	}
	return 0
}

func GetSyncReceiveThreadCount() uint8 {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerSyncReceiveThreadCount, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U8Val
	}
	return 0
}

func GetSyncSendThreadCount() uint8 {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerSyncSendThreadCount, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U8Val
	}
	return 0
}

func GetVoiceConnectionState() enums.VoiceConnectionState {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerVoiceConnectionState, 0)
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return enums.VoiceConnectionState(cDataResult.U8Val)
	}
	return 0
}

func GetAmmoHashForWeaponHash(weaponHash weapon_hash.ModelHash) ammo_type.AmmoHash {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerAmmoHashForWeaponHash, uint32(weaponHash))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return ammo_type.AmmoHash(cDataResult.U32Val)
	}
	return 0
}

func GetAmmoHashForWeaponName(weaponName string) ammo_type.AmmoHash {
	var w = lib.GetWrapper()
	ret, freeDataResultFunc := w.GetServerData(enums.ServerAmmoHashForWeaponHash, common.Hash(weaponName))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return ammo_type.AmmoHash(cDataResult.U32Val)
	}
	return 0
}

func GetEntitiesInDimension[T any](dimension int32) []*T {
	var entitys []*T
	for entity := range pools.GetAnyEntityIterator[T]() {
		if e, ok := any(entity).(*models.IPlayer); ok {
			if e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IBlip); ok {
			if e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IPed); ok {
			if e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IVehicle); ok {
			if e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IObject); ok {
			if e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IMarker); ok {
			if e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.ICheckpoint); ok {
			if e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IColshape); ok {
			if e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IVirtualEntity); ok {
			if e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
	}
	return entitys
}

func GetEntitiesInRange[T any](position *models.Vector3, dimension int32, inRange float32) []*T {
	var entitys []*T
	for entity := range pools.GetAnyEntityIterator[T]() {
		if e, ok := any(entity).(*models.IPlayer); ok {
			if e.GetPosition().Distance(position) <= inRange && e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IBlip); ok {
			if e.GetPosition().Distance(position) <= inRange && e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IPed); ok {
			if e.GetPosition().Distance(position) <= inRange && e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IVehicle); ok {
			if e.GetPosition().Distance(position) <= inRange && e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IObject); ok {
			if e.GetPosition().Distance(position) <= inRange && e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IMarker); ok {
			if e.GetPosition().Distance(position) <= inRange && e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.ICheckpoint); ok {
			if e.GetPosition().Distance(position) <= inRange && e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IColshape); ok {
			if e.GetPosition().Distance(position) <= inRange && e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
		if e, ok := any(entity).(*models.IVirtualEntity); ok {
			if e.GetPosition().Distance(position) <= inRange && e.GetDimension() == dimension {
				entitys = append(entitys, entity)
			}
		}
	}
	return entitys
}

func GetClosestEntities[T any](position *models.Vector3, dimension int32, inRange float32, limit int32, closestEntitiesOrderType closest_entities_order_type.ClosestEntitiesOrderType) []*T {
	entitys := GetEntitiesInRange[T](position, dimension, inRange)
	switch closestEntitiesOrderType {
	case closest_entities_order_type.ClosestEntitiesOrderDefault, closest_entities_order_type.ClosestEntitiesOrderAsc:
		entitiesQuickSort(position, entitys, 0, len(entitys)-1, closest_entities_order_type.ClosestEntitiesOrderAsc)
	case closest_entities_order_type.ClosestEntitiesOrderDesc:
		entitiesQuickSort(position, entitys, 0, len(entitys)-1, closest_entities_order_type.ClosestEntitiesOrderDesc)
	default:
	}
	return entitys[:limit]
}

func entitiesQuickSort[T any](target *models.Vector3, datas []*T, low, high int, closestEntitiesOrderType closest_entities_order_type.ClosestEntitiesOrderType) {
	if low < high {
		p := entitiesPartition(target, datas, low, high, closestEntitiesOrderType)
		entitiesQuickSort(target, datas, low, p-1, closestEntitiesOrderType)
		entitiesQuickSort(target, datas, p+1, high, closestEntitiesOrderType)
	}
}

func entitiesPartition[T any](target *models.Vector3, datas []*T, low, high int, closestEntitiesOrderType closest_entities_order_type.ClosestEntitiesOrderType) int {
	i := low - 1
	pivot := getEntityRange(target, datas[high])
	if closestEntitiesOrderType == closest_entities_order_type.ClosestEntitiesOrderAsc {
		for j := low; j < high; j++ {
			if getEntityRange(target, datas[j]) <= pivot {
				i++
				datas[i], datas[j] = datas[j], datas[i]
			}
		}
	} else {
		for j := low; j < high; j++ {
			if getEntityRange(target, datas[j]) >= pivot {
				i++
				datas[i], datas[j] = datas[j], datas[i]
			}
		}
	}
	datas[i+1], datas[high] = datas[high], datas[i+1]
	return i + 1
}

func getEntityRange[T any](target *models.Vector3, elem *T) float32 {
	if e, ok := any(elem).(*models.IPlayer); ok {
		return e.GetPosition().Distance(target)
	}
	if e, ok := any(elem).(*models.IBlip); ok {
		return e.GetPosition().Distance(target)
	}
	if e, ok := any(elem).(*models.IPed); ok {
		return e.GetPosition().Distance(target)
	}
	if e, ok := any(elem).(*models.IVehicle); ok {
		return e.GetPosition().Distance(target)
	}
	if e, ok := any(elem).(*models.IObject); ok {
		return e.GetPosition().Distance(target)
	}
	if e, ok := any(elem).(*models.IMarker); ok {
		return e.GetPosition().Distance(target)
	}
	if e, ok := any(elem).(*models.ICheckpoint); ok {
		return e.GetPosition().Distance(target)
	}
	if e, ok := any(elem).(*models.IColshape); ok {
		return e.GetPosition().Distance(target)
	}
	if e, ok := any(elem).(*models.IVirtualEntity); ok {
		return e.GetPosition().Distance(target)
	}
	return 0
}
