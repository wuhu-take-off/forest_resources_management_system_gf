package user

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_user"

	"forest_resources_management_system_gf/api/user/user_v2"
)

func (c *ControllerUser_v2) UserList(ctx context.Context, req *user_v2.UserListReq) (res *user_v2.UserListRes, err error) {
	return l_user.NewUserLogicV2().UserList(ctx, req)

}
