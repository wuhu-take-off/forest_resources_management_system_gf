package l_identity

import (
	"context"
	"forest_resources_management_system_gf/api/identity/identity_v1"
)

type defaultIdentity struct{}

type IIdentityLogic interface {
	IdentityList(ctx context.Context, req *identity_v1.IdentityListReq) (res *identity_v1.IdentityListRes, err error)
}

func NewIdentityLogic() IIdentityLogic {
	return &defaultIdentity{}
}
