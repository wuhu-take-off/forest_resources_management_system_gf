// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package resource

import (
	"context"

	"forest_resources_management_system_gf/api/resource/resource_v1"
)

type IResourceResource_v1 interface {
	ResourceVideo(ctx context.Context, req *resource_v1.ResourceVideoReq) (res *resource_v1.ResourceVideoRes, err error)
	ResourcePhoto(ctx context.Context, req *resource_v1.ResourcePhotoReq) (res *resource_v1.ResourcePhotoRes, err error)
}
