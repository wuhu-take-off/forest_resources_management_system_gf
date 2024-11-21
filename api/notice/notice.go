// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package notice

import (
	"context"

	"forest_resources_management_system_gf/api/notice/notice_v1"
)

type INoticeNotice_v1 interface {
	NoticeCreate(ctx context.Context, req *notice_v1.NoticeCreateReq) (res *notice_v1.NoticeCreateRes, err error)
	NoticeDel(ctx context.Context, req *notice_v1.NoticeDelReq) (res *notice_v1.NoticeDelRes, err error)
	NoticeInfoList(ctx context.Context, req *notice_v1.NoticeInfoListReq) (res *notice_v1.NoticeInfoListRes, err error)
	NoticeList(ctx context.Context, req *notice_v1.NoticeListReq) (res *notice_v1.NoticeListRes, err error)
	NoticeModify(ctx context.Context, req *notice_v1.NoticeModifyReq) (res *notice_v1.NoticeModifyRes, err error)
}
