// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package department_affiche

import (
	"context"

	"forest_resources_management_system_gf/api/department_affiche/department_affiche_v1"
)

type IDepartmentAfficheDepartment_affiche_v1 interface {
	DepartmentAfficheCreate(ctx context.Context, req *department_affiche_v1.DepartmentAfficheCreateReq) (res *department_affiche_v1.DepartmentAfficheCreateRes, err error)
	DepartmentAfficheDel(ctx context.Context, req *department_affiche_v1.DepartmentAfficheDelReq) (res *department_affiche_v1.DepartmentAfficheDelRes, err error)
	DepartmentAfficheList(ctx context.Context, req *department_affiche_v1.DepartmentAfficheListReq) (res *department_affiche_v1.DepartmentAfficheListRes, err error)
	DepartmentAfficheModify(ctx context.Context, req *department_affiche_v1.DepartmentAfficheModifyReq) (res *department_affiche_v1.DepartmentAfficheModifyRes, err error)
}
