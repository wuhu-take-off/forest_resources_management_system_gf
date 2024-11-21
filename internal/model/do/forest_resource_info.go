// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ForestResourceInfo is the golang structure of table forest_resource_info for DAO operations like Where/Data.
type ForestResourceInfo struct {
	g.Meta                    `orm:"table:forest_resource_info, do:true"`
	ForestResourceInfoId      interface{} // 林地资源信息主键
	ForestId                  interface{} // 林地id
	TreeSpeciesId             interface{} // 树种id
	ForestResourceExamineTime interface{} // 检查日期
	ForestResourceExamineType interface{} // 检查类型
	ForestStatus              interface{} // 林地状态
	ForestResourceExamineInfo interface{} // 检测数据
	ProblemDes                interface{} // 问题描述
	TreatmentSuggestion       interface{} // 处理建议
	ScenePhoto                interface{} // 现场照片
	SceneVideo                interface{} // 视频记录
}
