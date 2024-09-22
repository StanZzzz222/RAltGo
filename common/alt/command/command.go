package command

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/models"
	"reflect"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: command.go
*/

var groups = &sync.Map{}

type middlewareCallback = func(player *models.IPlayer, name string, args []any) bool
type Group struct {
	name        string
	commands    *sync.Map
	middlewares []middlewareCallback
}
type Command struct {
	name     string
	callback any
	greedy   bool
	desc     string
}

func NewCommandGroup(name string) *Group {
	group := &Group{
		name:        name,
		commands:    &sync.Map{},
		middlewares: make([]middlewareCallback, 0),
	}
	groups.Store(name, group)
	return group
}

func (g *Group) UseMiddleware(callback middlewareCallback) {
	g.middlewares = append(g.middlewares, callback)
}

func (g *Group) OnCommandDesc(name string, callback any, greedy bool, desc string) {
	var commandName string
	if name[0] == '/' {
		commandName = name
	} else {
		commandName = fmt.Sprintf("/%v", name)
	}
	command := &Command{
		name:     commandName,
		callback: callback,
		greedy:   greedy,
		desc:     desc,
	}
	g.commands.Store(name, command)
}

func (g *Group) OnCommand(name string, callback any, greedy bool) {
	var commandName string
	if name[0] == '/' {
		commandName = name
	} else {
		commandName = fmt.Sprintf("/%v", name)
	}
	command := &Command{
		name:     commandName,
		callback: callback,
		greedy:   greedy,
		desc:     "",
	}
	g.commands.Store(name, command)
}

func (g *Group) TriggerCommand(name string, player *models.IPlayer, args ...any) {
	var flag bool
	if len(g.middlewares) > 0 {
		for _, callback := range g.middlewares {
			flag = callback(player, name, args)
			if !flag {
				break
			}
		}
	}
	if flag {
		g.commands.Range(func(key, value any) bool {
			command := value.(*Command)
			if command.name == name {
				if command.greedy {
					triggerGreedyCommand(command, player, args...)
					return false
				}
				triggerCommand(command, player, args...)
				return false
			}
			return true
		})
	}
}

func GetCommandGroups() []*Group {
	var gs []*Group
	groups.Range(func(key, value any) bool {
		gs = append(gs, value.(*Group))
		return true
	})
	return gs
}

func GetCommandGroupByName(name string) *Group {
	if value, ok := groups.Load(name); ok {
		return value.(*Group)
	}
	return nil
}

func triggerCommand(command *Command, player *models.IPlayer, args ...any) {
	callbackValue := reflect.ValueOf(command.callback)
	inputs := make([]reflect.Value, 0)
	inputs = append(inputs, reflect.ValueOf(player))
	for _, arg := range args {
		inputs = append(inputs, reflect.ValueOf(arg))
	}
	callbackValue.Call(inputs)
}

func triggerGreedyCommand(command *Command, player *models.IPlayer, args ...any) {
	combinedArgs := ""
	for _, arg := range args {
		combinedArgs += fmt.Sprintf("%v ", arg)
	}
	callbackValue := reflect.ValueOf(command.callback)
	inputs := make([]reflect.Value, 0)
	inputs = append(inputs, reflect.ValueOf(player), reflect.ValueOf(combinedArgs))
	callbackValue.Call(inputs)
}
