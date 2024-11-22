package l_tst

import (
	"context"
	"forest_resources_management_system_gf/api/tst/tst_v1"
)

func (d defaultTst) Tst(ctx context.Context, req *tst_v1.TstReq) (res *tst_v1.TstRes, err error) {
	return &tst_v1.TstRes{
		SubContent1: &tst_v1.SubContent1{
			Content1: "Hello World1",
		},
		SubContent2: &tst_v1.SubContent2{
			Content2: "Hello World2",
		},
	}, nil
}
