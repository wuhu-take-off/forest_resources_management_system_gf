package l_chat

import (
	"context"
	"forest_resources_management_system_gf/api/chat/char_v1"
)

func (d *defaultChat) ChatHistory(ctx context.Context, req *char_v1.ChatHistoryReq) (res *char_v1.ChatHistoryRes, err error) {
	res = &char_v1.ChatHistoryRes{
		ChatGroupIdHistoryList: make([]*char_v1.ChatContent, 0),
	}
	if req.UserId != nil {
		if res.ChatGroupIdHistoryList, err = d.ChatPrivateHistory(ctx, *req.UserId); err != nil {
			return nil, err
		}
	}
	if req.GroupId != nil {
		if res.ChatGroupIdHistoryList, err = d.ChatGroupHistory(ctx, *req.GroupId); err != nil {
			return nil, err
		}
	}
	return
}
