package l_tree_species

import (
	"context"
	"forest_resources_management_system_gf/api/tree_species/tree_species_v1"
)

type defaultTreeSpecies struct{}

type ITreeSpeciesLogic interface {
	TreeSpeciesCreate(ctx context.Context, req *tree_species_v1.TreeSpeciesCreateReq) (res *tree_species_v1.TreeSpeciesCreateRes, err error)
	TreeSpeciesDel(ctx context.Context, req *tree_species_v1.TreeSpeciesDelReq) (res *tree_species_v1.TreeSpeciesDelRes, err error)
	TreeSpeciesList(ctx context.Context, req *tree_species_v1.TreeSpeciesListReq) (res *tree_species_v1.TreeSpeciesListRes, err error)
}

func NewTreeSpeciesLogic() ITreeSpeciesLogic {
	return &defaultTreeSpecies{}
}