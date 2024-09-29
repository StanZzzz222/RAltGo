package models

import (
	"github.com/StanZzzz222/RAltGo/internal/enums"
)

/*
   Create by zyx
   Date Time: 2024/9/19
   File: colshape.go
*/

type IVirtualEntityGroup struct {
	id                  uint32
	maxEntitiesInStream uint32
	*NetworkData
}

func (v *IVirtualEntityGroup) GetId() uint32                  { return v.id }
func (v *IVirtualEntityGroup) GetMaxEntitiesInStream() uint32 { return v.maxEntitiesInStream }

func (v *IVirtualEntityGroup) NewIVirtualEntityGroup(id, maxEntitiesInStream uint32) *IVirtualEntityGroup {
	return &IVirtualEntityGroup{
		id:                  id,
		maxEntitiesInStream: maxEntitiesInStream,
		NetworkData:         NewNetworkData(id, enums.VirtualEntityGroup),
	}
}

func (v *IVirtualEntityGroup) SetData(key string, value any) {
	v.datas.Store(key, value)
}

func (v *IVirtualEntityGroup) DelData(key string) {
	_, ok := v.datas.Load(key)
	if ok {
		v.datas.Delete(key)
	}
}

func (v *IVirtualEntityGroup) DelAllData() {
	v.datas.Range(func(key, value any) bool {
		v.datas.Delete(key)
		return true
	})
}

func (v *IVirtualEntityGroup) HasData(key string) bool {
	_, ok := v.datas.Load(key)
	if ok {
		return true
	}
	return false
}

func (v *IVirtualEntityGroup) GetData(key string) any {
	value, ok := v.datas.Load(key)
	if ok {
		return value
	}
	return value
}

func (v *IVirtualEntityGroup) GetDatas() []any {
	var datas []any
	v.datas.Range(func(key, value any) bool {
		datas = append(datas, key)
		return true
	})
	return datas
}
