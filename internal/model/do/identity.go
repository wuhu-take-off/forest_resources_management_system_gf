// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Identity is the golang structure of table identity for DAO operations like Where/Data.
type Identity struct {
	g.Meta       `orm:"table:identity, do:true"`
	IdentityId   interface{} // 身份ID
	IdentityName interface{} // 身份名称
}
