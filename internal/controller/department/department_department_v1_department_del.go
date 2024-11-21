package department

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_department"

	"forest_resources_management_system_gf/api/department/department_v1"
)

func (c *ControllerDepartment_v1) DepartmentDel(ctx context.Context, req *department_v1.DepartmentDelReq) (res *department_v1.DepartmentDelRes, err error) {
	return l_department.NewDepartmentLogic().DepartmentDel(ctx, req)
}
