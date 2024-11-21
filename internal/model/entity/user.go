// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	UserId       int    `json:"userId"       orm:"user_id"       ` // 用户主键
	UserName     string `json:"userName"     orm:"user_name"     ` // 用户名
	UserPassword string `json:"userPassword" orm:"user_password" ` // 用户密码
	DepartmentId int    `json:"departmentId" orm:"department_id" ` // 部门ID
	UserPhone    string `json:"userPhone"    orm:"user_phone"    ` // 联系电话
	IdentityId   int    `json:"identityId"   orm:"identity_id"   ` // 身份ID
}
