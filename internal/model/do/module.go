// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Module is the golang structure of table module for DAO operations like Where/Data.
type Module struct {
	g.Meta       `orm:"table:module, do:true"`
	ModuleId     interface{} // 模块Id
	ModuleName   interface{} // 模块名称
	ModuleDesc   interface{} // 模块描述
	ModuleGrade  interface{} // 模块等级
	ModuleParent interface{} // 父模块(0为顶级模块)
}
