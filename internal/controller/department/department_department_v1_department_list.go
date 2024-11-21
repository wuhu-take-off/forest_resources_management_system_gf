package department

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_department"

	"forest_resources_management_system_gf/api/department/department_v1"
)

func (c *ControllerDepartment_v1) DepartmentList(ctx context.Context, req *department_v1.DepartmentListReq) (res *department_v1.DepartmentListRes, err error) {
	return l_department.NewDepartmentLogic().DepartmentList(ctx, req)
}
