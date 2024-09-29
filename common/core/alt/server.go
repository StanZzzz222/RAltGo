package alt

import (
	"github.com/StanZzzz222/RAltGo/common/core/alt/alt_events"
	"github.com/StanZzzz222/RAltGo/common/core/pools"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/internal/entities"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: server.go
*/

func SendBroadcast(message string) {
	alt_events.EmitAllPlayer("chat:message", "", message)
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

func GetEntitiesInRange[T any](position *entities.Vector3, dimension int32, inRange float32) []*T {
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
