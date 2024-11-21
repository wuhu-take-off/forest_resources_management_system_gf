// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package policy

import (
	"context"

	"forest_resources_management_system_gf/api/policy/policy_v1"
)

type IPolicyPolicy_v1 interface {
	PolicyCreate(ctx context.Context, req *policy_v1.PolicyCreateReq) (res *policy_v1.PolicyCreateRes, err error)
	PolicyDel(ctx context.Context, req *policy_v1.PolicyDelReq) (res *policy_v1.PolicyDelRes, err error)
	PolicyInfoList(ctx context.Context, req *policy_v1.PolicyInfoListReq) (res *policy_v1.PolicyInfoListRes, err error)
	PolicyList(ctx context.Context, req *policy_v1.PolicyListReq) (res *policy_v1.PolicyListRes, err error)
	PolicyModify(ctx context.Context, req *policy_v1.PolicyModifyReq) (res *policy_v1.PolicyModifyRes, err error)
}
