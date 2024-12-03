package models

import (
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"github.com/StanZzzz222/RAltGo/hash_enums/marker_type"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enums"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/19
   File: colshape.go
*/

type IMarker struct {
	id         uint32
	markerType marker_type.MarkerType
	position   *Vector3
	rotation   *Vector3
	dimension  int32
	visible    bool
	bobUpDown  bool
	faceCamera bool
	rotating   bool
	dir        *Vector3
	scale      *Vector3
	color      *Rgba
	warpper    *lib.Warpper
	*NetworkData
}

func (m *IMarker) GetId() uint32                         { return m.id }
func (m *IMarker) GetMarkerType() marker_type.MarkerType { return m.markerType }
func (m *IMarker) GetPosition() *Vector3                 { return m.position }
func (m *IMarker) GetDimension() int32                   { return m.dimension }
func (m *IMarker) GetVisible() bool                      { return m.visible }
func (m *IMarker) GetBobUpDown() bool                    { return m.bobUpDown }
func (m *IMarker) GetFaceCamera() bool                   { return m.faceCamera }
func (m *IMarker) GetRotating() bool                     { return m.rotating }
func (m *IMarker) GetColor() *Rgba                       { return m.color }
func (m *IMarker) GetDir() *Vector3                      { return m.dir }
func (m *IMarker) GetScale() *Vector3                    { return m.scale }
func (m *IMarker) GetTarget() *IPlayer {
	ret, freeDataResultFunc := m.warpper.GetData(m.id, enums.Marker, uint8(enums.MarkerTarget))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		if p, ok := GetPools().GetPlayerPools().Load(cDataResult.U32Val); ok {
			return p.(*IPlayer)
		}
	}
	return nil
}

func (m *IMarker) NewIMarker(id uint32, markerType uint8, position *Vector3) *IMarker {
	return &IMarker{
		id:          id,
		markerType:  marker_type.MarkerType(markerType),
		position:    position,
		dimension:   hash_enums.DefaultDimension,
		visible:     true,
		bobUpDown:   false,
		faceCamera:  false,
		rotating:    false,
		dir:         nil,
		scale:       nil,
		color:       nil,
		warpper:     lib.GetWrapper(),
		NetworkData: NewNetworkData(id, enums.Marker),
	}
}

func (m *IMarker) SetDimension(dimension int32) {
	m.dimension = dimension
	m.warpper.SetMarkerData(m.id, enums.MarkerDimension, int64(dimension), 0, 0, 0, 0, 0)
}

func (m *IMarker) SetVisible(visible bool) {
	m.visible = visible
	value := 0
	if visible {
		value = 1
	}
	m.warpper.SetMarkerData(m.id, enums.MarkerVisible, int64(value), 0, 0, 0, 0, 0)
}

func (m *IMarker) SetColor(color *Rgba) {
	m.color = color
	m.warpper.SetMarkerData(m.id, enums.MarkerColor, 0, 0, color.R, color.G, color.B, color.A)
}

func (m *IMarker) SetBobUpDown(bobUpDown bool) {
	m.bobUpDown = bobUpDown
	value := 0
	if bobUpDown {
		value = 1
	}
	m.warpper.SetMarkerData(m.id, enums.MarkerBobUpDown, int64(value), 0, 0, 0, 0, 0)
}

func (m *IMarker) SetFaceCamera(faceCamera bool) {
	m.faceCamera = faceCamera
	value := 0
	if faceCamera {
		value = 1
	}
	m.warpper.SetMarkerData(m.id, enums.MarkerFaceCamera, int64(value), 0, 0, 0, 0, 0)
}

func (m *IMarker) SetRotating(rotating bool) {
	m.rotating = rotating
	value := 0
	if rotating {
		value = 1
	}
	m.warpper.SetMarkerData(m.id, enums.MarkerRotating, int64(value), 0, 0, 0, 0, 0)
}

func (m *IMarker) SetMarkerType(markerType marker_type.MarkerType) {
	m.markerType = markerType
	m.warpper.SetMarkerData(m.id, enums.MarkerType, int64(markerType), 0, 0, 0, 0, 0)
}

func (m *IMarker) SetDir(dir *Vector3) {
	m.dir = dir
	posData, posMetaData := int64(math.Float32bits(dir.X))|(int64(math.Float32bits(dir.Y))<<32), uint64(math.Float32bits(dir.Z))<<32
	m.warpper.SetMarkerData(m.id, enums.MarkerDir, posData, posMetaData, 0, 0, 0, 0)
}

func (m *IMarker) SetScale(scale *Vector3) {
	m.scale = scale
	posData, posMetaData := int64(math.Float32bits(scale.X))|(int64(math.Float32bits(scale.Y))<<32), uint64(math.Float32bits(scale.Z))<<32
	m.warpper.SetMarkerData(m.id, enums.MarkerScale, posData, posMetaData, 0, 0, 0, 0)
}

func (m *IMarker) SetPosition(position *Vector3) {
	m.position = position
	posData, posMetaData := int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32
	m.warpper.SetMarkerData(m.id, enums.MarkerPosition, posData, posMetaData, 0, 0, 0, 0)
}

func (m *IMarker) SetRotation(rotation *Vector3) {
	m.rotation = rotation
	rotData, rotMetaData := int64(math.Float32bits(rotation.X))|(int64(math.Float32bits(rotation.Y))<<32), uint64(math.Float32bits(rotation.Z))<<32
	m.warpper.SetMarkerData(m.id, enums.MarkerRotation, rotData, rotMetaData, 0, 0, 0, 0)
}

func (m *IMarker) Destroy() {
	m.warpper.SetMarkerData(m.id, enums.MarkerDestory, int64(0), 0, 0, 0, 0, 0)
	pools.DestroyMarker(m)
}
