package department_v1

import "github.com/gogf/gf/v2/frame/g"

type DepartmentCreateReq struct {
	g.Meta             `path:"/create" method:"post" tags:"department" summary:"创建部门" description:"创建部门"`
	DepartmentName     string `json:"department_name" description:"部门名称"`
	DepartmentParentId int64  `json:"department_parent_id" description:"父部门ID"`
}
type DepartmentCreateRes struct{}
