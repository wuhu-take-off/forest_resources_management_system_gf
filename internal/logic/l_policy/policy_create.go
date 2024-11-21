package l_policy

import (
	"context"
	"encoding/json"
	"forest_resources_management_system_gf/api/policy/policy_v1"
	"forest_resources_management_system_gf/api/universal_class/universal_class_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
	"time"
)

func (d defaultPolicy) PolicyCreate(ctx context.Context, req *policy_v1.PolicyCreateReq) (res *policy_v1.PolicyCreateRes, err error) {
	t := strconv.FormatInt(time.Now().Unix(), 10) + "/"
	annex := make([]*universal_class_v1.Annex, 0)

	// 保存附件
	for i := range req.PolicyAnnex {
		if _, err := req.PolicyAnnex[i].Save(consts.PolicyAnnexDir + t); err != nil {
			g.Log().Errorf(ctx, "Failed to save policy annex: %v", err)
			return nil, common.NewError(consts.InternalServerError, "Failed to save policy annex")
		}
		annex = append(annex, &universal_class_v1.Annex{
			Url:  consts.PolicyAnnexDir + t,
			Name: req.PolicyAnnex[i].Filename,
		})
	}
	// 保存资源信息
	annexJson, _ := json.Marshal(annex)
	var policyInfo *entity.Policy
	common.CopyFields(req, &policyInfo)
	policyInfo.PolicyAnnex = string(annexJson)

	if _, err = dao.Policy.Ctx(ctx).Insert(policyInfo); err != nil {
		g.Log().Errorf(ctx, "Failed to create forest resource info: %v", err)
		return nil, common.NewError(consts.InternalServerError, "Failed to create forest resource info")
	}
	return
}
