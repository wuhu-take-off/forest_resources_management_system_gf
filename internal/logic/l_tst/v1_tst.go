package l_tst

import (
	"context"
	"forest_resources_management_system_gf/api/tst/tst_v1"
	"forest_resources_management_system_gf/internal/dao"
)

func (d defaultTst) Tst(ctx context.Context, req *tst_v1.TstReq) (res *tst_v1.TstRes, err error) {

	dao.Identity.Ctx(ctx).Where(dao.Identity.Columns().IdentityId, map[int]struct{}{
		1: {},
	}).Scan()
}
