package l_forest

import (
	"context"
	"forest_resources_management_system_gf/api/forest/forest_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultForest) ForestCreate(ctx context.Context, req *forest_v1.ForestCreateReq) (res *forest_v1.ForestCreateRes, err error) {
	var forest *do.Forest
	common.CopyFields(req, &forest)
	if _, err = dao.Forest.Ctx(ctx).Insert(forest); err != nil {
		g.Log().Errorf(ctx, "create forest failed, err: %v", err)
		return nil, common.NewError(consts.InternalServerError, "create forest failed")
	}
	return
}
