package forest_resource_info_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ForestResourceInfoModifyReq struct {
	g.Meta                    `path:"/forest_resource_info/modify" method:"post" tags:"forest_resource_info" summary:"修改森林资源信息" description:"修改森林资源信息"`
	ForestResourceInfoId      int                             `json:"forest_resource_info_id"`     // 森林资源信息id
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
type ForestResourceInfoModifyRes struct{}
