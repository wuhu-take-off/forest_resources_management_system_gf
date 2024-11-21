package user_v1

import "github.com/gogf/gf/v2/frame/g"

type UserLoginReq struct {
	g.Meta       `path:"/login" method:"post" summary:"用户登录" tags:"user"`
	UserName     string `json:"username" `
	UserPassword string `json:"password" `
}
type UserLoginRes struct {
	UserId   int    `json:"userid"`    // 用户id
	UserName string `json:"user_name"` // 用户名
	Token    string `json:"token"`     // 用户登录成功后返回的token
}
