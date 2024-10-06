package command

import (
	"fmt"
	"github.com/StanZzzz222/RAltGo/common/core/alt/alt_events"
	"github.com/StanZzzz222/RAltGo/common/models"
	"github.com/StanZzzz222/RAltGo/logger"
	"reflect"
	"sync"
)

/*
   Create by zyx
   Date Time: 2024/9/22
   File: command.go
*/

var groups = &sync.Map{}
var customHandler customCommandError
var errorMessage = "Command parameter error"

type middlewareCallback = func(player *models.IPlayer, name string, args []any) bool
type customCommandError = func(player *models.IPlayer, name string, desc string, isParamInsufficient bool) string
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
type ExportCommand struct {
	name   string
	greedy bool
	desc   string
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

func SetCommandErrorCustomHandler(handler customCommandError) {
	customHandler = handler
}

func SetCommandErrorMessage(message string) {
	errorMessage = message
}

func TriggerLocalCommand(name string, args ...any) {
	if name[0] == '/' {
		name = name[1:]
	}
	if !checkZeroEventArgs(args) {
		logger.Logger().LogError("TriggerLocalCommand: should not be zero parameters")
		return
	}
	if !checkFirstEventArgs(args[0]) {
		logger.Logger().LogError("TriggerLocalCommand: The first parameter should be *models.IPlayer")
		return
	}
	groups.Range(func(key, value any) bool {
		flag := false
		group := value.(*Group)
		group.commands.Range(func(key, value any) bool {
			targetName := key.(string)
			if targetName == name {
				flag = true
				command := value.(*Command)
				if command.greedy {
					triggerGreedyCommand(command, args[0].(*models.IPlayer), args[1:]...)
					return false
				}
				triggerCommand(command, args[0].(*models.IPlayer), args[1:]...)
				return false
			}
			return true
		})
		return flag
	})
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
	if command, ok := g.getCommand(name); ok {
		var res = false
		if len(g.middlewares) <= 0 {
			res = true
		} else {
			for _, callback := range g.middlewares {
				res = callback(player, name, args)
				if !res {
					break
				}
			}
		}
		if res {
			if command.greedy {
				triggerGreedyCommand(command, player, args...)
				return
			}
			triggerCommand(command, player, args...)
		}
	} else {
		alt_events.Triggers().TriggerOnCommandError(player, false, name, "")
	}
}

func (g *Group) ExportCommands() []*ExportCommand {
	var exportCommands []*ExportCommand
	g.commands.Range(func(key, value any) bool {
		command := value.(*Command)
		exportCommands = append(exportCommands, &ExportCommand{
			name:   command.name,
			greedy: command.greedy,
			desc:   command.desc,
		})
		return true
	})
	return exportCommands
}

func (g *Group) getCommand(name string) (*Command, bool) {
	var res = false
	var resCommand *Command
	g.commands.Range(func(key, value any) bool {
		command := value.(*Command)
		if command.name == name {
			res = true
			resCommand = command
			return false
		}
		return true
	})
	return resCommand, res
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
	callbackType := reflect.TypeOf(command.callback)
	if callbackType.NumIn() != len(args)+1 {
		if customHandler != nil {
			var desc = ""
			if len(command.desc) > 0 {
				desc = command.desc
			}
			commandErrorMessage := customHandler(player, command.name, desc, true)
			alt_events.Triggers().TriggerOnCommandError(player, true, command.name, commandErrorMessage)
			return
		}
		if len(command.desc) <= 0 {
			alt_events.Triggers().TriggerOnCommandError(player, true, command.name, errorMessage)
			return
		}
		alt_events.Triggers().TriggerOnCommandError(player, true, command.name, command.desc)
		return
	}
	for i := 1; i < callbackType.NumIn(); i++ {
		argType := callbackType.In(i)
		targetType := reflect.TypeOf(args[i-1])
		if argType.Kind() != targetType.Kind() {
			if customHandler != nil {
				var desc = ""
				if len(command.desc) > 0 {
					desc = command.desc
				}
				commandErrorMessage := customHandler(player, command.name, desc, false)
				alt_events.Triggers().TriggerOnCommandError(player, true, command.name, commandErrorMessage)
				return
			}
			if len(command.desc) <= 0 {
				alt_events.Triggers().TriggerOnCommandError(player, true, command.name, errorMessage)
				return
			}
			alt_events.Triggers().TriggerOnCommandError(player, true, command.name, command.desc)
			return
		}
	}
	inputs := make([]reflect.Value, 0)
	inputs = append(inputs, reflect.ValueOf(player))
	if len(args) == 0 {
		callbackValue.Call(inputs)
		return
	}
	for _, arg := range args {
		inputs = append(inputs, reflect.ValueOf(arg))
	}
	callbackValue.Call(inputs)
}

func triggerGreedyCommand(command *Command, player *models.IPlayer, args ...any) {
	callbackValue := reflect.ValueOf(command.callback)
	callbackType := reflect.TypeOf(command.callback)
	if callbackType.NumIn() != len(args)+1 {
		if customHandler != nil {
			var desc = ""
			if len(command.desc) > 0 {
				desc = command.desc
			}
			commandErrorMessage := customHandler(player, command.name, desc, true)
			alt_events.Triggers().TriggerOnCommandError(player, true, command.name, commandErrorMessage)
			return
		}
		if len(command.desc) <= 0 {
			alt_events.Triggers().TriggerOnCommandError(player, true, command.name, errorMessage)
			return
		}
		alt_events.Triggers().TriggerOnCommandError(player, true, command.name, command.desc)
		return
	}
	for i := 1; i < callbackType.NumIn(); i++ {
		argType := callbackType.In(i)
		targetType := reflect.TypeOf(args[i-1])
		if argType.Kind() != targetType.Kind() {
			if customHandler != nil {
				var desc = ""
				if len(command.desc) > 0 {
					desc = command.desc
				}
				commandErrorMessage := customHandler(player, command.name, desc, false)
				alt_events.Triggers().TriggerOnCommandError(player, true, command.name, commandErrorMessage)
				return
			}
			if len(command.desc) <= 0 {
				alt_events.Triggers().TriggerOnCommandError(player, true, command.name, errorMessage)
				return
			}
			alt_events.Triggers().TriggerOnCommandError(player, true, command.name, command.desc)
			return
		}
	}
	combinedArgs := ""
	if len(args) > 0 {
		for _, arg := range args {
			combinedArgs += fmt.Sprintf("%v ", arg)
		}
	}
	inputs := make([]reflect.Value, 0)
	inputs = append(inputs, reflect.ValueOf(player), reflect.ValueOf(combinedArgs))
	callbackValue.Call(inputs)
}

func checkZeroEventArgs(args ...any) bool {
	if arr, ok := args[0].([]any); ok && len(arr) == 0 {
		return false
	}
	return len(args) != 0
}

func checkFirstEventArgs(arg any) bool {
	var paramType = reflect.TypeOf(arg)
	if paramType.Kind() == reflect.Ptr {
		elemType := paramType.Elem()
		if elemType == reflect.TypeOf((*models.IPlayer)(nil)).Elem() {
			return true
		}
	}
	return false
}
