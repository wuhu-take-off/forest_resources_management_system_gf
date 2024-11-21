package forest

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_forest"

	"forest_resources_management_system_gf/api/forest/forest_v1"
)

func (c *ControllerForest_v1) ForestCreate(ctx context.Context, req *forest_v1.ForestCreateReq) (res *forest_v1.ForestCreateRes, err error) {
	return l_forest.NewForestLogic().ForestCreate(ctx, req)

}
