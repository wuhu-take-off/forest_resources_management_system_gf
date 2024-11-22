package tst_v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TstReq struct {
	g.Meta `path:"/tst" method:"post" summary:"测试接口"`
}
type TstRes struct {
	*SubContent1
	*SubContent2
}
type SupContent struct {
}

type SubContent1 struct {
	Content1 string `json:"content1"`
}
type SubContent2 struct {
	Content2 string `json:"content2"`
}
