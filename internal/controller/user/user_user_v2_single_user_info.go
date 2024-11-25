package user

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_user"

	"forest_resources_management_system_gf/api/user/user_v2"
)

func (c *ControllerUser_v2) SingleUserInfo(ctx context.Context, req *user_v2.SingleUserInfoReq) (res *user_v2.SingleUserInfoRes, err error) {
	return l_user.NewUserLogicV2().SingleUserInfo(ctx, req)
}
