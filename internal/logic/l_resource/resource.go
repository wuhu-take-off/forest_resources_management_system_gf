package l_resource

import (
	"forest_resources_management_system_gf/api/resource"
)

type defaultResource struct{}

type IResourceLogic interface {
	resource.IResourceResource_v1
}

func NewResourceLogic() IResourceLogic {
	return &defaultResource{}
}
