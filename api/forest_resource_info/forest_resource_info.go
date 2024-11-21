// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package forest_resource_info

import (
	"context"

	"forest_resources_management_system_gf/api/forest_resource_info/forest_resource_info_v1"
)

type IForestResourceInfoForest_resource_info_v1 interface {
	ForestResourceInfoCreate(ctx context.Context, req *forest_resource_info_v1.ForestResourceInfoCreateReq) (res *forest_resource_info_v1.ForestResourceInfoCreateRes, err error)
	ForrestResourceInfoDel(ctx context.Context, req *forest_resource_info_v1.ForrestResourceInfoDelReq) (res *forest_resource_info_v1.ForrestResourceInfoDelRes, err error)
	ForestResourceInfoList(ctx context.Context, req *forest_resource_info_v1.ForestResourceInfoListReq) (res *forest_resource_info_v1.ForestResourceInfoListRes, err error)
	ForestResourceInfoModify(ctx context.Context, req *forest_resource_info_v1.ForestResourceInfoModifyReq) (res *forest_resource_info_v1.ForestResourceInfoModifyRes, err error)
}
