package tree_species

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_tree_species"

	"forest_resources_management_system_gf/api/tree_species/tree_species_v1"
)

func (c *ControllerTree_species_v1) TreeSpeciesList(ctx context.Context, req *tree_species_v1.TreeSpeciesListReq) (res *tree_species_v1.TreeSpeciesListRes, err error) {
	return l_tree_species.NewTreeSpeciesLogic().TreeSpeciesList(ctx, req)

}
