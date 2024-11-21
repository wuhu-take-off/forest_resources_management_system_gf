package l_identity

import (
	"context"
	"forest_resources_management_system_gf/api/identity/identity_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultIdentity) IdentityList(ctx context.Context, req *identity_v1.IdentityListReq) (res *identity_v1.IdentityListRes, err error) {
	var identities []*entity.Identity
	if err := dao.Identity.Ctx(ctx).Scan(&identities); err != nil {
		g.Log().Errorf(ctx, "Error occurred while scanning identities: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Error occurred while scanning identities")
	}
	res = &identity_v1.IdentityListRes{
		IdentityList: make([]*identity_v1.IdentityListStruct, 0),
	}
	common.CopyFields(identities, &res.IdentityList)
	return
}
