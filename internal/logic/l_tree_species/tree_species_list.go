package l_tree_species

import (
	"context"
	"forest_resources_management_system_gf/api/tree_species/tree_species_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultTreeSpecies) TreeSpeciesList(ctx context.Context, req *tree_species_v1.TreeSpeciesListReq) (res *tree_species_v1.TreeSpeciesListRes, err error) {
	var entityList []*entity.TreeSpecies
	if err := dao.TreeSpecies.Ctx(ctx).Scan(&entityList); err != nil {
		g.Log().Errorf(ctx, "Error occurred while scanning tree species: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Error occurred while scanning tree species")
	}
	res = &tree_species_v1.TreeSpeciesListRes{
		TreeSpeciesList: make([]*tree_species_v1.TreeSpeciesStruct, 0),
	}
	common.CopyFields(entityList, &res.TreeSpeciesList)
	return
}
