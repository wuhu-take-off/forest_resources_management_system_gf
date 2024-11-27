// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DepartmentAffiche is the golang structure of table department_affiche for DAO operations like Where/Data.
type DepartmentAffiche struct {
	g.Meta                      `orm:"table:department_affiche, do:true"`
	DepartmentAfficheId         interface{} // 部门公告主键
	DepartmentAfficheTitle      interface{} // 部门公告标题
	DepartmentAfficheContent    interface{} // 部门公告内容
	DepartmentId                interface{} // 发布部门
	UserId                      interface{} // 创建用户
	DepartmentAfficheCreateTime interface{} // 创建时间
	DepartmentAfficheUpdateTime interface{} // 修改时间
}
