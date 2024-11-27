package l_department_affiche

import (
	"context"
	"forest_resources_management_system_gf/api/department_affiche/department_affiche_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"forest_resources_management_system_gf/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func (d defaultDepartmentAdapter) DepartmentAfficheCreate(ctx context.Context, req *department_affiche_v1.DepartmentAfficheCreateReq) (res *department_affiche_v1.DepartmentAfficheCreateRes, err error) {
	claims := ctx.Value(middleware.TokenClaims).(*common.Claims)
	if claims == nil {
		return nil, common.NewError(consts.Unauthorized, "token is invalid")
	}

	var departmentAfficheInfo *do.DepartmentAffiche
	common.CopyFields(req, &departmentAfficheInfo)

	departmentAfficheInfo.UserId = claims.UserId
	departmentAfficheInfo.DepartmentAfficheCreateTime = time.Now().Unix()
	departmentAfficheInfo.DepartmentAfficheUpdateTime = time.Now().Unix()
	if _, err := dao.DepartmentAffiche.Ctx(ctx).Insert(departmentAfficheInfo); err != nil {
		g.Log().Errorf(ctx, "create department affiche failed: %v", err)
		return nil, common.NewError(consts.InternalServerError, "create department affiche failed")
	}
	return
}
