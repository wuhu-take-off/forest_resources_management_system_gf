package forest

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_forest"

	"forest_resources_management_system_gf/api/forest/forest_v1"
)

func (c *ControllerForest_v1) ForestList(ctx context.Context, req *forest_v1.ForestListReq) (res *forest_v1.ForestListRes, err error) {
	return l_forest.NewForestLogic().ForestList(ctx, req)

}
