package department_affiche

import (
	"context"
	"forest_resources_management_system_gf/internal/logic/l_department_affiche"

	"forest_resources_management_system_gf/api/department_affiche/department_affiche_v1"
)

func (c *ControllerDepartment_affiche_v1) DepartmentAfficheDel(ctx context.Context, req *department_affiche_v1.DepartmentAfficheDelReq) (res *department_affiche_v1.DepartmentAfficheDelRes, err error) {
	return l_department_affiche.NewDepartmentAfficheLogic().DepartmentAfficheDel(ctx, req)
}
