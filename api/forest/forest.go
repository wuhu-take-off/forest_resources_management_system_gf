// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package forest

import (
	"context"

	"forest_resources_management_system_gf/api/forest/forest_v1"
)

type IForestForest_v1 interface {
	ForestCreate(ctx context.Context, req *forest_v1.ForestCreateReq) (res *forest_v1.ForestCreateRes, err error)
	ForestDel(ctx context.Context, req *forest_v1.ForestDelReq) (res *forest_v1.ForestDelRes, err error)
	ForestList(ctx context.Context, req *forest_v1.ForestListReq) (res *forest_v1.ForestListRes, err error)
}
