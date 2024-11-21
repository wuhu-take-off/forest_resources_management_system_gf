package user_v2

import "github.com/gogf/gf/v2/frame/g"

type UserChatListReq struct {
	g.Meta `path:"/chat/list" method:"post" tags:"user"`
}
type UserChatListStruct struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
}
type UserChatListRes struct {
	UserChatList []*UserChatListStruct `json:"chat_list"`
}
