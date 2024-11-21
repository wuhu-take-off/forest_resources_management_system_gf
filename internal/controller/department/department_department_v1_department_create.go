package department

import (
	"context"
	"forest_resources_management_system_gf/api/department/department_v1"
	"forest_resources_management_system_gf/internal/logic/l_department"
)

func (c *ControllerDepartment_v1) DepartmentCreate(ctx context.Context, req *department_v1.DepartmentCreateReq) (res *department_v1.DepartmentCreateRes, err error) {
	return l_department.NewDepartmentLogic().DepartmentCreate(ctx, req)
}
