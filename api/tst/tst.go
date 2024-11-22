// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package tst

import (
	"context"

	"forest_resources_management_system_gf/api/tst/tst_v1"
)

type ITstTst_v1 interface {
	Tst(ctx context.Context, req *tst_v1.TstReq) (res *tst_v1.TstRes, err error)
}
