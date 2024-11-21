// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Department is the golang structure of table department for DAO operations like Where/Data.
type Department struct {
	g.Meta             `orm:"table:department, do:true"`
	DepartmentId       interface{} // 部门ID
	DepartmentName     interface{} // 部门名称
	DepartmentParentId interface{} // 父类部门
}
