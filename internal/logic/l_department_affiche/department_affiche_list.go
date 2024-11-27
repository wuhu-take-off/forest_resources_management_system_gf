package l_department_affiche

import (
	"context"
	"forest_resources_management_system_gf/api/department_affiche/department_affiche_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"forest_resources_management_system_gf/middleware"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultDepartmentAdapter) DepartmentAfficheList(ctx context.Context, req *department_affiche_v1.DepartmentAfficheListReq) (res *department_affiche_v1.DepartmentAfficheListRes, err error) {
	claims := ctx.Value(middleware.TokenClaims).(*common.Claims)
	if claims == nil {
		return nil, common.NewError(consts.Unauthorized, "token is invalid")
	}

	var departmentAfficheList []*entity.DepartmentAffiche

	////获取用户所在的部门
	//department := l_department.NewDepartmentTool().Ctx(ctx).GetDepartment(claims.UserId)
	//var departmentIds []int
	//for i := range department {
	//	departmentIds = append(departmentIds, department[i].DepartmentId)
	//}
	if err := dao.DepartmentAffiche.Ctx(ctx).WhereIn(dao.DepartmentAffiche.Columns().DepartmentId, req.DepartmentId).Scan(&departmentAfficheList); err != nil {
		g.Log().Errorf(ctx, "failed to scan department affiche list, err: %v", err)
		return nil, common.NewError(consts.InternalServerError, "failed to scan department affiche list")
	}

	res = &department_affiche_v1.DepartmentAfficheListRes{
		DepartmentAfficheList: make([]*department_affiche_v1.DepartmentAfficheListStruct, 0),
	}

	common.CopyFields(departmentAfficheList, &res.DepartmentAfficheList)
	return
}
