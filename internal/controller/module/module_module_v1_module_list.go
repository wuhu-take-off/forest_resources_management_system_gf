package module

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_module"

	"forest_resources_management_system_gf/api/module/module_v1"
)

func (c *ControllerModule_v1) ModuleList(ctx context.Context, req *module_v1.ModuleListReq) (res *module_v1.ModuleListRes, err error) {
	return l_module.NewModuleLogic().ModuleList(ctx, req)
}
