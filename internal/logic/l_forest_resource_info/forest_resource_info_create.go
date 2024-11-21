package l_forest_resource_info

import (
	"context"
	"encoding/json"
	"forest_resources_management_system_gf/api/forest_resource_info/forest_resource_info_v1"
	"forest_resources_management_system_gf/api/universal_class/universal_class_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
	"time"
)

func (d defaultForestResourceInfo) ForestResourceInfoCreate(ctx context.Context, req *forest_resource_info_v1.ForestResourceInfoCreateReq) (res *forest_resource_info_v1.ForestResourceInfoCreateRes, err error) {
	t := strconv.FormatInt(time.Now().Unix(), 10) + "/"
	photo := make([]*universal_class_v1.Photo, 0)
	video := make([]*universal_class_v1.Video, 0)

	// 保存图片
	for i := range req.ScenePhoto {
		if _, err := req.ScenePhoto[i].Save(consts.ForestResourceInfoPhotosDir + t); err != nil {
			g.Log().Errorf(ctx, "Failed to save scene photo: %v", err)
			return nil, common.NewError(consts.InternalServerError, "Failed to save scene photo")
		}
		photo = append(photo, &universal_class_v1.Photo{
			Url:  consts.ForestResourceInfoPhotosDir + t,
			Name: req.ScenePhoto[i].Filename,
		})
	}
	// 保存视频
	for i := range req.SceneVideo {
		if _, err := req.SceneVideo[i].Save(consts.ForestResourceInfoVideosDir + t); err != nil {
			g.Log().Errorf(ctx, "Failed to save scene video: %v", err)
			return nil, common.NewError(consts.InternalServerError, "Failed to save scene video")
		}
		video = append(video, &universal_class_v1.Video{
			Photo: universal_class_v1.Photo{
				Url:  consts.ForestResourceInfoVideosDir + t,
				Name: req.SceneVideo[i].Filename,
			},
		})
	}

	// 保存资源信息
	photoJson, _ := json.Marshal(photo)
	videoJson, _ := json.Marshal(video)
	forestResourceExamineInfoJson, _ := json.Marshal(req.ForestResourceExamineInfo)
	var forestResourceInfo *entity.ForestResourceInfo
	common.CopyFields(req, &forestResourceInfo)
	forestResourceInfo.ScenePhoto = string(photoJson)
	forestResourceInfo.SceneVideo = string(videoJson)
	forestResourceInfo.ForestResourceExamineInfo = string(forestResourceExamineInfoJson)

	if _, err = dao.ForestResourceInfo.Ctx(ctx).Insert(forestResourceInfo); err != nil {
		g.Log().Errorf(ctx, "Failed to create forest resource info: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Failed to create forest resource info")
	}
	return
}
