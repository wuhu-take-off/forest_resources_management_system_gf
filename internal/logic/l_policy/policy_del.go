package l_policy

import (
	"context"
	"forest_resources_management_system_gf/api/policy/policy_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/do"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultPolicy) PolicyDel(ctx context.Context, req *policy_v1.PolicyDelReq) (res *policy_v1.PolicyDelRes, err error) {
	var policyInfo *do.Policy
	common.CopyFields(req, &policyInfo)
	if _, err = dao.Policy.Ctx(ctx).Delete(policyInfo); err != nil {
		g.Log().Errorf(ctx, "delete Policy info failed, err:%v", err)
		return nil, common.NewError(consts.InternalServerError, "delete Policy info failed")
	}
	return
}
