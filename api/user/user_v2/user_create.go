package user_v2

import "github.com/gogf/gf/v2/frame/g"

type UserCreateReq struct {
	g.Meta       `path:"/create" method:"post" tags:"user" summary:"创建用户"`
	UserName     string `json:"username"`
	UserPassword string `json:"password"`
	UserPhone    string `json:"phone"`
	DepartmentId int    `json:"department_id"` // 部门ID 不能为0
	IdentityId   int    `json:"identity_id"`   // 身份ID 不能为0
}
type UserCreateRes struct {
}
