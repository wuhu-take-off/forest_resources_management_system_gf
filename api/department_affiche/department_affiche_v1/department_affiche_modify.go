package department_affiche_v1

import "github.com/gogf/gf/v2/frame/g"

type DepartmentAfficheModifyReq struct {
	g.Meta              `path:"/department_affiche/modify" method:"post" tags:"department_affiche"`
	DepartmentAfficheId int `json:"department_affiche_id"         ` // 部门公告主键

	DepartmentAfficheTitle   string `json:"title"`
	DepartmentAfficheContent string `json:"content" `
	DepartmentId             int32  `json:"department_id" description:"部门ID"`
}

type DepartmentAfficheModifyRes struct{}
