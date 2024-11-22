package l_chat

import (
	"context"
	"encoding/json"
	"forest_resources_management_system_gf/api/chat/char_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/logic/l_department"
	"forest_resources_management_system_gf/internal/model/entity"
	"forest_resources_management_system_gf/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
	"time"
)

func (d *defaultChat) ChatGroupSend(ctx context.Context, req *char_v1.ChatGroupSendReq) (res *char_v1.ChatGroupSendRes, err error) {
	claims := ctx.Value(middleware.TokenClaims).(*common.Claims)
	if claims == nil {
		return nil, common.NewError(consts.Unauthorized, "token is invalid")
	}

	var message char_v1.ChatContent
	if req.File != nil {
		fileName := strconv.Itoa(req.GroupChatReceiverId)
		if _, err := req.File.Save(consts.GroupChatDir + fileName); err != nil {
			g.Log().Errorf(ctx, "Failed to save scene photo: %v", err)
			return nil, common.NewError(consts.InternalServerError, "Failed to save scene photo")
		}
		message = char_v1.ChatContent{
			Type:     consts.ChatTypeFile,
			SendId:   claims.UserId,
			SendTime: time.Now().Unix(),
			Content:  consts.GroupChatDir + fileName + "/" + req.File.Filename,
		}

	} else if req.Message != nil {
		message = char_v1.ChatContent{
			Type:     consts.ChatTypeText,
			SendId:   claims.UserId,
			SendTime: time.Now().Unix(),
			Content:  *req.Message,
		}
	} else {
		message = char_v1.ChatContent{
			Type:     consts.ChatTypeEmpty,
			SendId:   claims.UserId,
			SendTime: time.Now().Unix(),
			Content:  "",
		}
	}
	messageJson, _ := json.Marshal(message)
	groupChat := entity.GroupChat{
		GroupChatSendId:     claims.UserId,
		GroupChatReceiverId: req.GroupChatReceiverId,
		GroupChatCreateTime: message.SendTime,
		GroupChatMessage:    string(messageJson),
	}
	if _, err = dao.GroupChat.Ctx(ctx).Insert(groupChat); err != nil {
		g.Log().Errorf(ctx, "Failed to insert private chat: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Failed to insert private chat")
	}
	//获取部门下所有成员
	employees := l_department.NewDepartmentTool().Ctx(ctx).GetDepartmentEmployees(req.GroupChatReceiverId)
	var userIds []int
	for i := range employees {
		userIds = append(userIds, employees[i].UserId)
	}
	//发送消息到私聊
	d.chatConn.SendMessage(message, userIds...)
	return
}
