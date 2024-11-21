// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package chat

import (
	"context"

	"forest_resources_management_system_gf/api/chat/char_v1"
)

type IChatChar_v1 interface {
	CallChat(ctx context.Context, req *char_v1.CallChatReq) (res *char_v1.CallChatRes, err error)
	ChatPrivateHistory(ctx context.Context, req *char_v1.ChatPrivateHistoryReq) (res *char_v1.ChatPrivateHistoryRes, err error)
	ChatPrivateSend(ctx context.Context, req *char_v1.ChatPrivateSendReq) (res *char_v1.ChatPrivateSendRes, err error)
}
