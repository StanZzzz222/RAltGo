package tests

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common"
	"github.com/StanZzzz222/RAltGo/common/core/alt"
	"github.com/StanZzzz222/RAltGo/common/core/pools"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/hash_enums"
	"testing"
)

/*
   Create by zyx
   Date Time: 2024/9/29
   File: any_test.go
*/

func TestAny(t *testing.T) {
	var p = &models.IPlayer{}
	p = p.NewIPlayer(1, "test", "127.0.0.1", "test", 1, 1, common.NewVector3(0, 0, 0), common.NewVector3(0, 0, 0))
	models.GetPools().PutPlayer(p)
	getP := pools.GetAnyEntity[models.IPlayer](p.GetId())
	fmt.Println(getP.GetName())
	getPlayerPools := pools.GetAnyEntitys[models.IPlayer]()
	for _, target := range getPlayerPools {
		fmt.Println(target.GetName())
	}
	for target := range pools.GetAnyEntityIterator[models.IPlayer]() {
		fmt.Println(target.GetName())
	}
	entites := alt.GetEntitiesInDimension[models.IPlayer](hash_enums.DefaultDimension)
	fmt.Println(len(entites))
}
