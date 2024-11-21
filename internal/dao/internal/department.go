// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DepartmentDao is the data access object for table department.
type DepartmentDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns DepartmentColumns // columns contains all the column names of Table for convenient usage.
}

// DepartmentColumns defines and stores column names for table department.
type DepartmentColumns struct {
	DepartmentId       string // 部门ID
	DepartmentName     string // 部门名称
	DepartmentParentId string // 父类部门
}

// departmentColumns holds the columns for table department.
var departmentColumns = DepartmentColumns{
	DepartmentId:       "department_id",
	DepartmentName:     "department_name",
	DepartmentParentId: "department_parent_id",
}

// NewDepartmentDao creates and returns a new DAO object for table data access.
func NewDepartmentDao() *DepartmentDao {
	return &DepartmentDao{
		group:   "default",
		table:   "department",
		columns: departmentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DepartmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DepartmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DepartmentDao) Columns() DepartmentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DepartmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DepartmentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DepartmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
