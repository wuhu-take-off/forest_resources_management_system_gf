package l_department

import (
	"context"
	"forest_resources_management_system_gf/api/department/department_v1"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

func (d defaultDepartment) DepartmentList(ctx context.Context, req *department_v1.DepartmentListReq) (res *department_v1.DepartmentListRes, err error) {
	var departments []*entity.Department
	if err := dao.Department.Ctx(ctx).OrderDesc(dao.Department.Columns().DepartmentId).Scan(&departments); err != nil {
		g.Log().Errorf(ctx, "scan department failed: %v", err)
		return nil, common.NewError(consts.InternalServerError, "scan department failed")
	}
	departmentMap := make(map[int][]*department_v1.DepartmentListStruct)
	res = &department_v1.DepartmentListRes{
		DepartmentList: make([]*department_v1.DepartmentListStruct, 0),
	}
	for _, department := range departments {
		if department.DepartmentParentId == 0 {
			res.DepartmentList = append(res.DepartmentList, &department_v1.DepartmentListStruct{
				DepartmentId:       department.DepartmentId,
				DepartmentName:     department.DepartmentName,
				DepartmentParentId: department.DepartmentParentId,
				Children:           departmentMap[department.DepartmentId],
			})
		} else {
			var tmp *department_v1.DepartmentListStruct
			common.CopyFields(department, &tmp)
			departmentMap[department.DepartmentParentId] = append(departmentMap[department.DepartmentParentId], tmp)
		}
	}
	return
}
