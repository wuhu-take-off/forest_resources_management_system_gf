package char_v1

import "github.com/gogf/gf/v2/frame/g"

type ChatPrivateHistoryReq struct {
	g.Meta `path:"/chat/private/history" method:"post" tags:"chat"`
	UserId int `json:"user_id"`
}
type ChatPrivateHistoryRes struct {
	ChatPrivateHistoryList []*ChatContent `json:"list"`
}
