package char_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ChatPrivateSendReq struct {
	g.Meta                `path:"/chat/private/send" method:"post" tags:"chat"`
	PrivateChatReceiverId int               `json:"receiver_id"`
	File                  *ghttp.UploadFile `json:"file"`
	Message               *string           `json:"message"`
}

type ChatPrivateSendRes struct {
}
