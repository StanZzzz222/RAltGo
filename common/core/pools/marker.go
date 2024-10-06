package pools

import (
	"github.com/StanZzzz222/RAltGo/common/models"
)

/*
   Create by zyx
   Date Time: 2024/9/18
   File: checkpoint.go
*/

func GetMarker(id uint32) *models.IMarker {
	return models.GetPools().GetMarker(id)
}

func GetMarkers() []*models.IMarker {
	var markers []*models.IMarker
	var pools = models.GetPools().GetMarkerPools()
	pools.Range(func(key, value any) bool {
		markers = append(markers, value.(*models.IMarker))
		return true
	})
	return markers
}

func GetMarkerIterator() <-chan *models.IMarker {
	var markerChan = make(chan *models.IMarker)
	var pools = models.GetPools().GetMarkerPools()
	go func() {
		defer close(markerChan)
		pools.Range(func(key, value any) bool {
			markerChan <- value.(*models.IMarker)
			return true
		})
	}()
	return markerChan
}
