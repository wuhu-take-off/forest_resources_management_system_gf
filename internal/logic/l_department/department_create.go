package l_department

import (
	"context"
	"forest_resources_management_system_gf/api/department/department_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultDepartment) DepartmentCreate(ctx context.Context, req *department_v1.DepartmentCreateReq) (res *department_v1.DepartmentCreateRes, err error) {
	var department *do.Department
	common.CopyFields(req, &department)
	if _, err = dao.Department.Ctx(ctx).Insert(department); err != nil {
		g.Log().Errorf(ctx, "create department failed, err: %v", err)
		return nil, common.NewError(consts.InternalServerError, "新建部门失败")
	}

	return
}
