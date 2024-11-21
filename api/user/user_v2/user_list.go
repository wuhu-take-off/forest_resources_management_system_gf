package user_v2

import "github.com/gogf/gf/v2/frame/g"

type UserListReq struct {
	g.Meta       `path:"/list" method:"post" summary:"获取用户列表" tags:"user"`
	DepartmentId int `json:"department_id" description:"部门ID"`
}
type UserListStruct struct {
	UserId       int    `json:"user_id"`
	UserName     string `json:"username"`
	DepartmentId int    `json:"department_id"`
	UserPhone    string `json:"phone"`
	IdentityId   int    `json:"identity_id"`
}
type UserListRes struct {
	UserLists []*UserListStruct `json:"user_lists"`
}
