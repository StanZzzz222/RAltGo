package pools

import "github.com/StanZzzz222/RAltGo/common/models"

/*
   Create by zyx
   Date Time: 2024/9/18
   File: player.go
*/

func GetPlayer(id uint32) *models.IPlayer {
	return models.GetPools().GetPlayer(id)
}

func GetPlayers() []*models.IPlayer {
	var players []*models.IPlayer
	var pools = models.GetPools().GetPlayerPools()
	pools.Range(func(key, value any) bool {
		players = append(players, value.(*models.IPlayer))
		return true
	})
	return players
}

func GetPlayerIterator() <-chan *models.IPlayer {
	var playerChan = make(chan *models.IPlayer)
	var pools = models.GetPools().GetPlayerPools()
	go func() {
		defer close(playerChan)
		pools.Range(func(key, value any) bool {
			playerChan <- value.(*models.IPlayer)
			return true
		})
	}()
	return playerChan
}
