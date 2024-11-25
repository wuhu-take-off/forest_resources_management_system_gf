package char_v1

import "github.com/gogf/gf/v2/frame/g"

type ChatHistoryReq struct {
	g.Meta  `path:"/chat/history" method:"post" tags:"chat"`
	UserId  *int `json:"user_id"`
	GroupId *int `json:"group_id"`
}
type ChatHistoryRes struct {
	ChatGroupIdHistoryList []*ChatContent `json:"chat_history_list"`
}
