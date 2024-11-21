package l_chat

import (
	"context"
	"forest_resources_management_system_gf/api/chat/char_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
)

func (d *defaultChat) CallChat(ctx context.Context, req *char_v1.CallChatReq) (res *char_v1.CallChatRes, err error) {

	request := g.RequestFromCtx(ctx)
	ws, err := common.UptoWS(request)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to get up to ws: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Failed to get up to ws")
	}
	var value *common.Claims
	if _, p, err := ws.ReadMessage(); err != nil {
		g.Log().Errorf(ctx, "Failed to read message: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Failed to read message")
	} else {
		if value, err = common.ParseToken(string(p)); err != nil {
			g.Log().Errorf(ctx, "Failed to parse token: %v", err)
			return nil, common.NewError(consts.InternalServerError, "Failed to parse token")
		}
	}

	d.chatConn.AddConn(value.UserId, ws)

	////监听聊天消息
	//d.chatConn.SendMessage(char_v1.ChatContent{
	//	Type:    1,
	//	SendId:  value.UserId,
	//	Content: "你好，欢迎来到Forest Resources Management System",
	//}, value.UserId)
	//d.chatConn.SendMessage(char_v1.ChatContent{
	//	Type:    1,
	//	SendId:  2,
	//	Content: "你可以通过输入指令来与系统进行交互，例如：/help 查看指令列表",
	//}, value.UserId)
	return
}
