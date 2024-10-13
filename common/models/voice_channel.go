package models

import (
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/enums"
	"github.com/StanZzzz222/RAltGo/internal/lib"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/19
   File: colshape.go
*/

type IVoiceChannel struct {
	id          uint32
	spatial     bool
	maxDistance float32
	players     *sync.Map
	warpper     *lib.Warpper
	*NetworkData
}

func (v *IVoiceChannel) GetId() uint32           { return v.id }
func (v *IVoiceChannel) GetSpatial() bool        { return v.spatial }
func (v *IVoiceChannel) GetMaxDistance() float32 { return v.maxDistance }
func (v *IVoiceChannel) GetFilter() uint32 {
	ret, freeDataResultFunc := v.warpper.GetData(v.id, enums.VoiceChannel, uint8(enums.VoiceChannelFilter))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return cDataResult.U32Val
	}
	return 0
}
func (v *IVoiceChannel) GetPriority() int32 {
	ret, freeDataResultFunc := v.warpper.GetData(v.id, enums.VoiceChannel, uint8(enums.VoiceChannelPriority))
	cDataResult := entities.ConverCDataResult(ret)
	if cDataResult != nil {
		freeDataResultFunc()
		return int32(int64(cDataResult.U64Val))
	}
	return 0
}

func (v *IVoiceChannel) NewIVoiceChannel(id uint32, spatial bool, maxDistance float32) *IVoiceChannel {
	return &IVoiceChannel{
		id:          id,
		spatial:     spatial,
		maxDistance: maxDistance,
		players:     &sync.Map{},
		warpper:     lib.GetWarpper(),
		NetworkData: NewNetworkData(id, enums.VoiceChannel),
	}
}

func (v *IVoiceChannel) AddPlayer(player *IPlayer) {
	if _, ok := v.players.Load(player.GetId()); !ok {
		v.players.Store(player.GetId(), player)
		v.warpper.SetVoiceChannelData(v.id, enums.VoiceChannelAddPlayer, int64(player.GetId()))
	}
}

func (v *IVoiceChannel) RemovePlayer(player *IPlayer) {
	if _, ok := v.players.Load(player.GetId()); ok {
		v.players.Delete(player.GetId())
		v.warpper.SetVoiceChannelData(v.id, enums.VoiceChannelRemovePlayer, int64(player.GetId()))
	}
}

func (v *IVoiceChannel) MutePlayer(player *IPlayer) {
	if _, ok := v.players.Load(player.GetId()); ok {
		v.warpper.SetVoiceChannelData(v.id, enums.VoiceChannelMutePlayer, int64(player.GetId()))
	}
}

func (v *IVoiceChannel) UnmutePlayer(player *IPlayer) {
	if _, ok := v.players.Load(player.GetId()); ok {
		v.warpper.SetVoiceChannelData(v.id, enums.VoiceChannelUnmutePlayer, int64(player.GetId()))
	}
}

func (v *IVoiceChannel) SetFilter(filter uint32) {
	v.warpper.SetVoiceChannelData(v.id, enums.VoiceChannelFilter, int64(filter))
}

func (v *IVoiceChannel) SetPriority(priority uint32) {
	v.warpper.SetVoiceChannelData(v.id, enums.VoiceChannelPriority, int64(priority))
}

func (v *IVoiceChannel) HasPlayer(player *IPlayer) bool {
	_, ok := v.players.Load(player.GetId())
	return ok
}

func (v *IVoiceChannel) IsPlayerMuted(player *IPlayer) bool {
	if _, ok := v.players.Load(player.GetId()); ok {
		ret, freeDataResultFunc := v.warpper.GetMetaData(v.id, enums.VoiceChannel, uint8(enums.VoiceChannelIsPlayerMuted), int64(player.GetId()))
		cDataResult := entities.ConverCDataResult(ret)
		if cDataResult != nil {
			freeDataResultFunc()
			return cDataResult.BoolVal
		}
		return false
	}
	return false
}

func (v *IVoiceChannel) Destroy() {
	v.warpper.SetVoiceChannelData(v.id, enums.VoiceChannelDestroy, 0)
	pools.DestroyVoiceChannel(v)
}
