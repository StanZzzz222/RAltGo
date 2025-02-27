package models

import (
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"github.com/StanZzzz222/RAltGo/internal/enums"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/19
   File: colshape.go
*/

type IVirtualEntity struct {
	id                uint32
	streamingDistance uint32
	position          *Vector3
	visible           bool
	dimension         int32
	warpper           *lib.Wrapper
	*NetworkData
}

func (v *IVirtualEntity) GetId() uint32                { return v.id }
func (v *IVirtualEntity) GetPosition() *Vector3        { return v.position }
func (v *IVirtualEntity) GetStreamingDistance() uint32 { return v.streamingDistance }
func (v *IVirtualEntity) GetDimension() int32          { return v.dimension }

func (v *IVirtualEntity) NewIVirtualEntity(id, streamingDistance uint32, position *Vector3) *IVirtualEntity {
	return &IVirtualEntity{
		id:                id,
		streamingDistance: streamingDistance,
		position:          position,
		dimension:         hash_enums.DefaultDimension,
		visible:           true,
		NetworkData:       NewNetworkData(id, enums.VirtualEntity),
	}
}

func (v *IVirtualEntity) SetDimension(dimension int32) {
	v.dimension = dimension
	v.warpper.SetVirtualEntityData(v.id, enums.VirtualEntityDimension, int64(dimension), 0)
}

func (v *IVirtualEntity) SetVisible(visible bool) {
	v.visible = visible
	value := 0
	if visible {
		value = 1
	}
	v.warpper.SetVirtualEntityData(v.id, enums.VirtualEntityVisible, int64(value), 0)
}

func (v *IVirtualEntity) SetPosition(position *Vector3) {
	v.position = position
	posData, posMetaData := int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	v.warpper.SetVirtualEntityData(v.id, enums.VirtualEntityPosition, posData, posMetaData)
}

func (v *IVirtualEntity) Destroy() {
	v.warpper.SetVirtualEntityData(v.id, enums.VirtualEntityDestory, 0, 0)
}
