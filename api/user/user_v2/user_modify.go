package user_v2

import "github.com/gogf/gf/v2/frame/g"

type UserModifyReq struct {
	g.Meta       `path:"/modify" method:"post" tags:"user" summary:"修改用户信息" description:"修改用户信息"`
	UserId       int    `json:"user_id"`
	Username     string `json:"username"`
	UserPhone    string `json:"phone"`
	UserPassword string `json:"password"`
	IdentityId   int    `json:"identity_id"`
}
type UserModifyRes struct{}
