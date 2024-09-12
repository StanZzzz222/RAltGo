package models

import (
	"github.com/StanZzzz222/RAltGo/common/utils"
	"github.com/StanZzzz222/RAltGo/enums/blip_type"
	"github.com/StanZzzz222/RAltGo/internal/entitys"
	"github.com/StanZzzz222/RAltGo/internal/enum"
	"math"
)

/*
   Create by zyx
   Date Time: 2024/9/12
   File: blip.go
*/

type IBlip struct {
	id                       uint32
	blipType                 uint32
	color                    uint32
	spriteId                 uint32
	alpha                    uint32
	category                 uint32
	flashInterval            uint32
	flashTimer               uint32
	rot                      float32
	visible                  bool
	display                  bool
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
	name                     string
	routeColor               *utils.Rgba
	rgbaColor                *utils.Rgba
	position                 *entitys.Vector3
}

func (b *IBlip) GetId() uint32                     { return b.id }
func (b *IBlip) GetName() string                   { return b.name }
func (b *IBlip) GetBlipType() blip_type.BlipType   { return blip_type.BlipType(b.blipType) }
func (b *IBlip) GetBlipColor() uint32              { return b.color }
func (b *IBlip) GetSpriteId() uint32               { return b.spriteId }
func (b *IBlip) GetAlpha() uint32                  { return b.alpha }
func (b *IBlip) GetCategory() uint32               { return b.category }
func (b *IBlip) GetFlashInterval() uint32          { return b.flashInterval }
func (b *IBlip) GetFlashTimer() uint32             { return b.flashTimer }
func (b *IBlip) GetRot() float32                   { return b.rot }
func (b *IBlip) GetVisible() bool                  { return b.visible }
func (b *IBlip) GetDisplay() bool                  { return b.display }
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
func (b *IBlip) GetRouteColor() *utils.Rgba        { return b.routeColor }
func (b *IBlip) GetRgbaColor() *utils.Rgba         { return b.rgbaColor }
func (b *IBlip) GetPosition() *entitys.Vector3     { return b.position }

func (b *IBlip) NewIBlip(id, blipType, spriteId, color uint32, rot float32, position *entitys.Vector3) *IBlip {
	return &IBlip{
		id:       id,
		blipType: blipType,
		spriteId: spriteId,
		color:    color,
		rot:      rot,
		position: position,
	}
}

func (b *IBlip) SetBlipSprite(spriteId uint32) {
	b.spriteId = spriteId
	w.SetBlipData(b.id, enum.Sprite, int64(spriteId))
}

func (b *IBlip) SetBlipColor(color uint32) {
	b.color = color
	w.SetBlipData(b.id, enum.Color, int64(color))
}

func (b *IBlip) SetBlipRGBA(rgbaColor *utils.Rgba) {
	b.rgbaColor = rgbaColor
	w.SetBlipMetaData(b.id, enum.RgbaColor, 0, 0, "", rgbaColor.R, rgbaColor.G, rgbaColor.B, rgbaColor.A)
}

func (b *IBlip) SetBlipVisible(visible bool) {
	b.visible = visible
	value := 0
	if visible {
		value = 1
	}
	w.SetBlipData(b.id, enum.BlipVisible, int64(value))
}

func (b *IBlip) SetBlipDisplay(display bool) {
	b.display = display
	value := 0
	if display {
		value = 1
	}
	w.SetBlipData(b.id, enum.Display, int64(value))
}

func (b *IBlip) SetBlipAlpha(alpha uint32) {
	b.alpha = alpha
	w.SetBlipData(b.id, enum.Alpha, int64(alpha))
}

func (b *IBlip) SetBlipFriendly(friendly bool) {
	b.friendly = friendly
	value := 0
	if friendly {
		value = 1
	}
	w.SetBlipData(b.id, enum.Friendly, int64(value))
}

func (b *IBlip) SetBlipHighDetail(highDetail bool) {
	b.highDetail = highDetail
	value := 0
	if highDetail {
		value = 1
	}
	w.SetBlipData(b.id, enum.HighDetail, int64(value))
}

func (b *IBlip) SetBlipMissionCreator(missionCreator bool) {
	b.missionCreator = missionCreator
	value := 0
	if missionCreator {
		value = 1
	}
	w.SetBlipData(b.id, enum.MissionCreator, int64(value))
}

func (b *IBlip) SetBlipShortRange(shortRange bool) {
	b.shortRange = shortRange
	value := 0
	if shortRange {
		value = 1
	}
	w.SetBlipData(b.id, enum.ShortRange, int64(value))
}

func (b *IBlip) SetBlipBright(bright bool) {
	b.bright = bright
	value := 0
	if bright {
		value = 1
	}
	w.SetBlipData(b.id, enum.Bright, int64(value))
}

func (b *IBlip) SetBlipCrewIndicatorVisible(crewIndicatorVisible bool) {
	b.crewIndicatorVisible = crewIndicatorVisible
	value := 0
	if crewIndicatorVisible {
		value = 1
	}
	w.SetBlipData(b.id, enum.CrewIndicatorVisible, int64(value))
}

func (b *IBlip) SetBlipCategory(category uint32) {
	b.category = category
	w.SetBlipData(b.id, enum.Category, int64(category))
}

func (b *IBlip) SetBlipFlashInterval(flashInterval uint32) {
	b.flashInterval = flashInterval
	w.SetBlipData(b.id, enum.FlashInterval, int64(flashInterval))
}

func (b *IBlip) SetBlipFlashTimer(flashTimer uint32) {
	b.flashTimer = flashTimer
	w.SetBlipData(b.id, enum.FlashTimer, int64(flashTimer))
}

func (b *IBlip) SetBlipFlashes(flashes bool) {
	b.flashes = flashes
	value := 0
	if flashes {
		value = 1
	}
	w.SetBlipData(b.id, enum.Flashes, int64(value))
}

func (b *IBlip) SetBlipFlashesAlternate(flashesAlternate bool) {
	b.flashesAlternate = flashesAlternate
	value := 0
	if flashesAlternate {
		value = 1
	}
	w.SetBlipData(b.id, enum.FlashesAlternate, int64(value))
}

func (b *IBlip) SetBlipGlobal(global bool) {
	b.global = global
	value := 0
	if global {
		value = 1
	}
	w.SetBlipData(b.id, enum.Global, int64(value))
}

func (b *IBlip) SetBlipMinimalOnEdge(minimalOnEdge bool) {
	b.minimalOnEdge = minimalOnEdge
	value := 0
	if minimalOnEdge {
		value = 1
	}
	w.SetBlipData(b.id, enum.MinimalOnEdge, int64(value))
}

func (b *IBlip) SetBlipRoute(route bool) {
	b.route = route
	value := 0
	if route {
		value = 1
	}
	w.SetBlipData(b.id, enum.Route, int64(value))
}

func (b *IBlip) SetBlipPulse(pulse bool) {
	b.pulse = pulse
	value := 0
	if pulse {
		value = 1
	}
	w.SetBlipData(b.id, enum.Pulse, int64(value))
}

func (b *IBlip) SetBlipHiddenOnLegend(hiddenOnLegend bool) {
	b.hiddenOnLegend = hiddenOnLegend
	value := 0
	if hiddenOnLegend {
		value = 1
	}
	w.SetBlipData(b.id, enum.HiddenOnLegend, int64(value))
}

func (b *IBlip) SetBlipOutlineIndicatorVisible(outlineIndicatorVisible bool) {
	b.outlineIndicatorVisible = outlineIndicatorVisible
	value := 0
	if outlineIndicatorVisible {
		value = 1
	}
	w.SetBlipData(b.id, enum.OutlineIndicatorVisible, int64(value))
}

func (b *IBlip) SetBlipRot(rot float32) {
	b.rot = rot
	w.SetBlipData(b.id, enum.BlipRot, int64(math.Float64bits(float64(rot))))
}

func (b *IBlip) SetBlipShrinked(shrinked bool) {
	b.shrinked = shrinked
	value := 0
	if shrinked {
		value = 1
	}
	w.SetBlipData(b.id, enum.Shrinked, int64(value))
}

func (b *IBlip) SetBlipShowCone(showCone bool) {
	b.showCone = showCone
	value := 0
	if showCone {
		value = 1
	}
	w.SetBlipData(b.id, enum.ShowCone, int64(value))
}

func (b *IBlip) SetBlipTickVisible(tickVisible bool) {
	b.tickVisible = tickVisible
	value := 0
	if tickVisible {
		value = 1
	}
	w.SetBlipData(b.id, enum.TickVisible, int64(value))
}

func (b *IBlip) SetBlipUseHeightIndicatorOnEdge(useHeightIndicatorOnEdge bool) {
	b.useHeightIndicatorOnEdge = useHeightIndicatorOnEdge
	value := 0
	if useHeightIndicatorOnEdge {
		value = 1
	}
	w.SetBlipData(b.id, enum.UseHeightIndicatorOnEdge, int64(value))
}

func (b *IBlip) SetBlipPosition(position *entitys.Vector3) {
	b.position = position
	w.SetBlipMetaData(b.id, enum.BlipPosition, int64(math.Float32bits(position.X))|(int64(math.Float32bits(position.Y))<<32), uint64(math.Float32bits(position.Z))<<32, "", 0, 0, 0, 0)
}

func (b *IBlip) SetBlipName(name string) {
	b.name = name
	w.SetBlipMetaData(b.id, enum.Name, 0, 0, name, 0, 0, 0, 0)
}

func (b *IBlip) SetBlipRouteColor(rgbaColor *utils.Rgba) {
	b.routeColor = rgbaColor
	w.SetBlipMetaData(b.id, enum.RouteColor, 0, 0, "", rgbaColor.R, rgbaColor.G, rgbaColor.B, rgbaColor.A)
}
