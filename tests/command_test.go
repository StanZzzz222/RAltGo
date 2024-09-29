package tests

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common"
	"github.com/StanZzzz222/RAltGo/common/command"
	"github.com/StanZzzz222/RAltGo/common/core/alt/alt_events"
	"github.com/StanZzzz222/RAltGo/common/models"
	"strconv"
	"strings"
	"testing"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: command_test.go
*/

func TestCommand(t *testing.T) {
	var p = &models.IPlayer{}
	p = p.NewIPlayer(1, "test", "127.0.0.1", "test", 1, 1, common.NewVector3(0, 0, 0), common.NewVector3(0, 0, 0))
	command.SetCommandErrorCustomHandler(func(player *models.IPlayer, name string, desc string, isParamInsufficient bool) string {
		return "参数不正确, 应为: " + desc
	})
	alt_events.Events().OnCommandError(func(player *models.IPlayer, commandName, desc string) {
		fmt.Println(commandName, " | ", desc)
	})
	group := command.NewCommandGroup("PublicCommands")
	group.UseMiddleware(func(player *models.IPlayer, name string, args []any) bool {
		fmt.Println("Call middleware")
		return true
	})
	{
		group.OnCommandDesc("test", func(player *models.IPlayer, num int64) {
			fmt.Println(num)
		}, false, "/test [num]")
	}
	var params []any
	for _, param := range []string{"12.25"} {
		if strings.Contains(param, ".") {
			if value, err := strconv.ParseFloat(param, 64); err == nil {
				params = append(params, value)
				continue
			}
			params = append(params, param)
			continue
		}
		if value, err := strconv.ParseInt(param, 10, 64); err == nil {
			params = append(params, value)
			continue
		}
		if value, err := strconv.ParseBool(param); err == nil {
			params = append(params, value)
			continue
		}
		params = append(params, param)
	}
	group.TriggerCommand("/test", p, params...)
	command.TriggerLocalCommand("/test", p, int64(64))
}
