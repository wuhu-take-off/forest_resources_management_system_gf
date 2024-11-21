// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PeopleApplication is the golang structure of table people_application for DAO operations like Where/Data.
type PeopleApplication struct {
	g.Meta                   `orm:"table:people_application, do:true"`
	PeopleApplicationId      interface{} // 民事申请ID
	PeopleApplicationTitle   interface{} // 民事申请标题
	PeopleApplicant          interface{} // 申请人
	PeopleApplicationPhone   interface{} // 联系电话
	PeopleApplicationType    interface{} // 申请类型
	PeopleApplicationContent interface{} // 申请内容
	PeopleApplicationAnnex   interface{} // 附件
	ModifyTime               interface{} // 修改时间
	PeopleApplicationState   interface{} // 申请状态(0:待审核 1:已处理)
	PeopleApplicationDesc    interface{} // 处理描述
}
