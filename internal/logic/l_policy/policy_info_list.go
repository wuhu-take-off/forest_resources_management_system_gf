package l_policy

import (
	"context"
	"forest_resources_management_system_gf/api/policy/policy_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func (d defaultPolicy) PolicyInfoList(ctx context.Context, req *policy_v1.PolicyInfoListReq) (res *policy_v1.PolicyInfoListRes, err error) {
	var policies []*entity.Policy
	db := dao.Policy.Ctx(ctx)
	if req.Type != nil && *req.Type != 0 {
		db = db.Where(dao.Policy.Columns().PolicyType, *req.Type)
	}
	//限定时间范围
	if req.ExpirationDate != nil && *req.ExpirationDate != 0 {
		db = db.WhereLTE(dao.Policy.Columns().PolicyTime, min(*req.ExpirationDate, time.Now().Unix()))
	} else {
		db = db.WhereLTE(dao.Policy.Columns().PolicyTime, time.Now().Unix())
	}
	if err := db.Page(req.Page, req.Limit).Order(dao.Policy.Columns().PolicyId).Scan(&policies); err != nil {
		g.Log().Errorf(ctx, "Policy info list Error: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Policy info list Error")
	}
	res = &policy_v1.PolicyInfoListRes{
		PolityList: make([]*policy_v1.PolicyInfoListStruct, 0),
	}
	common.CopyFields(policies, &res.PolityList)

	if total, err := db.Count(); err != nil {
		g.Log().Errorf(ctx, "Policy info list Error: %v", err)
		return nil, common.NewError(consts.InternalServerError, " Policy info list Error")
	} else {
		res.TotalCount = total
	}

	//policyInfo.PolicyTime = common.DateConversion(req.PolicyTime).(int64)
	return

}
