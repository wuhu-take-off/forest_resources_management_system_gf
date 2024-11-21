package resource_v1

import (
	"forest_resources_management_system_gf/api/universal_class/universal_class_v1"
	"github.com/gogf/gf/v2/frame/g"
)

type ResourceVideoReq struct {
	g.Meta `path:"/resource/video" method:"post" tags:"resource" summary:"获取视频资源"`
	Video  universal_class_v1.Video `json:"video"`
}
type ResourceVideoRes struct {
	FileNames string `json:"file_names"`
	Content   []byte `json:"content"`
}
type ResourcePhotoReq struct {
	g.Meta `path:"/resource/photo" method:"post" tags:"resource" summary:"获取图片资源"`
	Photo  universal_class_v1.Photo `json:"photo"`
}
type ResourcePhotoRes struct {
	FileNames string `json:"file_names"`
	Content   []byte `json:"content"`
}
