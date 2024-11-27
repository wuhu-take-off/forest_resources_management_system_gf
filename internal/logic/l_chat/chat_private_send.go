package l_chat

import (
	"context"
	"encoding/json"
	"forest_resources_management_system_gf/api/chat/char_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"forest_resources_management_system_gf/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
	"time"
)

func (d *defaultChat) ChatPrivateSend(ctx context.Context, req *char_v1.ChatPrivateSendReq) (res *char_v1.ChatPrivateSendRes, err error) {
	claims := ctx.Value(middleware.TokenClaims).(*common.Claims)
	if claims == nil {
		return nil, common.NewError(consts.Unauthorized, "token is invalid")
	}

	var message char_v1.ChatContent
	if req.File != nil {
		var fileName string
		if claims.UserId > req.PrivateChatReceiverId {
			fileName = strconv.Itoa(claims.UserId) + "-" + strconv.Itoa(req.PrivateChatReceiverId)
		} else {
			fileName = strconv.Itoa(req.PrivateChatReceiverId) + "-" + strconv.Itoa(claims.UserId)
		}
		if _, err := req.File.Save(consts.PrivateChatDir + fileName); err != nil {
			g.Log().Errorf(ctx, "Failed to save scene photo: %v", err)
			return nil, common.NewError(consts.InternalServerError, "Failed to save scene photo")
		}
		message = char_v1.ChatContent{
			Type:    consts.ChatTypeFile,
			Content: consts.PrivateChatDir + fileName + "/" + req.File.Filename,
		}

	} else if req.Message != nil {
		message = char_v1.ChatContent{
			Type:    consts.ChatTypeText,
			Content: *req.Message,
		}
	} else {
		message = char_v1.ChatContent{
			Type:    consts.ChatTypeEmpty,
			Content: "",
		}
	}

	message.SendTime = time.Now().Unix()
	message.SendId = claims.UserId
	message.ReceiverId = &req.PrivateChatReceiverId

	messageJson, _ := json.Marshal(message)
	privateChatInfo := entity.PrivateChat{
		PrivateChatSendId:     claims.UserId,
		PrivateChatReceiverId: req.PrivateChatReceiverId,
		PrivateChatCreateTime: message.SendTime,
		PrivateChatMessage:    string(messageJson),
	}
	if _, err = dao.PrivateChat.Ctx(ctx).Insert(privateChatInfo); err != nil {
		g.Log().Errorf(ctx, "Failed to insert private chat: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Failed to insert private chat")
	}
	//发送消息到私聊
	d.chatConn.SendMessage(message, claims.UserId, req.PrivateChatReceiverId)
	return
}
