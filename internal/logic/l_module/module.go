package l_module

import (
	"forest_resources_management_system_gf/api/module"
)

type defaultModule struct{}

type IModuleLogic interface {
	module.IModuleModule_v1
}

func NewModuleLogic() IModuleLogic {
	return &defaultModule{}
}
