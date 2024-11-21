package forest_resource_info_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ForestResourceInfoCreateReq struct {
	g.Meta                    `path:"/forest_resource_info/create" method:"post" tags:"forest_resource_info"`
	ForestId                  int                             `json:"forestId"                   ` // 林地id
	TreeSpeciesId             int                             `json:"treeSpeciesId"              ` // 树种id
	ForestResourceExamineTime string                          `json:"forestResourceExamineTime"  ` // 检查日期
	ForestResourceExamineType int                             `json:"forestResourceExamineType"  ` // 检查类型
	ForestStatus              int                             `json:"forestStatus"               ` // 林地状态
	ForestResourceExamineInfo ForestResourceExamineInfoStruct `json:"forestResourceExamineInfo"  ` // 检测数据
	ProblemDes                string                          `json:"problemDes"                 ` // 问题描述
	TreatmentSuggestion       string                          `json:"treatmentSuggestion"        ` // 处理建议
	ScenePhoto                []*ghttp.UploadFile             `json:"scenePhoto"                 ` // 现场照片
	SceneVideo                []*ghttp.UploadFile             `json:"sceneVideo"                 ` // 视频记录
}
type ForestResourceExamineInfoStruct struct {
	MeanHeight float64 `json:"meanHeight"` // 平均树高
	MeanDBH    float64 `json:"meanDBH"`    // 平均树径
	Density    float64 `json:"density"`    // 密度 (树/亩)
}
type ForestResourceInfoCreateRes struct{}
