package user

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_user"

	"forest_resources_management_system_gf/api/user/user_v1"
)

func (c *ControllerUser_v1) UserLogin(ctx context.Context, req *user_v1.UserLoginReq) (res *user_v1.UserLoginRes, err error) {
	return l_user.NewUserLogicV1().UserLogin(ctx, req)
}
