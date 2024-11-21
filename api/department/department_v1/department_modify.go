package department_v1

import "github.com/gogf/gf/v2/frame/g"

type DepartmentModifyReq struct {
	g.Meta             `path:"/modify" method:"post" tags:"department" summary:"修改部门信息" description:"修改部门信息"`
	DepartmentId       int64  `json:"department_id" description:"部门ID"`
	DepartmentName     string `json:"department_name" description:"部门名称"`
	DepartmentParentId int64  `json:"department_parent_id" description:"父部门ID"`
}
type DepartmentModifyRes struct{}
