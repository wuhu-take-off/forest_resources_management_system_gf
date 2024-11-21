package policy_v1

import "github.com/gogf/gf/v2/frame/g"

type PolicyListStruct struct {
	PolicyId         int    `json:"policyId"         ` // 政策Id
	PolicyHeadline   string `json:"policyHeadline"   ` // 政策标题
	PolicyDepartment string `json:"policyDepartment" ` // 发布部门
	PolicyTime       int64  `json:"policyTime"       ` // 发布日期
	PolicyType       int    `json:"policyType"       ` // 政策类型
	PolicyContent    string `json:"policyContent"    ` // 政策内容
	PolicyAnnex      string `json:"policyAnnex"      ` // 附件
}
type PolicyListRes struct {
	PolityList []*PolicyListStruct `json:"polityList"` // 政策列表
	TotalCount int                 `json:"total_count"`
}
type PolicyListReq struct {
	g.Meta `path:"/policy/list" method:"post" tags:"policy"`
	Page   int `json:"page"`  // 页码
	Limit  int `json:"limit"` // 每页数量
}
