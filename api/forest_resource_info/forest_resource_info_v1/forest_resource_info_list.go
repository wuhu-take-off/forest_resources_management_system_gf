package forest_resource_info_v1

import (
	"forest_resources_management_system_gf/api/universal_class/universal_class_v1"
	"github.com/gogf/gf/v2/frame/g"
)

type ForestResourceInfoListReq struct {
	g.Meta        `path:"/forest_resource_info/list" method:"post" summary:"获取森林资源信息列表" tags:"forest_resource_info"`
	Page          *int `json:"page" `
	Limit         *int `json:"limit" `
	ForestId      *int `json:"forestId"                   ` // 林地id
	TreeSpeciesId *int `json:"treeSpeciesId"              ` // 树种id
}

type ForestResourceInfoListRes struct {
	ForestResourceInfoList []*ForestResourceInfoStruct `json:"forest_resource_info_list"` // 森林资源信息列表
	TotalCount             int                         `json:"total_count"`
}

type ForestResourceInfoStruct struct {
	ForestResourceInfoId      int                             `json:"forest_resource_info_id"`     // 森林资源信息id
	ForestId                  int                             `json:"forestId"                   ` // 林地id
	TreeSpeciesId             int                             `json:"treeSpeciesId"              ` // 树种id
	ForestResourceExamineTime string                          `json:"forestResourceExamineTime"  ` // 检查日期
	ForestResourceExamineType int                             `json:"forestResourceExamineType"  ` // 检查类型
	ForestStatus              int                             `json:"forestStatus"               ` // 林地状态
	ForestResourceExamineInfo ForestResourceExamineInfoStruct `json:"forestResourceExamineInfo"  ` // 检测数据
	ProblemDes                string                          `json:"problemDes"                 ` // 问题描述
	TreatmentSuggestion       string                          `json:"treatmentSuggestion"        ` // 处理建议
	ScenePhoto                []*universal_class_v1.Photo     `json:"scenePhoto"                 ` // 现场照片
	SceneVideo                []*universal_class_v1.Video     `json:"sceneVideo"                 ` // 视频记录
}
