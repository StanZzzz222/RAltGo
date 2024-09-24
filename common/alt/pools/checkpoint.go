package pools

import "github.com/StanZzzz222/RAltGo/common/models"

/*
   Create by zyx
   Date Time: 2024/9/18
   File: checkpoint.go
*/

func GetCheckpoint(id uint32) *models.ICheckpoint {
	return models.GetPools().GetCheckpoint(id)
}

func GetCheckpoints() []*models.ICheckpoint {
	var checkpoints []*models.ICheckpoint
	var pools = models.GetPools().GetCheckpointPools()
	pools.Range(func(key, value any) bool {
		checkpoints = append(checkpoints, value.(*models.ICheckpoint))
		return true
	})
	return checkpoints
}

func GetCheckpointIterator() <-chan *models.ICheckpoint {
	var checkpointChan = make(chan *models.ICheckpoint)
	var pools = models.GetPools().GetCheckpointPools()
	go func() {
		defer close(checkpointChan)
		pools.Range(func(key, value any) bool {
			checkpointChan <- value.(*models.ICheckpoint)
			return true
		})
	}()
	return checkpointChan
}
