// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Policy is the golang structure of table policy for DAO operations like Where/Data.
type Policy struct {
	g.Meta           `orm:"table:policy, do:true"`
	PolicyId         interface{} // 政策Id
	PolicyHeadline   interface{} // 政策标题
	PolicyDepartment interface{} // 发布部门
	PolicyTime       interface{} // 发布日期
	PolicyType       interface{} // 政策类型
	PolicyContent    interface{} // 政策内容
	PolicyAnnex      interface{} // 附件
}
