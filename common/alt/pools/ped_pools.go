package pools

/*
   Create by zyx
   Date Time: 2024/9/18
   File: ped_pools.go
*/

func GetPed(id uint32) *IPed {
	ped, ok := pedPools.peds.Load(id)
	if ok {
		return ped.(*IPed)
	}
	return nil
}

func GetPeds() []*IPed {
	var peds []*IPed
	pedPools.peds.Range(func(key, value any) bool {
		peds = append(peds, value.(*IPed))
		return true
	})
	return peds
}

func GetPedsIterator() <-chan *IPed {
	pedsChan := make(chan *IPed)
	go func() {
		defer close(pedsChan)
		playerPools.players.Range(func(key, value any) bool {
			pedsChan <- value.(*IPed)
			return true
		})
	}()
	return pedsChan
}
