package l_forest_resource_info

import (
	"context"
	"encoding/json"
	"forest_resources_management_system_gf/api/forest_resource_info/forest_resource_info_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultForestResourceInfo) ForestResourceInfoList(ctx context.Context, req *forest_resource_info_v1.ForestResourceInfoListReq) (res *forest_resource_info_v1.ForestResourceInfoListRes, err error) {
	var forestResourceInfos []*entity.ForestResourceInfo
	//获取资源列表
	var condition do.ForestResourceInfo
	//条件查询
	if req.ForestId != nil && *req.ForestId != 0 {
		condition.ForestId = *req.ForestId
	}
	if req.TreeSpeciesId != nil && *req.TreeSpeciesId != 0 {
		condition.TreeSpeciesId = *req.TreeSpeciesId
	}
	var db = dao.ForestResourceInfo.Ctx(ctx)
	if req.Page == nil || *req.Page == 0 || req.Limit == nil || *req.Limit == 0 {
		db = db.Where(condition)
	} else {
		db = db.Where(condition).Page(*req.Page, *req.Limit)
	}
	if err := db.Scan(&forestResourceInfos); err != nil {
		g.Log().Errorf(ctx, "Error occurred while scanning forest resource infos: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Error occurred while scanning forest resource infos")
	}
	res = &forest_resource_info_v1.ForestResourceInfoListRes{
		ForestResourceInfoList: make([]*forest_resource_info_v1.ForestResourceInfoStruct, 0),
	}
	common.CopyFields(forestResourceInfos, &res.ForestResourceInfoList)

	//获取资源信息
	for i := range res.ForestResourceInfoList {
		json.Unmarshal([]byte(forestResourceInfos[i].ForestResourceExamineInfo), &res.ForestResourceInfoList[i].ForestResourceExamineInfo)
		json.Unmarshal([]byte(forestResourceInfos[i].ScenePhoto), &res.ForestResourceInfoList[i].ScenePhoto)
		json.Unmarshal([]byte(forestResourceInfos[i].SceneVideo), &res.ForestResourceInfoList[i].SceneVideo)
	}

	//获取数据总数
	if count, _ := dao.ForestResourceInfo.Ctx(ctx).Count(condition); count != 0 {
		res.TotalCount = count
	}
	return
}
