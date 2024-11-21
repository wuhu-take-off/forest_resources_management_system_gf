package policy_v1

import "github.com/gogf/gf/v2/frame/g"

type PolicyDelReq struct {
	g.Meta   `path:"/policy/del" method:"post" tags:"policy" `
	PolicyId int `json:"policy_id" `
}
type PolicyDelRes struct {
}
