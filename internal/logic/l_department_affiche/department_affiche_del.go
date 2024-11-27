package l_department_affiche

import (
	"context"
	"forest_resources_management_system_gf/api/department_affiche/department_affiche_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultDepartmentAdapter) DepartmentAfficheDel(ctx context.Context, req *department_affiche_v1.DepartmentAfficheDelReq) (res *department_affiche_v1.DepartmentAfficheDelRes, err error) {
	var departmentAffiche *do.DepartmentAffiche
	common.CopyFields(req, &departmentAffiche)
	if _, err = dao.DepartmentAffiche.Ctx(ctx).Delete(departmentAffiche); err != nil {
		g.Log().Errorf(ctx, "delete department affiche failed, err:%v", err)
		return nil, common.NewError(consts.InternalServerError, "delete department affiche failed")
	}
	return
}
