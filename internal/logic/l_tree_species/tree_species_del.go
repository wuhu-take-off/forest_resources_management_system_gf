package l_tree_species

import (
	"context"
	"forest_resources_management_system_gf/api/tree_species/tree_species_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultTreeSpecies) TreeSpeciesDel(ctx context.Context, req *tree_species_v1.TreeSpeciesDelReq) (res *tree_species_v1.TreeSpeciesDelRes, err error) {
	var treeSpecies *do.TreeSpecies
	common.CopyFields(req, &treeSpecies)
	if _, err = dao.TreeSpecies.Ctx(ctx).Delete(treeSpecies); err != nil {
		g.Log().Errorf(ctx, "delete tree species failed, err:%v", err)
		return nil, common.NewError(consts.InternalServerError, "delete tree species failed")
	}
	return
}
