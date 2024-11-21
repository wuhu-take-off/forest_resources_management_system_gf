package tree_species_v1

import "github.com/gogf/gf/v2/frame/g"

type TreeSpeciesDelReq struct {
	g.Meta        `path:"/tree_species/del" method:"post" tags:"tree_species"`
	TreeSpeciesId int `json:"tree_species_id"`
}
type TreeSpeciesDelRes struct {
}
