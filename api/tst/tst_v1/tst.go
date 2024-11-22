package tst_v1

import "github.com/gogf/gf/v2/frame/g"

type TstReq struct {
	g.Meta `path:"/l_tst" method:"post" summary:"测试接口"`
}
type TstRes struct{}
