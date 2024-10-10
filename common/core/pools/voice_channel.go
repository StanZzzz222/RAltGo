package pools

import (
	"github.com/StanZzzz222/RAltGo/common/models"
)

/*
   Create by zyx
   Date Time: 2024/10/10
   File: voice_channel.go
*/

func GetVoiceChannel(id uint32) *models.IVoiceChannel {
	return models.GetPools().GetVoiceChannel(id)
}

func GetVoiceChannels() []*models.IVoiceChannel {
	var voiceChannels []*models.IVoiceChannel
	var pools = models.GetPools().GetVoiceChannelPools()
	pools.Range(func(key, value any) bool {
		voiceChannels = append(voiceChannels, value.(*models.IVoiceChannel))
		return true
	})
	return voiceChannels
}

func GetVoiceChannelIterator() <-chan *models.IVoiceChannel {
	var voiceChannelChan = make(chan *models.IVoiceChannel)
	var pools = models.GetPools().GetVoiceChannelPools()
	go func() {
		defer close(voiceChannelChan)
		pools.Range(func(key, value any) bool {
			voiceChannelChan <- value.(*models.IVoiceChannel)
			return true
		})
	}()
	return voiceChannelChan
}
