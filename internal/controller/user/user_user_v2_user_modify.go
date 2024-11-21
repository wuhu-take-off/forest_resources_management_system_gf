package user

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_user"

	"forest_resources_management_system_gf/api/user/user_v2"
)

func (c *ControllerUser_v2) UserModify(ctx context.Context, req *user_v2.UserModifyReq) (res *user_v2.UserModifyRes, err error) {
	return l_user.NewUserLogicV2().UserModify(ctx, req)

}
