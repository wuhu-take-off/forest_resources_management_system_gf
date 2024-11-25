package l_user

import (
	"context"
	"forest_resources_management_system_gf/api/user/user_v2"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/logic/l_department"
	"forest_resources_management_system_gf/internal/model/entity"
	"forest_resources_management_system_gf/middleware"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultUserV2) SingleUserInfo(ctx context.Context, req *user_v2.SingleUserInfoReq) (res *user_v2.SingleUserInfoRes, err error) {
	claims := ctx.Value(middleware.TokenClaims).(*common.Claims)
	if claims == nil {
		return nil, common.NewError(consts.Unauthorized, "token is invalid")
	}
	if one, err := dao.User.Ctx(ctx).Where(dao.User.Columns().UserId, claims.UserId).One(); err != nil {
		g.Log().Errorf(ctx, "failed to get user info: %v", err)
		return nil, common.NewError(consts.InternalServerError, "failed to get user info")
	} else {
		var userInfoTmp *entity.User
		if err := one.Struct(&userInfoTmp); err != nil {
			g.Log().Errorf(ctx, "failed to get user info: %v", err)
			return nil, common.NewError(consts.InternalServerError, "failed to get user info")
		}
		common.CopyFields(userInfoTmp, &res)
	}
	department := l_department.NewDepartmentTool().Ctx(ctx).GetDepartment(claims.UserId)
	for i := range department {
		res.AffiliatedDepartment = append(res.AffiliatedDepartment, department[i].DepartmentId)
	}
	return
}
