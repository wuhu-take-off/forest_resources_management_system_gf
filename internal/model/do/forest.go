// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Forest is the golang structure of table forest for DAO operations like Where/Data.
type Forest struct {
	g.Meta     `orm:"table:forest, do:true"`
	ForestId   interface{} // 林地主键
	ForestName interface{} // 林地名
}
