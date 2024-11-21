package forest

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_forest"

	"forest_resources_management_system_gf/api/forest/forest_v1"
)

func (c *ControllerForest_v1) ForestDel(ctx context.Context, req *forest_v1.ForestDelReq) (res *forest_v1.ForestDelRes, err error) {
	return l_forest.NewForestLogic().ForestDel(ctx, req)

}
