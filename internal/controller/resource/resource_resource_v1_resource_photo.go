package resource

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_resource"

	"forest_resources_management_system_gf/api/resource/resource_v1"
)

func (c *ControllerResource_v1) ResourcePhoto(ctx context.Context, req *resource_v1.ResourcePhotoReq) (res *resource_v1.ResourcePhotoRes, err error) {
	return l_resource.NewResourceLogic().ResourcePhoto(ctx, req)
}
