package l_forest_resource_info

import (
	"context"
	"forest_resources_management_system_gf/api/forest_resource_info/forest_resource_info_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultForestResourceInfo) ForrestResourceInfoDel(ctx context.Context, req *forest_resource_info_v1.ForrestResourceInfoDelReq) (res *forest_resource_info_v1.ForrestResourceInfoDelRes, err error) {
	var forestResourceInfo *do.ForestResourceInfo
	common.CopyFields(req, &forestResourceInfo)
	if _, err = dao.ForestResourceInfo.Ctx(ctx).Delete(forestResourceInfo); err != nil {
		g.Log().Errorf(ctx, "delete forest resource info failed, err:%v", err)
		return nil, common.NewError(consts.InternalServerError, "delete forest resource info failed")
	}
	return
}
