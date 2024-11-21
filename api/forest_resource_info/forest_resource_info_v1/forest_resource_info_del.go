package forest_resource_info_v1

import "github.com/gogf/gf/v2/frame/g"

type ForrestResourceInfoDelReq struct {
	g.Meta               `path:"/forest_resource_info/del" method:"post" tags:"forest_resource_info"`
	ForestResourceInfoId int `json:"forest_resource_info_id"`
}
type ForrestResourceInfoDelRes struct{}
