package l_forest_resource_info

import (
	"forest_resources_management_system_gf/api/forest_resource_info"
)

type defaultForestResourceInfo struct{}

type IForestResourceInfoLogic interface {
	forest_resource_info.IForestResourceInfoForest_resource_info_v1
}

func NewForestResourceInfoLogic() IForestResourceInfoLogic {
	return &defaultForestResourceInfo{}
}
