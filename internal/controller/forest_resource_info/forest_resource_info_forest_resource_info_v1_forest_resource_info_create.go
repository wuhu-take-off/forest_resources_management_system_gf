package forest_resource_info

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_forest_resource_info"

	"forest_resources_management_system_gf/api/forest_resource_info/forest_resource_info_v1"
)

func (c *ControllerForest_resource_info_v1) ForestResourceInfoCreate(ctx context.Context, req *forest_resource_info_v1.ForestResourceInfoCreateReq) (res *forest_resource_info_v1.ForestResourceInfoCreateRes, err error) {
	return l_forest_resource_info.NewForestResourceInfoLogic().ForestResourceInfoCreate(ctx, req)
}
