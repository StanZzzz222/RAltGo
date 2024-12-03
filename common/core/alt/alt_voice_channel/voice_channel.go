package alt_checkpoint

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/internal/entities"
	"github.com/StanZzzz222/RAltGo/internal/lib"
)

/*
   Create by zyx
   Date Time: 2024/10/10
   File: voice_channel.go
*/

func CreateVoiceChannel(maxDistance float32) *models.IVoiceChannel {
	var w = lib.GetWrapper()
	var v = &models.IVoiceChannel{}
	ret, freePtrFunc := w.CreateVoiceChannel(false, maxDistance)
	cVoiceChannel := entities.ConvertCVoiceChannel(ret)
	if cVoiceChannel != nil {
		freePtrFunc()
		v = v.NewIVoiceChannel(cVoiceChannel.ID, cVoiceChannel.Spatial, cVoiceChannel.MaxDistance)
		pools := models.GetPools()
		pools.PutVoiceChannel(v)
		return v
	}
	return nil
}

func CreateVoiceChannelSpatial() *models.IVoiceChannel {
	var w = lib.GetWrapper()
	var v = &models.IVoiceChannel{}
	ret, freePtrFunc := w.CreateVoiceChannel(true, 0)
	cVoiceChannel := entities.ConvertCVoiceChannel(ret)
	if cVoiceChannel != nil {
		freePtrFunc()
		v = v.NewIVoiceChannel(cVoiceChannel.ID, cVoiceChannel.Spatial, cVoiceChannel.MaxDistance)
		pools := models.GetPools()
		pools.PutVoiceChannel(v)
		return v
	}
	return nil
}
