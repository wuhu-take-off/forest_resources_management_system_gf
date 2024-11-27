package l_department

import (
	"context"
	"forest_resources_management_system_gf/internal/common"
	"forest_resources_management_system_gf/internal/consts"
	"forest_resources_management_system_gf/internal/dao"
	"forest_resources_management_system_gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// DepartmentTool 部门工具类
type DepartmentTool struct {
	ctx context.Context
	err error
}

func NewDepartmentTool() *DepartmentTool {
	return new(DepartmentTool)
}

func (d *DepartmentTool) Err() error {
	return d.err
}

func (d *DepartmentTool) Ctx(ctx context.Context) *DepartmentTool {
	d.ctx = ctx
	return d
}
func (d *DepartmentTool) GetDepartmentEmployees(departmentParentIds int) []*entity.User {
	if d.err != nil {
		return []*entity.User{}
	}
	list1 := d.GetParentDepartmentList(departmentParentIds)
	list2 := d.GetSubDepartment(departmentParentIds)
	list := append(list1, list2[1:]...)
	var departmentIds []int
	for _, v := range list {
		departmentIds = append(departmentIds, v.DepartmentId)
	}
	var userInfoList []*entity.User
	if err := dao.User.Ctx(d.ctx).WhereIn(dao.User.Columns().DepartmentId, departmentIds).
		FieldsEx(dao.User.Columns().UserPassword).Scan(&userInfoList); err != nil {
		g.Log().Errorf(d.ctx, "Failed to get department employees: %v", err)
		d.err = common.NewError(consts.InsufficientFunds, "Failed to get department employees")
		return []*entity.User{}
	}
	return userInfoList
}

// GetDepartmentSub  获取部门以及其子部门
func (d *DepartmentTool) GetDepartmentSub(departmentParentIds ...int) []int {
	if len(departmentParentIds) == 0 || d.err != nil {
		return departmentParentIds
	}
	if d.ctx == nil {
		d.ctx = context.Background()
	}
	if array, err := dao.Department.Ctx(d.ctx).WhereIn(dao.Department.Columns().DepartmentParentId, departmentParentIds).
		Fields(dao.Department.Columns().DepartmentId).Array(); err != nil {
		g.Log().Errorf(d.ctx, "Failed to get department employees: %v", err)
		d.err = common.NewError(consts.InsufficientFunds, "Failed to get department employees")
		return departmentParentIds
	} else {
		var departmentIds []int
		for _, v := range array {
			departmentIds = append(departmentIds, v.Int())
		}
		departmentIds = append(departmentIds, d.GetDepartmentSub(departmentIds...)...)
		//递归获取所有子部门的员工信息
		departmentIds = append(departmentIds, departmentParentIds...)
		return departmentIds
	}
}

// GetDepartment 获取用户所在部门信息
func (d *DepartmentTool) GetDepartment(userId int) []*entity.Department {
	if d.err != nil {
		return []*entity.Department{}
	}
	if userId == 1 {
		//如果用户是超级管理员，则返回所有部门信息
		var departmentList []*entity.Department
		if err := dao.Department.Ctx(d.ctx).Scan(&departmentList); err != nil {
			g.Log().Errorf(d.ctx, "Failed to get department employees: %v", err)
			d.err = common.NewError(consts.InsufficientFunds, "Failed to get department employees")
			return departmentList
		}
		return departmentList
	}

	var userInfo *entity.User
	if err := dao.User.Ctx(d.ctx).Where(dao.User.Columns().UserId, userId).Fields(dao.User.Columns().DepartmentId).Scan(&userInfo); err != nil {
		g.Log().Errorf(d.ctx, "Failed to get department: %v", err)
		d.err = common.NewError(consts.InsufficientFunds, "Failed to get department")
		return []*entity.Department{}
	}
	list1 := d.GetParentDepartmentList(userInfo.DepartmentId)
	list2 := d.GetSubDepartment(userInfo.DepartmentId)
	return append(list1, list2[1:]...)
	//return append(d.GetParentDepartmentList(userInfo.DepartmentId), d.GetSubDepartment(userInfo.DepartmentId)...)
}

// GetParentDepartmentList 获取父部门列表
func (d *DepartmentTool) GetParentDepartmentList(departmentId int) (departmentList []*entity.Department) {
	if d.err != nil {
		return []*entity.Department{}
	}
	if err := dao.Department.Ctx(d.ctx).WhereIn(dao.Department.Columns().DepartmentId, departmentId).Scan(&departmentList); err != nil {
		g.Log().Errorf(d.ctx, "Failed to get department employees: %v", err)
		d.err = common.NewError(consts.InsufficientFunds, "Failed to get department employees")
		return []*entity.Department{}
	} else {
		if len(departmentList) > 0 {
			departmentList = append(d.GetParentDepartmentList(departmentList[0].DepartmentParentId), departmentList...)
		}
		return departmentList
	}
}

// GetSubDepartment 获取子类部门
func (d *DepartmentTool) GetSubDepartment(departmentId int) (departmentList []*entity.Department) {
	if d.err != nil {
		return []*entity.Department{}
	}
	if err := dao.Department.Ctx(d.ctx).WhereIn(dao.Department.Columns().DepartmentId, departmentId).Scan(&departmentList); err != nil {
		g.Log().Errorf(d.ctx, "Failed to get department employees: %v", err)
		d.err = common.NewError(consts.InsufficientFunds, "Failed to get department employees")
		return []*entity.Department{}
	}

	for i := range departmentList {
		var departmentTmpList []*entity.Department
		if err := dao.Department.Ctx(d.ctx).WhereIn(dao.Department.Columns().DepartmentParentId, departmentList[i].DepartmentId).Scan(&departmentTmpList); err != nil {
			g.Log().Errorf(d.ctx, "Failed to get department employees: %v", err)
			d.err = common.NewError(consts.InsufficientFunds, "Failed to get department employees")
			return []*entity.Department{}
		}
		departmentList = append(departmentList, departmentTmpList...)
	}
	return departmentList
}
