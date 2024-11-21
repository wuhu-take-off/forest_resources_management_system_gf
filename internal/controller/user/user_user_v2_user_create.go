package user

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_user"

	"forest_resources_management_system_gf/api/user/user_v2"
)

func (c *ControllerUser_v2) UserCreate(ctx context.Context, req *user_v2.UserCreateReq) (res *user_v2.UserCreateRes, err error) {
	return l_user.NewUserLogicV2().UserCreate(ctx, req)

}
