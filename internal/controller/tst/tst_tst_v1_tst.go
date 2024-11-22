package tst

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_tst"

	"forest_resources_management_system_gf/api/tst/tst_v1"
)

func (c *ControllerTst_v1) Tst(ctx context.Context, req *tst_v1.TstReq) (res *tst_v1.TstRes, err error) {
	return l_tst.NewTstLogic().Tst(ctx, req)
}
