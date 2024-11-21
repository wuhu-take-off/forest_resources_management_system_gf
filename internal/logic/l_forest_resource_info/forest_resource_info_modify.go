package l_forest_resource_info

import (
	"context"
	"encoding/json"
	"forest_resources_management_system_gf/api/forest_resource_info/forest_resource_info_v1"
	"forest_resources_management_system_gf/api/universal_class/universal_class_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
	"time"
)

func (d defaultForestResourceInfo) ForestResourceInfoModify(ctx context.Context, req *forest_resource_info_v1.ForestResourceInfoModifyReq) (res *forest_resource_info_v1.ForestResourceInfoModifyRes, err error) {
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

	var forestResourceInfo *do.ForestResourceInfo
	// 保存资源信息
	if len(photo) > 0 {
		photoJson, _ := json.Marshal(photo)
		forestResourceInfo.ScenePhoto = string(photoJson)
	}
	if len(video) > 0 {
		videoJson, _ := json.Marshal(video)
		forestResourceInfo.SceneVideo = string(videoJson)
	}

	forestResourceExamineInfoJson, _ := json.Marshal(req.ForestResourceExamineInfo)
	common.CopyFields(req, &forestResourceInfo)
	forestResourceInfo.ForestResourceExamineInfo = string(forestResourceExamineInfoJson)
	forestResourceInfo.ForestResourceInfoId = nil
	if _, err = dao.ForestResourceInfo.Ctx(ctx).Where(dao.ForestResourceInfo.Columns().ForestResourceInfoId, req.ForestResourceInfoId).Update(forestResourceInfo); err != nil {
		g.Log().Errorf(ctx, "DepartmentModify error: %v", err)
		return nil, common.NewError(consts.InternalServerError, "DepartmentModify error")
	}
	return
}
