package l_forest

import (
	"context"
	"forest_resources_management_system_gf/api/forest/forest_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultForest) ForestList(ctx context.Context, req *forest_v1.ForestListReq) (res *forest_v1.ForestListRes, err error) {
	var forests []*entity.Forest
	if err := dao.Forest.Ctx(ctx).Scan(&forests); err != nil {
		g.Log().Errorf(ctx, "Error occurred while scanning forests: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Error occurred while scanning forests")
	}
	res = &forest_v1.ForestListRes{
		ForestList: make([]*forest_v1.ForestListStruct, 0),
	}
	common.CopyFields(forests, &res.ForestList)
	return
}
