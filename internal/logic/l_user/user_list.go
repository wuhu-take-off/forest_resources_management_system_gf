package l_user

import (
	"context"
	"forest_resources_management_system_gf/api/user/user_v2"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultUserV2) UserList(ctx context.Context, req *user_v2.UserListReq) (res *user_v2.UserListRes, err error) {
	var user *do.User
	common.CopyFields(req, &user)
	var users []*entity.User
	if err := dao.User.Ctx(ctx).Where(user).Scan(&users); err != nil {
		g.Log().Errorf(ctx, "UserList failed: %v", err)
		return nil, common.NewError(consts.InternalServerError, "UserList failed")
	}
	res = &user_v2.UserListRes{
		UserLists: make([]*user_v2.UserListStruct, 0),
	}
	common.CopyFields(users, &res.UserLists)
	return
}
