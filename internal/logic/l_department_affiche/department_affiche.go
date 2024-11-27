package l_department_affiche

import "forest_resources_management_system_gf/api/department_affiche"

type defaultDepartmentAdapter struct{}

type IDepartmentAfficheLogic interface {
	department_affiche.IDepartmentAfficheDepartment_affiche_v1
}

func NewDepartmentAfficheLogic() IDepartmentAfficheLogic {
	return &defaultDepartmentAdapter{}
}
