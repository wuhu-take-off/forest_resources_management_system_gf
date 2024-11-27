package l_department_affiche

import (
	"context"
	"forest_resources_management_system_gf/api/department_affiche/department_affiche_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func (d defaultDepartmentAdapter) DepartmentAfficheModify(ctx context.Context, req *department_affiche_v1.DepartmentAfficheModifyReq) (res *department_affiche_v1.DepartmentAfficheModifyRes, err error) {
	var departmentAffiche *do.DepartmentAffiche
	common.CopyFields(req, &departmentAffiche)
	departmentAffiche.DepartmentAfficheUpdateTime = time.Now().Unix()
	if _, err = dao.DepartmentAffiche.Ctx(ctx).Where(dao.DepartmentAffiche.Columns().DepartmentAfficheId, req.DepartmentAfficheId).Update(departmentAffiche); err != nil {
		g.Log().Errorf(ctx, "DepartmentAfficheModify error: %v", err)
		return nil, common.NewError(consts.InternalServerError, "modify department affiche error")
	}
	return
}
