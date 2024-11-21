package forest_resource_info

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_forest_resource_info"

	"forest_resources_management_system_gf/api/forest_resource_info/forest_resource_info_v1"
)

func (c *ControllerForest_resource_info_v1) ForestResourceInfoModify(ctx context.Context, req *forest_resource_info_v1.ForestResourceInfoModifyReq) (res *forest_resource_info_v1.ForestResourceInfoModifyRes, err error) {
	return l_forest_resource_info.NewForestResourceInfoLogic().ForestResourceInfoModify(ctx, req)

}