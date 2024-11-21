package department_v1

import "github.com/gogf/gf/v2/frame/g"

type DepartmentDelReq struct {
	g.Meta       `path:"/del" method:"post" summary:"删除部门" tags:"department"`
	DepartmentId int `json:"department_id" description:"部门ID"`
}
type DepartmentDelRes struct{}
