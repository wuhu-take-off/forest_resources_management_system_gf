package policy_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type PolicyCreateReq struct {
	g.Meta           `path:"/policy/create" method:"post" summary:"创建政策" tags:"policy"`
	PolicyHeadline   string              `json:"policyHeadline"   ` // 政策标题
	PolicyDepartment string              `json:"policyDepartment" ` // 发布部门
	PolicyTime       int64               `json:"policyTime"       ` // 发布日期
	PolicyType       int                 `json:"policyType"       ` // 政策类型
	PolicyContent    string              `json:"policyContent"    ` // 政策内容
	PolicyAnnex      []*ghttp.UploadFile `json:"policyAnnex"      ` // 附件
}
type PolicyCreateRes struct{}
