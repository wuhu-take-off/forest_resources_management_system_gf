package l_notice

import (
	"context"
	"forest_resources_management_system_gf/api/notice/notice_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func (d defaultNotice) NoticeInfoList(ctx context.Context, req *notice_v1.NoticeInfoListReq) (res *notice_v1.NoticeInfoListRes, err error) {
	var notices []*entity.Notice
	db := dao.Notice.Ctx(ctx)
	if req.NoticePriority != nil && *req.NoticePriority != 0 {
		db = db.Where(dao.Notice.Columns().NoticePriority, *req.NoticePriority)
	}
	//限定时间范围
	if req.ExpirationDate != nil && *req.ExpirationDate != 0 {
		db = db.WhereLTE(dao.Notice.Columns().NoticeTime, min(*req.ExpirationDate, time.Now().Unix()))
	} else {
		db = db.WhereLTE(dao.Notice.Columns().NoticeTime, time.Now().Unix())
	}
	if err := db.Page(req.Page, req.Limit).Order(dao.Notice.Columns().NoticeId).Scan(&notices); err != nil {
		g.Log().Errorf(ctx, "Notice info list Error: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Notice info list Error")
	}
	res = &notice_v1.NoticeInfoListRes{
		NoticeInfoList: make([]*notice_v1.NoticeInfoListStruct, 0),
	}
	common.CopyFields(notices, &res.NoticeInfoList)

	if total, err := db.Count(); err != nil {
		g.Log().Errorf(ctx, "Notice info list Error: %v", err)
		return nil, common.NewError(consts.InternalServerError, " Notice info list Error")
	} else {
		res.TotalCount = total
	}

	//policyInfo.PolicyTime = common.DateConversion(req.PolicyTime).(int64)
	return
}
