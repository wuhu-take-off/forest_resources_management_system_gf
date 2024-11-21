package tree_species_v1

import "github.com/gogf/gf/v2/frame/g"

type TreeSpeciesListReq struct {
	g.Meta `path:"/tree_species_list" method:"get" tags:"tree_species"`
}
type TreeSpeciesListRes struct {
	TreeSpeciesList []*TreeSpeciesStruct `json:"tree_species_list"`
}
type TreeSpeciesStruct struct {
	TreeSpeciesId   int    `json:"tree_species_id"`
	TreeSpeciesName string `json:"tree_species_name"`
}
