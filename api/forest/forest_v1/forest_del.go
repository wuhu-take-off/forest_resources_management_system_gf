package forest_v1

import "github.com/gogf/gf/v2/frame/g"

type ForestDelReq struct {
	g.Meta   `path:"/forest/del" method:"post" tags:"forest"`
	ForestId int `json:"forest_id" `
}
type ForestDelRes struct{}
