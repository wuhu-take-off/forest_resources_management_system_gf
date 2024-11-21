package forest_v1

import "github.com/gogf/gf/v2/frame/g"

type ForestCreateReq struct {
	g.Meta     `path:"/forest/create" method:"post" tags:"forest"`
	ForestName string `json:"forest_name"`
}
type ForestCreateRes struct {
}
