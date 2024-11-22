package chat

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_chat"

	"forest_resources_management_system_gf/api/chat/char_v1"
)

func (c *ControllerChar_v1) ChatGroupSend(ctx context.Context, req *char_v1.ChatGroupSendReq) (res *char_v1.ChatGroupSendRes, err error) {
	return l_chat.NewChatLogic().ChatGroupSend(ctx, req)

}
