package l_forest

import (
	"context"
	"forest_resources_management_system_gf/api/forest/forest_v1"
)

type defaultForest struct{}

type IForestLogic interface {
	ForestCreate(ctx context.Context, req *forest_v1.ForestCreateReq) (res *forest_v1.ForestCreateRes, err error)
	ForestDel(ctx context.Context, req *forest_v1.ForestDelReq) (res *forest_v1.ForestDelRes, err error)
	ForestList(ctx context.Context, req *forest_v1.ForestListReq) (res *forest_v1.ForestListRes, err error)
}

func NewForestLogic() IForestLogic {
	return &defaultForest{}
}
