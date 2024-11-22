// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ForestResourceInfo is the golang structure for table forest_resource_info.
type ForestResourceInfo struct {
	ForestResourceInfoId      int    `json:"forestResourceInfoId"      orm:"forest_resource_info_id"      ` // 林地资源信息主键
	ForestId                  int    `json:"forestId"                  orm:"forest_id"                    ` // 林地id
	TreeSpeciesId             int    `json:"treeSpeciesId"             orm:"tree_species_id"              ` // 树种id
	ForestResourceExamineTime int64  `json:"forestResourceExamineTime" orm:"forest_resource_examine_time" ` // 检查日期
	ForestResourceExamineType int    `json:"forestResourceExamineType" orm:"forest_resource_examine_type" ` // 检查类型
	ForestStatus              int    `json:"forestStatus"              orm:"forest_status"                ` // 林地状态
	ForestResourceExamineInfo string `json:"forestResourceExamineInfo" orm:"forest_resource_examine_info" ` // 检测数据
	ProblemDes                string `json:"problemDes"                orm:"problem_des"                  ` // 问题描述
	TreatmentSuggestion       string `json:"treatmentSuggestion"       orm:"treatment_suggestion"         ` // 处理建议
	ScenePhoto                string `json:"scenePhoto"                orm:"scene_photo"                  ` // 现场照片
	SceneVideo                string `json:"sceneVideo"                orm:"scene_video"                  ` // 视频记录
}
