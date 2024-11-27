// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Module is the golang structure for table module.
type Module struct {
	ModuleId     int    `json:"moduleId"     orm:"module_id"     ` // 模块Id
	ModuleName   string `json:"moduleName"   orm:"module_name"   ` // 模块名称
	ModuleDesc   string `json:"moduleDesc"   orm:"module_desc"   ` // 模块描述
	ModuleGrade  int    `json:"moduleGrade"  orm:"module_grade"  ` // 模块等级
	ModuleParent int    `json:"moduleParent" orm:"module_parent" ` // 父模块(0为顶级模块)
}
