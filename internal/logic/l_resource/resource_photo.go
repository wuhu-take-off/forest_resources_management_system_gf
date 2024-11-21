package l_resource

import (
	"context"
	"forest_resources_management_system_gf/api/resource/resource_v1"
	"os"
)

func (d defaultResource) ResourcePhoto(ctx context.Context, req *resource_v1.ResourcePhotoReq) (res *resource_v1.ResourcePhotoRes, err error) {
	url := req.Photo.Url
	res = &resource_v1.ResourcePhotoRes{}
	fileNames := req.Photo.Name
	res.FileNames = fileNames
	res.Content, _ = os.ReadFile(url + fileNames)
	return res, nil
}
