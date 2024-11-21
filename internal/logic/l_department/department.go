package l_department

import (
	"context"
	"forest_resources_management_system_gf/api/department/department_v1"
)

type defaultDepartment struct{}

type IDepartmentLogic interface {
	DepartmentCreate(ctx context.Context, req *department_v1.DepartmentCreateReq) (res *department_v1.DepartmentCreateRes, err error)
	DepartmentDel(ctx context.Context, req *department_v1.DepartmentDelReq) (res *department_v1.DepartmentDelRes, err error)
	DepartmentList(ctx context.Context, req *department_v1.DepartmentListReq) (res *department_v1.DepartmentListRes, err error)
	DepartmentModify(ctx context.Context, req *department_v1.DepartmentModifyReq) (res *department_v1.DepartmentModifyRes, err error)
}

func NewDepartmentLogic() IDepartmentLogic {
	return &defaultDepartment{}
}
