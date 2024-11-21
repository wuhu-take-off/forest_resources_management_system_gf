package l_user

import (
	"context"
	"forest_resources_management_system_gf/api/user/user_v2"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultUserV2) UserCreate(ctx context.Context, req *user_v2.UserCreateReq) (res *user_v2.UserCreateRes, err error) {
	var user *do.User
	common.CopyFields(req, &user)
	if _, err = dao.User.Ctx(ctx).Insert(user); err != nil {
		g.Log().Errorf(ctx, "Error inserting user: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Error inserting user")
	}

	return
}
