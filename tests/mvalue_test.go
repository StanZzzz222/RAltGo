package tests

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common"
	"github.com/StanZzzz222/RAltGo/common/models"
	"testing"
)

/*
   Create by zyx
   Date Time: 2024/9/17
   File: mvalue_test.go
*/

type User struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Likes []string `json:"likes"`
}

func TestMValues(t *testing.T) {
	mvalues := models.NewMValues("test")
	fmt.Println(mvalues.Dump())

	var users []User
	for i := 0; i < 8000; i++ {
		users = append(users, User{Id: i, Name: fmt.Sprintf("test%v", i), Likes: []string{"like1", "like2"}})
	}
	var p = &models.IPlayer{}
	p = p.NewIPlayer(1, "test", "127.0.0.1", "testToken", "test", 123, 123, 456, common.NewVector3(0, 0, 0), common.NewVector3(0, 0, 0))
	for i := 0; i < 3000; i++ {
		p.Emit("test", p, "test", 1, true, 1.1, users)
	}
}
