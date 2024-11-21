package user_v2

import "github.com/gogf/gf/v2/frame/g"

type UserDelReq struct {
	g.Meta `path:"/del" method:"post" summary:"删除用户" tags:"user"`
	UserId int `json:"user_id" description:"用户ID" `
}
type UserDelRes struct{}
