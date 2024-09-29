package pools

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"reflect"
	"sync"
	"unsafe"
)

/*
   Create by zyx
   Date Time: 2024/9/18
   File: blip.go
*/

func GetAnyEntity[T any](id uint32) T {
	var entity T
	tType := reflect.TypeOf(entity)
	switch tType {
	case reflect.TypeOf((*models.IPlayer)(nil)).Elem():
		entity = *(*T)(unsafe.Pointer(models.GetPools().GetPlayer(id)))
		break
	case reflect.TypeOf((*models.IBlip)(nil)).Elem():
		entity = *(*T)(unsafe.Pointer(models.GetPools().GetBlip(id)))
		break
	case reflect.TypeOf((*models.IPed)(nil)).Elem():
		entity = *(*T)(unsafe.Pointer(models.GetPools().GetPed(id)))
		break
	case reflect.TypeOf((*models.IVehicle)(nil)).Elem():
		entity = *(*T)(unsafe.Pointer(models.GetPools().GetVehicle(id)))
		break
	case reflect.TypeOf((*models.IObject)(nil)).Elem():
		entity = *(*T)(unsafe.Pointer(models.GetPools().GetObject(id)))
		break
	case reflect.TypeOf((*models.IMarker)(nil)).Elem():
		entity = *(*T)(unsafe.Pointer(models.GetPools().GetMarker(id)))
		break
	case reflect.TypeOf((*models.ICheckpoint)(nil)).Elem():
		entity = *(*T)(unsafe.Pointer(models.GetPools().GetCheckpoint(id)))
		break
	case reflect.TypeOf((*models.IColshape)(nil)).Elem():
		entity = *(*T)(unsafe.Pointer(models.GetPools().GetColshape(id)))
		break
	case reflect.TypeOf((*models.IVirtualEntity)(nil)).Elem():
		entity = *(*T)(unsafe.Pointer(models.GetPools().GetVirtualEntity(id)))
		break
	case reflect.TypeOf((*models.IVirtualEntityGroup)(nil)).Elem():
		entity = *(*T)(unsafe.Pointer(models.GetPools().GetVirtualEntityGroup(id)))
		break
	}
	return entity
}

func GetAnyEntitys[T any]() []*T {
	var entitys []*T
	var entity T
	var pools *sync.Map
	tType := reflect.TypeOf(entity)
	switch tType {
	case reflect.TypeOf((*models.IPlayer)(nil)).Elem():
		pools = models.GetPools().GetPlayerPools()
		break
	case reflect.TypeOf((*models.IBlip)(nil)).Elem():
		pools = models.GetPools().GetBlipPools()
		break
	case reflect.TypeOf((*models.IPed)(nil)).Elem():
		pools = models.GetPools().GetPedPools()
		break
	case reflect.TypeOf((*models.IVehicle)(nil)).Elem():
		pools = models.GetPools().GetVehiclePools()
		break
	case reflect.TypeOf((*models.IObject)(nil)).Elem():
		pools = models.GetPools().GetObjectPools()
		break
	case reflect.TypeOf((*models.IMarker)(nil)).Elem():
		pools = models.GetPools().GetMarkerPools()
		break
	case reflect.TypeOf((*models.ICheckpoint)(nil)).Elem():
		pools = models.GetPools().GetCheckpointPools()
		break
	case reflect.TypeOf((*models.IColshape)(nil)).Elem():
		pools = models.GetPools().GetColshapePools()
		break
	case reflect.TypeOf((*models.IVirtualEntity)(nil)).Elem():
		pools = models.GetPools().GetVirtualEntityPools()
		break
	case reflect.TypeOf((*models.IVirtualEntityGroup)(nil)).Elem():
		pools = models.GetPools().GetVirtualEntityGroupPools()
		break
	}
	if pools != nil {
		pools.Range(func(key, value any) bool {
			if v, ok := value.(*T); ok {
				entitys = append(entitys, v)
			}
			return true
		})
	}
	return entitys
}

func GetAnyEntityIterator[T any]() <-chan *T {
	var entity T
	var anyEntityChan = make(chan *T)
	var pools *sync.Map
	tType := reflect.TypeOf(entity)
	switch tType {
	case reflect.TypeOf((*models.IPlayer)(nil)).Elem():
		pools = models.GetPools().GetPlayerPools()
		break
	case reflect.TypeOf((*models.IBlip)(nil)).Elem():
		pools = models.GetPools().GetBlipPools()
		break
	case reflect.TypeOf((*models.IPed)(nil)).Elem():
		pools = models.GetPools().GetPedPools()
		break
	case reflect.TypeOf((*models.IVehicle)(nil)).Elem():
		pools = models.GetPools().GetVehiclePools()
		break
	case reflect.TypeOf((*models.IObject)(nil)).Elem():
		pools = models.GetPools().GetObjectPools()
		break
	case reflect.TypeOf((*models.IMarker)(nil)).Elem():
		pools = models.GetPools().GetMarkerPools()
		break
	case reflect.TypeOf((*models.ICheckpoint)(nil)).Elem():
		pools = models.GetPools().GetCheckpointPools()
		break
	case reflect.TypeOf((*models.IColshape)(nil)).Elem():
		pools = models.GetPools().GetColshapePools()
		break
	case reflect.TypeOf((*models.IVirtualEntity)(nil)).Elem():
		pools = models.GetPools().GetVirtualEntityPools()
		break
	case reflect.TypeOf((*models.IVirtualEntityGroup)(nil)).Elem():
		pools = models.GetPools().GetVirtualEntityGroupPools()
		break
	}
	go func() {
		defer close(anyEntityChan)
		if pools != nil {
			pools.Range(func(key, value any) bool {
				if v, ok := value.(*T); ok {
					anyEntityChan <- v
				}
				return true
			})
		}
	}()
	return anyEntityChan
}
