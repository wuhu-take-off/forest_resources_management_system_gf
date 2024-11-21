package tree_species_v1

import "github.com/gogf/gf/v2/frame/g"

type TreeSpeciesCreateReq struct {
	g.Meta          `path:"/tree_species/create" method:"post" tags:"tree_species"`
	TreeSpeciesName string `json:"tree_species_name"`
}
type TreeSpeciesCreateRes struct{}
