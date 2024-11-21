// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package department

import (
	"context"

	"forest_resources_management_system_gf/api/department/department_v1"
)

type IDepartmentDepartment_v1 interface {
	DepartmentCreate(ctx context.Context, req *department_v1.DepartmentCreateReq) (res *department_v1.DepartmentCreateRes, err error)
	DepartmentDel(ctx context.Context, req *department_v1.DepartmentDelReq) (res *department_v1.DepartmentDelRes, err error)
	DepartmentList(ctx context.Context, req *department_v1.DepartmentListReq) (res *department_v1.DepartmentListRes, err error)
	DepartmentModify(ctx context.Context, req *department_v1.DepartmentModifyReq) (res *department_v1.DepartmentModifyRes, err error)
}
