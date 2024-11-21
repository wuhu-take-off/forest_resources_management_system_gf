package chat

import (
	"context"
	"forest_resources_management_system_gf/api/chat/char_v1"
	"forest_resources_management_system_gf/internal/logic/l_chat"
)

func (c *ControllerChar_v1) CallChat(ctx context.Context, req *char_v1.CallChatReq) (res *char_v1.CallChatRes, err error) {
	return l_chat.NewChatLogic().CallChat(ctx, req)
}
