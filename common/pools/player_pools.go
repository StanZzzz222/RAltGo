package pools

import (
	"github.com/StanZzzz222/RAltGo/common/models"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/18
   File: player_pools.go
*/

var playerPools = newPlayerPools()

type PlayerPool struct {
	players *sync.Map
}

func newPlayerPools() *PlayerPool {
	return &PlayerPool{
		players: &sync.Map{},
	}
}

func (p *PlayerPool) Put(player *models.IPlayer) {
	if _, ok := p.players.Load(player.GetId()); !ok {
		p.players.Store(player.GetId(), player)
	}
}

func (p *PlayerPool) Remove(player *models.IPlayer) {
	if _, ok := p.players.Load(player.GetId()); ok {
		p.players.Delete(player.GetId())
	}
}

func GetPlayer(id int) *models.IPlayer {
	player, ok := playerPools.players.Load(id)
	if ok {
		return player.(*models.IPlayer)
	}
	return nil
}

func GetPlayers() []*models.IPlayer {
	var players []*models.IPlayer
	playerPools.players.Range(func(key, value any) bool {
		players = append(players, value.(*models.IPlayer))
		return true
	})
	return players
}

func GetPlayerIterator() <-chan *models.IPlayer {
	playerChan := make(chan *models.IPlayer)
	go func() {
		defer close(playerChan)
		playerPools.players.Range(func(key, value any) bool {
			playerChan <- value.(*models.IPlayer)
			return true
		})
	}()
	return playerChan
}

func GetPlayerPools() *PlayerPool {
	return playerPools
}
