package char_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ChatGroupSendReq struct {
	g.Meta              `path:"/chat/group/send" method:"post" tags:"chat"`
	GroupChatReceiverId int               `json:"receiver_id"`
	File                *ghttp.UploadFile `json:"file"`
	Message             *string           `json:"message"`
}

type ChatGroupSendRes struct {
}
