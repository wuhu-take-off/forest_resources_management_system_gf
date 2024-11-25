package user_v2

import "github.com/gogf/gf/v2/frame/g"

type SingleUserInfoReq struct {
	g.Meta `path:"/single/user/info" method:"post" tags:"user"`
}
type SingleUserInfoRes struct {
	UserId               int    `json:"user_id"`
	UserName             string `json:"user_name"`
	DepartmentId         int    `json:"department_id"`
	UserPhone            string `json:"user_phone"`
	IdentityId           int    `json:"identity_id"`
	AffiliatedDepartment []int  `json:"affiliated_department"` // 附属部门
}
