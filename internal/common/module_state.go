package common

/*
   Create by zyx
   Date Time: 2024/9/10
   File: module_state.go
*/

type ModuleTick struct{}

var moduleTickDone bool

func (m *ModuleTick) ModuleTickDone() {
	moduleTickDone = true
}

func (m *ModuleTick) GetModuleTickDone() bool {
	return moduleTickDone
}
