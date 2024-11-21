package l_policy

import (
	"context"
	"forest_resources_management_system_gf/api/policy/policy_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultPolicy) PolicyList(ctx context.Context, req *policy_v1.PolicyListReq) (res *policy_v1.PolicyListRes, err error) {
	var policies []*entity.Policy
	if err := dao.Policy.Ctx(ctx).Page(req.Page, req.Limit).Order(dao.Policy.Columns().PolicyId).Scan(&policies); err != nil {
		g.Log().Errorf(ctx, "PolicyList Error: %v", err)
		return nil, common.NewError(consts.InternalServerError, "PolicyList Error")
	}
	res = &policy_v1.PolicyListRes{
		PolityList: make([]*policy_v1.PolicyListStruct, 0),
	}
	common.CopyFields(policies, &res.PolityList)

	if total, err := dao.Policy.Ctx(ctx).Count(); err != nil {
		g.Log().Errorf(ctx, "PolicyList Error: %v", err)
		return nil, common.NewError(consts.InternalServerError, "PolicyList Error")
	} else {
		res.TotalCount = total
	}

	//policyInfo.PolicyTime = common.DateConversion(req.PolicyTime).(int64)
	return
}
