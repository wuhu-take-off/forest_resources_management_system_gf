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

func (d *defaultChat) ChatGroupHistory(ctx context.Context, groupId int) (res []*char_v1.ChatContent, err error) {
	claims := ctx.Value(middleware.TokenClaims).(*common.Claims)
	if claims == nil {
		return nil, common.NewError(consts.Unauthorized, "token is invalid")
	}

	//todo:目前只支持群聊
	var chatGroupList []*entity.GroupChat
	if err := dao.GroupChat.Ctx(ctx).Where(do.GroupChat{GroupChatReceiverId: groupId}).
		Scan(&chatGroupList); err != nil {
		g.Log().Errorf(ctx, "failed to get group chat history: %v", err)
		return nil, common.NewError(consts.InternalServerError, "failed to get group chat history")
	}
	res = make([]*char_v1.ChatContent, 0)
	for _, chatHistory := range chatGroupList {
		var chatContent *char_v1.ChatContent
		json.Unmarshal([]byte(chatHistory.GroupChatMessage), &chatContent)
		res = append(res, chatContent)
	}
	return
}
