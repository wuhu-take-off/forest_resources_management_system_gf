package department_affiche_v1

import "github.com/gogf/gf/v2/frame/g"

type DepartmentAfficheModifyReq struct {
	g.Meta                   `path:"/department_affiche/modify" method:"post" tags:"department_affiche"`
	DepartmentAfficheTitle   string `json:"title"`
	DepartmentAfficheContent string `json:"content" `
	DepartmentId             int32  `json:"department_id" description:"部门ID"`
}
