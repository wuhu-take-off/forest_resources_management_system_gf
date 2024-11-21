// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Policy is the golang structure for table policy.
type Policy struct {
	PolicyId         int    `json:"policyId"         orm:"policy_id"         ` // 政策Id
	PolicyHeadline   string `json:"policyHeadline"   orm:"policy_headline"   ` // 政策标题
	PolicyDepartment string `json:"policyDepartment" orm:"policy_department" ` // 发布部门
	PolicyTime       int64  `json:"policyTime"       orm:"policy_time"       ` // 发布日期
	PolicyType       int    `json:"policyType"       orm:"policy_type"       ` // 政策类型
	PolicyContent    string `json:"policyContent"    orm:"policy_content"    ` // 政策内容
	PolicyAnnex      string `json:"policyAnnex"      orm:"policy_annex"      ` // 附件
}
