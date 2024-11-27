// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DepartmentAfficheDao is the data access object for table department_affiche.
type DepartmentAfficheDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns DepartmentAfficheColumns // columns contains all the column names of Table for convenient usage.
}

// DepartmentAfficheColumns defines and stores column names for table department_affiche.
type DepartmentAfficheColumns struct {
	DepartmentAfficheId         string // 部门公告主键
	DepartmentAfficheTitle      string // 部门公告标题
	DepartmentAfficheContent    string // 部门公告内容
	DepartmentId                string // 发布部门
	UserId                      string // 创建用户
	DepartmentAfficheCreateTime string // 创建时间
	DepartmentAfficheUpdateTime string // 修改时间
}

// departmentAfficheColumns holds the columns for table department_affiche.
var departmentAfficheColumns = DepartmentAfficheColumns{
	DepartmentAfficheId:         "department_affiche_id",
	DepartmentAfficheTitle:      "department_affiche_title",
	DepartmentAfficheContent:    "department_affiche_content",
	DepartmentId:                "department_id",
	UserId:                      "user_id",
	DepartmentAfficheCreateTime: "department_affiche_create_time",
	DepartmentAfficheUpdateTime: "department_affiche_update_time",
}

// NewDepartmentAfficheDao creates and returns a new DAO object for table data access.
func NewDepartmentAfficheDao() *DepartmentAfficheDao {
	return &DepartmentAfficheDao{
		group:   "default",
		table:   "department_affiche",
		columns: departmentAfficheColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DepartmentAfficheDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DepartmentAfficheDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DepartmentAfficheDao) Columns() DepartmentAfficheColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DepartmentAfficheDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DepartmentAfficheDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DepartmentAfficheDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
