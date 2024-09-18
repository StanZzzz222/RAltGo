package pools

/*
   Create by zyx
   Date Time: 2024/9/18
   File: blip_pools.go
*/

func GetBlip(id uint32) *IBlip {
	blip, ok := blipPools.blips.Load(id)
	if ok {
		return blip.(*IBlip)
	}
	return nil
}

func GetBlips() []*IBlip {
	var blips []*IBlip
	blipPools.blips.Range(func(key, value any) bool {
		blips = append(blips, value.(*IBlip))
		return true
	})
	return blips
}

func GetBlipIterator() <-chan *IBlip {
	blipChan := make(chan *IBlip)
	go func() {
		defer close(blipChan)
		blipPools.blips.Range(func(key, value any) bool {
			blipChan <- value.(*IBlip)
			return true
		})
	}()
	return blipChan
}
