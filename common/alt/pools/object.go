package pools

import "github.com/StanZzzz222/RAltGo/common/models"

/*
   Create by zyx
   Date Time: 2024/9/18
   File: checkpoint.go
*/

func GetObject(id uint32) *models.IObject {
	return models.GetPools().GetObject(id)
}

func GetObjectss() []*models.IObject {
	var objects []*models.IObject
	var pools = models.GetPools().GetObjectPools()
	pools.Range(func(key, value any) bool {
		objects = append(objects, value.(*models.IObject))
		return true
	})
	return objects
}

func GetObjectIterator() <-chan *models.IObject {
	var objectChan = make(chan *models.IObject)
	var pools = models.GetPools().GetObjectPools()
	go func() {
		defer close(objectChan)
		pools.Range(func(key, value any) bool {
			objectChan <- value.(*models.IObject)
			return true
		})
	}()
	return objectChan
}
