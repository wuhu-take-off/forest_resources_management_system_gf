package user

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_user"

	"forest_resources_management_system_gf/api/user/user_v2"
)

func (c *ControllerUser_v2) UserChatList(ctx context.Context, req *user_v2.UserChatListReq) (res *user_v2.UserChatListRes, err error) {
	return l_user.NewUserLogicV2().UserChatList(ctx, req)
}
