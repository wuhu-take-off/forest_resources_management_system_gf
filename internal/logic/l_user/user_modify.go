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

func (d defaultUserV2) UserModify(ctx context.Context, req *user_v2.UserModifyReq) (res *user_v2.UserModifyRes, err error) {
	var user *do.User
	common.CopyFields(req, &user)
	if _, err = dao.User.Ctx(ctx).Where(dao.User.Columns().UserId, req.UserId).Update(user); err != nil {
		g.Log().Errorf(ctx, "UserModify error: %v", err)
		return nil, common.NewError(consts.InternalServerError, "用户修改失败")
	}
	return
}
