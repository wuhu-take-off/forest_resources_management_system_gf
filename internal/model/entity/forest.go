// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Forest is the golang structure for table forest.
type Forest struct {
	ForestId   int    `json:"forestId"   orm:"forest_id"   ` // 林地主键
	ForestName string `json:"forestName" orm:"forest_name" ` // 林地名
}
