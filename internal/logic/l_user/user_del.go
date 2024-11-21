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

func (d defaultUserV2) UserDel(ctx context.Context, req *user_v2.UserDelReq) (res *user_v2.UserDelRes, err error) {
	if req.UserId == 1 {
		return nil, common.NewError(consts.Forbidden, "cannot delete admin user")
	}
	var user *do.User
	common.CopyFields(req, &user)
	if _, err = dao.User.Ctx(ctx).Delete(user); err != nil {
		g.Log().Errorf(ctx, "delete user failed, err:%v", err)
		return nil, common.NewError(consts.InternalServerError, "delete user failed")
	}
	return
}
