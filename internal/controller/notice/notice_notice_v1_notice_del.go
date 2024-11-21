package notice

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_notice"

	"forest_resources_management_system_gf/api/notice/notice_v1"
)

func (c *ControllerNotice_v1) NoticeDel(ctx context.Context, req *notice_v1.NoticeDelReq) (res *notice_v1.NoticeDelRes, err error) {
	return l_notice.NewNoticeLogic().NoticeDel(ctx, req)

}
