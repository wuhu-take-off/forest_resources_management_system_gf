package user

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_user"

	"forest_resources_management_system_gf/api/user/user_v2"
)

func (c *ControllerUser_v2) UserDel(ctx context.Context, req *user_v2.UserDelReq) (res *user_v2.UserDelRes, err error) {
	return l_user.NewUserLogicV2().UserDel(ctx, req)

}
