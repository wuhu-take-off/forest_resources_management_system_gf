package l_resource

import (
	"context"
	"forest_resources_management_system_gf/api/resource/resource_v1"
	"os"
)

func (d defaultResource) ResourceVideo(ctx context.Context, req *resource_v1.ResourceVideoReq) (res *resource_v1.ResourceVideoRes, err error) {
	url := req.Video.Url
	fileNames := req.Video.Name
	res = &resource_v1.ResourceVideoRes{}
	res.FileNames = fileNames
	res.Content, _ = os.ReadFile(url + fileNames)
	return res, nil
}
