// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Identity is the golang structure for table identity.
type Identity struct {
	IdentityId      int    `json:"identityId"      orm:"identity_id"      ` // 身份ID
	IdentityName    string `json:"identityName"    orm:"identity_name"    ` // 身份名称
	ModuleAuthority string `json:"moduleAuthority" orm:"module_authority" ` // 模块访问权限
}
