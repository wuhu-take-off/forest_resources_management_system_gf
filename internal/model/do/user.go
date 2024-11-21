// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta       `orm:"table:user, do:true"`
	UserId       interface{} // 用户主键
	UserName     interface{} // 用户名
	UserPassword interface{} // 用户密码
	DepartmentId interface{} // 部门ID
	UserPhone    interface{} // 联系电话
	IdentityId   interface{} // 身份ID
}
