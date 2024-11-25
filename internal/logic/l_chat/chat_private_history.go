package l_chat

import (
	"context"
	"encoding/json"
	"forest_resources_management_system_gf/api/chat/char_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"forest_resources_management_system_gf/internal/model/entity"
	"forest_resources_management_system_gf/middleware"
	"github.com/gogf/gf/v2/frame/g"
)

func (d *defaultChat) ChatPrivateHistory(ctx context.Context, userId int) (res []*char_v1.ChatContent, err error) {
	claims := ctx.Value(middleware.TokenClaims).(*common.Claims)
	if claims == nil {
		return nil, common.NewError(consts.Unauthorized, "token is invalid")
	}
	var chatHistoryList []*entity.PrivateChat
	if err := dao.PrivateChat.Ctx(ctx).WhereOr(do.PrivateChat{PrivateChatSendId: claims.UserId, PrivateChatReceiverId: userId}).
		WhereOr(do.PrivateChat{PrivateChatSendId: userId, PrivateChatReceiverId: claims.UserId}).Scan(&chatHistoryList); err != nil {
		g.Log().Errorf(ctx, "failed to get private chat history: %v", err)
		return nil, common.NewError(consts.InternalServerError, "failed to get private chat history")
	}
	res = make([]*char_v1.ChatContent, 0)
	for _, chatHistory := range chatHistoryList {
		var chatContent *char_v1.ChatContent
		json.Unmarshal([]byte(chatHistory.PrivateChatMessage), &chatContent)
		res = append(res, chatContent)
	}
	return
}
