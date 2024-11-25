package department_affiche_v1

import "github.com/gogf/gf/v2/frame/g"

type DepartmentAfficheDelReq struct {
	g.Meta              `path:"/department_affiche/del" method:"post" tags:"department_affiche"`
	DepartmentAfficheId int `json:"department_affiche_id"`
}
type DepartmentAfficheDelRes struct{}
