package department_affiche

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_department_affiche"

	"forest_resources_management_system_gf/api/department_affiche/department_affiche_v1"
)

func (c *ControllerDepartment_affiche_v1) DepartmentAfficheList(ctx context.Context, req *department_affiche_v1.DepartmentAfficheListReq) (res *department_affiche_v1.DepartmentAfficheListRes, err error) {
	return l_department_affiche.NewDepartmentAfficheLogic().DepartmentAfficheList(ctx, req)
}
