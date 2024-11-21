package policy

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_policy"

	"forest_resources_management_system_gf/api/policy/policy_v1"
)

func (c *ControllerPolicy_v1) PolicyDel(ctx context.Context, req *policy_v1.PolicyDelReq) (res *policy_v1.PolicyDelRes, err error) {
	return l_policy.NewPolicyLogic().PolicyDel(ctx, req)

}
