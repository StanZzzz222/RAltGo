package models

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/hash_enums/blip_type"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"math"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/12
   File: blip.go
*/

type IBlip struct {
	id                       uint32
	blipType                 blip_type.BlipType
	color                    uint32
	spriteId                 uint32
	alpha                    uint32
	category                 uint32
	flashInterval            uint32
	flashTimer               uint32
	display                  uint32
	number                   int32
	rot                      float32
	visible                  bool
	friendly                 bool
	highDetail               bool
	missionCreator           bool
	shortRange               bool
	bright                   bool
	flashes                  bool
	global                   bool
	route                    bool
	pulse                    bool
	shrinked                 bool
	showCone                 bool
	tickVisible              bool
	useHeightIndicatorOnEdge bool
	outlineIndicatorVisible  bool
	hiddenOnLegend           bool
	minimalOnEdge            bool
	flashesAlternate         bool
	crewIndicatorVisible     bool
	headingIndicatorVisible  bool
	shortHeightThreshold     bool
	name                     string
	gxtName                  string
	routeColor               *entities.Rgba
	rgbaColor                *entities.Rgba
	position                 *entities.Vector3
	scale                    *entities.Vector3
	datas                    *sync.Map
}

func (b *IBlip) NewIBlip(id, blipType, spriteId, color uint32, name string, rot float32, position *entities.Vector3) *IBlip {
	return &IBlip{
		id:            id,
		blipType:      blip_type.BlipType(blipType),
		spriteId:      spriteId,
		name:          name,
		color:         color,
		rot:           rot,
		position:      position,
		alpha:         255,
		flashInterval: 0,
		flashTimer:    0,
		number:        0,
		display:       2,
		visible:       true,
		datas:         &sync.Map{},
	}
}

func (b *IBlip) GetId() uint32                     { return b.id }
func (b *IBlip) GetName() string                   { return b.name }
func (b *IBlip) GetBlipType() blip_type.BlipType   { return b.blipType }
func (b *IBlip) GetBlipColor() uint32              { return b.color }
func (b *IBlip) GetSpriteId() uint32               { return b.spriteId }
func (b *IBlip) GetAlpha() uint32                  { return b.alpha }
func (b *IBlip) GetFlashInterval() uint32          { return b.flashInterval }
func (b *IBlip) GetFlashTimer() uint32             { return b.flashTimer }
func (b *IBlip) GetVisible() bool                  { return b.visible }
func (b *IBlip) GetDisplay() uint32                { return b.display }
func (b *IBlip) GetFriendly() bool                 { return b.friendly }
func (b *IBlip) GetHighDetail() bool               { return b.highDetail }
func (b *IBlip) GetMissionCreator() bool           { return b.missionCreator }
func (b *IBlip) GetShortRange() bool               { return b.shortRange }
func (b *IBlip) GetBright() bool                   { return b.bright }
func (b *IBlip) GetFlashes() bool                  { return b.flashes }
func (b *IBlip) GetGlobal() bool                   { return b.global }
func (b *IBlip) GetRoute() bool                    { return b.route }
func (b *IBlip) GetPulse() bool                    { return b.pulse }
func (b *IBlip) GetShrinked() bool                 { return b.shrinked }
func (b *IBlip) GetShowCone() bool                 { return b.showCone }
func (b *IBlip) GetTickVisible() bool              { return b.tickVisible }
func (b *IBlip) GetUseHeightIndicatorOnEdge() bool { return b.useHeightIndicatorOnEdge }
func (b *IBlip) GetOutlineIndicatorVisible() bool  { return b.outlineIndicatorVisible }
func (b *IBlip) GetHiddenOnLegend() bool           { return b.hiddenOnLegend }
func (b *IBlip) GetMinimalOnEdge() bool            { return b.minimalOnEdge }
func (b *IBlip) GetFlashesAlternate() bool         { return b.flashesAlternate }
func (b *IBlip) GetCrewIndicatorVisible() bool     { return b.crewIndicatorVisible }
func (b *IBlip) GetHeadingIndicatorVisible() bool  { return b.headingIndicatorVisible }
func (b *IBlip) GetShortHeightThreshold() bool     { return b.shortHeightThreshold }
func (b *IBlip) GetRouteColor() *entities.Rgba     { return b.routeColor }
func (b *IBlip) GetRgbaColor() *entities.Rgba      { return b.rgbaColor }
func (b *IBlip) GetScale() *entities.Vector3       { return b.scale }
func (b *IBlip) GetNumber() int32                  { return b.number }
func (b *IBlip) GetGxtName() string {
	ret, freeDataResultFunc := w.GetData(b.id, enum.Blip, uint8(enum.Category))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.StringVal
	}
	return ""
}
func (b *IBlip) GetCategory() uint32 {
	ret, freeDataResultFunc := w.GetData(b.id, enum.Blip, uint8(enum.Category))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U32Val
	}
	return 0
}
func (b *IBlip) GetPosition() *entities.Vector3 {
	ret, freeDataResultFunc := w.GetData(b.id, enum.Blip, uint8(enum.BlipPosition))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.Vector3Val
	}
	return nil
}
func (b *IBlip) GetPositionString() string {
	position := b.GetPosition()
	return fmt.Sprintf("%v,%v,%v", position.X, position.Y, position.Z)
}

func (b *IBlip) SetSprite(spriteId uint32) {
	b.spriteId = spriteId
	w.SetBlipData(b.id, enum.Sprite, int64(spriteId))
}

func (b *IBlip) SetColor(color uint32) {
	b.color = color
	w.SetBlipData(b.id, enum.Color, int64(color))
}

func (b *IBlip) SetRGBAColor(rgbaColor *entities.Rgba) {
	b.rgbaColor = rgbaColor
	w.SetBlipMetaData(b.id, enum.RgbaColor, 0, 0, "", rgbaColor.R, rgbaColor.G, rgbaColor.B, rgbaColor.A)
}

func (b *IBlip) SetVisible(visible bool) {
	b.visible = visible
	value := 0
	if visible {
		value = 1
	}
	w.SetBlipData(b.id, enum.BlipVisible, int64(value))
}

func (b *IBlip) SetDisplay(display uint32) {
	b.display = display
	w.SetBlipData(b.id, enum.Display, int64(display))
}

func (b *IBlip) SetAlpha(alpha uint32) {
	b.alpha = alpha
	w.SetBlipData(b.id, enum.Alpha, int64(alpha))
}

func (b *IBlip) SetFriendly(friendly bool) {
	b.friendly = friendly
	value := 0
	if friendly {
		value = 1
	}
	w.SetBlipData(b.id, enum.Friendly, int64(value))
}

func (b *IBlip) SetHighDetail(highDetail bool) {
	b.highDetail = highDetail
	value := 0
	if highDetail {
		value = 1
	}
	w.SetBlipData(b.id, enum.HighDetail, int64(value))
}

func (b *IBlip) SetMissionCreator(missionCreator bool) {
	b.missionCreator = missionCreator
	value := 0
	if missionCreator {
		value = 1
	}
	w.SetBlipData(b.id, enum.MissionCreator, int64(value))
}

func (b *IBlip) SetShortRange(shortRange bool) {
	b.shortRange = shortRange
	value := 0
	if shortRange {
		value = 1
	}
	w.SetBlipData(b.id, enum.ShortRange, int64(value))
}

func (b *IBlip) SetBright(bright bool) {
	b.bright = bright
	value := 0
	if bright {
		value = 1
	}
	w.SetBlipData(b.id, enum.Bright, int64(value))
}

func (b *IBlip) SetCrewIndicatorVisible(crewIndicatorVisible bool) {
	b.crewIndicatorVisible = crewIndicatorVisible
	value := 0
	if crewIndicatorVisible {
		value = 1
	}
	w.SetBlipData(b.id, enum.CrewIndicatorVisible, int64(value))
}

func (b *IBlip) SetCategory(category uint32) {
	b.category = category
	w.SetBlipData(b.id, enum.Category, int64(category))
}

func (b *IBlip) SetFlashInterval(flashInterval uint32) {
	b.flashInterval = flashInterval
	w.SetBlipData(b.id, enum.FlashInterval, int64(flashInterval))
}

func (b *IBlip) SetFlashTimer(flashTimer uint32) {
	b.flashTimer = flashTimer
	w.SetBlipData(b.id, enum.FlashTimer, int64(flashTimer))
}

func (b *IBlip) SetFlashes(flashes bool) {
	b.flashes = flashes
	value := 0
	if flashes {
		value = 1
	}
	w.SetBlipData(b.id, enum.Flashes, int64(value))
}

func (b *IBlip) SetFlashesAlternate(flashesAlternate bool) {
	b.flashesAlternate = flashesAlternate
	value := 0
	if flashesAlternate {
		value = 1
	}
	w.SetBlipData(b.id, enum.FlashesAlternate, int64(value))
}

func (b *IBlip) SetGlobal(global bool) {
	b.global = global
	value := 0
	if global {
		value = 1
	}
	w.SetBlipData(b.id, enum.Global, int64(value))
}

func (b *IBlip) SetMinimalOnEdge(minimalOnEdge bool) {
	b.minimalOnEdge = minimalOnEdge
	value := 0
	if minimalOnEdge {
		value = 1
	}
	w.SetBlipData(b.id, enum.MinimalOnEdge, int64(value))
}

func (b *IBlip) SetRoute(route bool) {
	b.route = route
	value := 0
	if route {
		value = 1
	}
	w.SetBlipData(b.id, enum.Route, int64(value))
}

func (b *IBlip) SetPulse(pulse bool) {
	b.pulse = pulse
	value := 0
	if pulse {
		value = 1
	}
	w.SetBlipData(b.id, enum.Pulse, int64(value))
}

func (b *IBlip) SetHiddenOnLegend(hiddenOnLegend bool) {
	b.hiddenOnLegend = hiddenOnLegend
	value := 0
	if hiddenOnLegend {
		value = 1
	}
	w.SetBlipData(b.id, enum.HiddenOnLegend, int64(value))
}

func (b *IBlip) SetOutlineIndicatorVisible(outlineIndicatorVisible bool) {
	b.outlineIndicatorVisible = outlineIndicatorVisible
	value := 0
	if outlineIndicatorVisible {
		value = 1
	}
	w.SetBlipData(b.id, enum.OutlineIndicatorVisible, int64(value))
}

func (b *IBlip) SetRot(rot float32) {
	b.rot = rot
	w.SetBlipData(b.id, enum.BlipRot, int64(math.Float64bits(float64(rot))))
}

func (b *IBlip) SetShrinked(shrinked bool) {
	b.shrinked = shrinked
	value := 0
	if shrinked {
		value = 1
	}
	w.SetBlipData(b.id, enum.Shrinked, int64(value))
}

func (b *IBlip) SetShowCone(showCone bool) {
	b.showCone = showCone
	value := 0
	if showCone {
		value = 1
	}
	w.SetBlipData(b.id, enum.ShowCone, int64(value))
}

func (b *IBlip) SetTickVisible(tickVisible bool) {
	b.tickVisible = tickVisible
	value := 0
	if tickVisible {
		value = 1
	}
	w.SetBlipData(b.id, enum.TickVisible, int64(value))
}

func (b *IBlip) SetUseHeightIndicatorOnEdge(useHeightIndicatorOnEdge bool) {
	b.useHeightIndicatorOnEdge = useHeightIndicatorOnEdge
	value := 0
	if useHeightIndicatorOnEdge {
		value = 1
	}
	w.SetBlipData(b.id, enum.UseHeightIndicatorOnEdge, int64(value))
}

func (b *IBlip) SetPosition(position *entities.Vector3) {
	b.position = position
	w.SetBlipMetaData(b.id, enum.BlipPosition, int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32, "", 0, 0, 0, 0)
}

func (b *IBlip) SetName(name string) {
	b.name = name
	w.SetBlipMetaData(b.id, enum.Name, 0, 0, name, 0, 0, 0, 0)
}

func (b *IBlip) SetRouteColor(rgbaColor *entities.Rgba) {
	b.routeColor = rgbaColor
	w.SetBlipMetaData(b.id, enum.RouteColor, 0, 0, "", rgbaColor.R, rgbaColor.G, rgbaColor.B, rgbaColor.A)
}

func (b *IBlip) SetHeadingIndicatorVisible(headingIndicatorVisible bool) {
	b.headingIndicatorVisible = headingIndicatorVisible
	value := 0
	if headingIndicatorVisible {
		value = 1
	}
	w.SetBlipData(b.id, enum.HeadingIndicatorVisible, int64(value))
}

func (b *IBlip) SetShortHeightThreshold(shortHeightThreshold bool) {
	b.shortHeightThreshold = shortHeightThreshold
	value := 0
	if shortHeightThreshold {
		value = 1
	}
	w.SetBlipData(b.id, enum.ShortHeightThreshold, int64(value))
}

func (b *IBlip) SetNumber(number int32) {
	b.number = number
	w.SetBlipData(b.id, enum.Number, int64(number))
}

func (b *IBlip) SetType(blipType blip_type.BlipType) {
	b.blipType = blipType
	w.SetBlipData(b.id, enum.BlipType, int64(blipType))
}

func (b *IBlip) SetGxtName(gxtName string) {
	b.gxtName = gxtName
	w.SetBlipMetaData(b.id, enum.GxtName, 0, 0, gxtName, 0, 0, 0, 0)
}

func (b *IBlip) SetScale(scale *entities.Vector3) {
	b.scale = scale
	w.SetBlipMetaData(b.id, enum.Scale, int64(math.Float32bits(scale.X))|(int64(math.Float32bits(scale.Y))<<32), 0, "", 0, 0, 0, 0)
}

func (b *IBlip) Destroy() {
	w.SetBlipData(b.id, enum.BlipDestroy, int64(0))
	pools.DestroyBlip(b)
}

func (b *IBlip) SetData(key string, value any) {
	b.datas.Store(key, value)
}

func (b *IBlip) DelData(key string) {
	_, ok := b.datas.Load(key)
	if ok {
		b.datas.Delete(key)
	}
}

func (b *IBlip) DelAllData() {
	b.datas.Range(func(key, value any) bool {
		b.datas.Delete(key)
		return true
	})
}

func (b *IBlip) HasData(key string) bool {
	_, ok := b.datas.Load(key)
	if ok {
		return true
	}
	return false
}

func (b *IBlip) GetData(key string) any {
	value, ok := b.datas.Load(key)
	if ok {
		return value
	}
	return value
}

func (b *IBlip) GetDatas() []any {
	var datas []any
	b.datas.Range(func(key, value any) bool {
		datas = append(datas, key)
		return true
	})
	return datas
}
