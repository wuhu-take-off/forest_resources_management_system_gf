package l_module

import (
	"context"
	"forest_resources_management_system_gf/api/module/module_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"forest_resources_management_system_gf/middleware"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultModule) ModuleList(ctx context.Context, req *module_v1.ModuleListReq) (res *module_v1.ModuleListRes, err error) {
	claims := ctx.Value(middleware.TokenClaims).(*common.Claims)
	if claims == nil {
		return nil, common.NewError(consts.Unauthorized, "token is invalid")
	}
	var moduleIdList []int
	if value, err := dao.User.Ctx(ctx).Where(dao.User.Columns().UserId, claims.UserId).Fields(dao.User.Columns().IdentityId).Value(); err != nil {
		g.Log().Errorf(ctx, "failed to get identity id of user: %v", err)
		return nil, common.NewError(consts.InternalServerError, "failed to get identity id of user")
	} else {
		if v, err := dao.Identity.Ctx(ctx).Where(dao.Identity.Columns().IdentityId, value.Int()).Value(dao.Identity.Columns().ModuleAuthority); err != nil {
			g.Log().Errorf(ctx, "failed to get authority of identity: %v", err)
			return nil, common.NewError(consts.InternalServerError, "failed to get authority of identity")
		} else {
			if err := v.Struct(&moduleIdList); err != nil {
				g.Log().Errorf(ctx, "failed to get authority of identity: %v", err)
				return nil, common.NewError(consts.InternalServerError, "failed to get authority of identity")
			}
		}
	}
	var modules []*entity.Module
	if err := dao.Module.Ctx(ctx).WhereIn(dao.Module.Columns().ModuleId, moduleIdList).Scan(&modules); err != nil {
		g.Log().Errorf(ctx, "failed to get module list: %v", err)
		return nil, common.NewError(consts.InternalServerError, "failed to get module list")
	}
	res = &module_v1.ModuleListRes{
		ModuleList: make([]*module_v1.ModuleListStruct, 0),
	}
	for i := range modules {
		if modules[i].ModuleParent == 0 {
			var moduleInfo *module_v1.ModuleListStruct
			common.CopyFields(modules[i], &moduleInfo)
			moduleInfo.ModuleChild = make([]*module_v1.ModuleListStruct, 0)
			res.ModuleList = append(res.ModuleList, moduleInfo)
		} else {
			for i2 := range res.ModuleList {
				if res.ModuleList[i2].ModuleId == modules[i].ModuleParent {
					var moduleInfo *module_v1.ModuleListStruct
					common.CopyFields(modules[i], &moduleInfo)
					res.ModuleList[i2].ModuleChild = append(res.ModuleList[i2].ModuleChild, moduleInfo)
					break
				}
			}
		}
	}
	return
}
