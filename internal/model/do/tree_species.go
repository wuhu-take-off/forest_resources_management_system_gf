// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TreeSpecies is the golang structure of table tree_species for DAO operations like Where/Data.
type TreeSpecies struct {
	g.Meta          `orm:"table:tree_species, do:true"`
	TreeSpeciesId   interface{} // 树种ID
	TreeSpeciesName interface{} // 树种名称
}
