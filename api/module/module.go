// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package module

import (
	"context"

	"forest_resources_management_system_gf/api/module/module_v1"
)

type IModuleModule_v1 interface {
	ModuleList(ctx context.Context, req *module_v1.ModuleListReq) (res *module_v1.ModuleListRes, err error)
}
