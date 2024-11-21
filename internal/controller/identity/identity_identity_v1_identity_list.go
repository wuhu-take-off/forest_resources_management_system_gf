package identity

import (
	"context"
	"forest_resources_management_system_gf/api/identity/identity_v1"
	"forest_resources_management_system_gf/internal/logic/l_identity"
)

func (c *ControllerIdentity_v1) IdentityList(ctx context.Context, req *identity_v1.IdentityListReq) (res *identity_v1.IdentityListRes, err error) {
	return l_identity.NewIdentityLogic().IdentityList(ctx, req)
}
