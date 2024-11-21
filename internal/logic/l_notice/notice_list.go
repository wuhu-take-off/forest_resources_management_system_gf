package l_notice

import (
	"context"
	"forest_resources_management_system_gf/api/notice/notice_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultNotice) NoticeList(ctx context.Context, req *notice_v1.NoticeListReq) (res *notice_v1.NoticeListRes, err error) {
	var notices []*entity.Notice
	if err := dao.Notice.Ctx(ctx).Page(req.Page, req.Limit).Order(dao.Notice.Columns().NoticeId).Scan(&notices); err != nil {
		g.Log().Errorf(ctx, "NoticeList Error: %v", err)
		return nil, common.NewError(consts.InternalServerError, "NoticeList Error")
	}
	res = &notice_v1.NoticeListRes{
		NoticeList: make([]*notice_v1.NoticeListStruct, 0),
	}
	common.CopyFields(notices, &res.NoticeList)

	if total, err := dao.Notice.Ctx(ctx).Count(); err != nil {
		g.Log().Errorf(ctx, "NoticeList  Error: %v", err)
		return nil, common.NewError(consts.InternalServerError, "NoticeList Error")
	} else {
		res.TotalCount = total
	}

	//policyInfo.PolicyTime = common.DateConversion(req.PolicyTime).(int64)
	return
}
