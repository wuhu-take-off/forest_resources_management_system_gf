package tree_species

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_tree_species"

	"forest_resources_management_system_gf/api/tree_species/tree_species_v1"
)

func (c *ControllerTree_species_v1) TreeSpeciesDel(ctx context.Context, req *tree_species_v1.TreeSpeciesDelReq) (res *tree_species_v1.TreeSpeciesDelRes, err error) {
	return l_tree_species.NewTreeSpeciesLogic().TreeSpeciesDel(ctx, req)

}
