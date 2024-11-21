// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package identity

import (
	"context"

	"forest_resources_management_system_gf/api/identity/identity_v1"
)

type IIdentityIdentity_v1 interface {
	IdentityList(ctx context.Context, req *identity_v1.IdentityListReq) (res *identity_v1.IdentityListRes, err error)
}
