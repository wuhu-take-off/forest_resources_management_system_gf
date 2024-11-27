package module_v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ModuleListReq struct {
	g.Meta `path:"/module/list" method:"post" tags:"module" `
}
type ModuleListStruct struct {
	ModuleId    int                 `json:"module_id"`    // 模块Id
	ModuleName  string              `json:"module_name"`  // 模块名称
	ModuleDesc  string              `json:"module_desc"`  // 模块描述
	ModuleGrade int                 `json:"module_grade"` // 模块等级
	ModuleChild []*ModuleListStruct `json:"module_child"` // 子模块
}
type ModuleListRes struct {
	ModuleList []*ModuleListStruct `json:"module_list"`
}
