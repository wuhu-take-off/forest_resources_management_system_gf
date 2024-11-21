package department_v1

import "github.com/gogf/gf/v2/frame/g"

type DepartmentListReq struct {
	g.Meta `path:"/list" method:"post" summary:"获取部门列表" tags:"department"`
}
type DepartmentListStruct struct {
	DepartmentId       int                     `json:"department_id"`
	DepartmentName     string                  `json:"department_name"`
	DepartmentParentId int                     `json:"department_parent_id"`
	Children           []*DepartmentListStruct `json:"children"`
}
type DepartmentListRes struct {
	DepartmentList []*DepartmentListStruct `json:"department_list"`
}
