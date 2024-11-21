package l_notice

import (
	"context"
	"forest_resources_management_system_gf/api/notice/notice_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultNotice) NoticeDel(ctx context.Context, req *notice_v1.NoticeDelReq) (res *notice_v1.NoticeDelRes, err error) {
	var noticeInfo *do.Notice
	common.CopyFields(req, &noticeInfo)
	if _, err = dao.Notice.Ctx(ctx).Delete(noticeInfo); err != nil {
		g.Log().Errorf(ctx, "delete Notice info failed, err:%v", err)
		return nil, common.NewError(consts.InternalServerError, "delete Notice info failed")
	}
	return
}
