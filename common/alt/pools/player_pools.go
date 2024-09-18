package pools

/*
   Create by zyx
   Date Time: 2024/9/18
   File: player_pools.go
*/

func GetPlayer(id uint32) *IPlayer {
	player, ok := playerPools.players.Load(id)
	if ok {
		return player.(*IPlayer)
	}
	return nil
}

func GetPlayers() []*IPlayer {
	var players []*IPlayer
	playerPools.players.Range(func(key, value any) bool {
		players = append(players, value.(*IPlayer))
		return true
	})
	return players
}

func GetPlayerIterator() <-chan *IPlayer {
	playerChan := make(chan *IPlayer)
	go func() {
		defer close(playerChan)
		playerPools.players.Range(func(key, value any) bool {
			playerChan <- value.(*IPlayer)
			return true
		})
	}()
	return playerChan
}
