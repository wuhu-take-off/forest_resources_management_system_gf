// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Department is the golang structure for table department.
type Department struct {
	DepartmentId       int    `json:"departmentId"       orm:"department_id"        ` // 部门Id
	DepartmentName     string `json:"departmentName"     orm:"department_name"      ` // 部门名称
	DepartmentParentId int    `json:"departmentParentId" orm:"department_parent_id" ` // 父类部门
}
