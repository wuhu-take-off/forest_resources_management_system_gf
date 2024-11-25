// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// DepartmentAffiche is the golang structure for table department_affiche.
type DepartmentAffiche struct {
	DepartmentAfficheId         int    `json:"departmentAfficheId"         orm:"department_affiche_id"          ` // 部门公告主键
	DepartmentAfficheTitle      string `json:"departmentAfficheTitle"      orm:"department_affiche_title"       ` // 部门公告标题
	DepartmentAfficheContent    string `json:"departmentAfficheContent"    orm:"department_affiche_content"     ` // 部门公告内容
	DepartmentId                int    `json:"departmentId"                orm:"department_id"                  ` // 发布部门
	UserId                      int    `json:"userId"                      orm:"user_id"                        ` // 创建用户
	DepartmentAfficheCreateTime int64  `json:"departmentAfficheCreateTime" orm:"department_affiche_create_time" ` // 创建时间
}
