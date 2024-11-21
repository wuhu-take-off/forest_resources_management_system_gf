package forest_v1

import "github.com/gogf/gf/v2/frame/g"

type ForestListReq struct {
	g.Meta `path:"/forest_list" method:"post" summary:"获取森林列表" tags:"forest"`
}
type ForestListStruct struct {
	ForestId   int    `json:"forest_id"`
	ForestName string `json:"forest_name"`
}
type ForestListRes struct {
	ForestList []*ForestListStruct `json:"forest_list"`
}
