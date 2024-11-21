package department

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_department"

	"forest_resources_management_system_gf/api/department/department_v1"
)

func (c *ControllerDepartment_v1) DepartmentModify(ctx context.Context, req *department_v1.DepartmentModifyReq) (res *department_v1.DepartmentModifyRes, err error) {
	return l_department.NewDepartmentLogic().DepartmentModify(ctx, req)
}
