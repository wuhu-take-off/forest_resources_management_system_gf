package chat

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_chat"

	"forest_resources_management_system_gf/api/chat/char_v1"
)

func (c *ControllerChar_v1) ChatPrivateHistory(ctx context.Context, req *char_v1.ChatPrivateHistoryReq) (res *char_v1.ChatPrivateHistoryRes, err error) {
	return l_chat.NewChatLogic().ChatPrivateHistory(ctx, req)
}
