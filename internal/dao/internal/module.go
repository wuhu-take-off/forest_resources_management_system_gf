// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ModuleDao is the data access object for table module.
type ModuleDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns ModuleColumns // columns contains all the column names of Table for convenient usage.
}

// ModuleColumns defines and stores column names for table module.
type ModuleColumns struct {
	ModuleId     string // 模块Id
	ModuleName   string // 模块名称
	ModuleDesc   string // 模块描述
	ModuleGrade  string // 模块等级
	ModuleParent string // 父模块(0为顶级模块)
}

// moduleColumns holds the columns for table module.
var moduleColumns = ModuleColumns{
	ModuleId:     "module_id",
	ModuleName:   "module_name",
	ModuleDesc:   "module_desc",
	ModuleGrade:  "module_grade",
	ModuleParent: "module_parent",
}

// NewModuleDao creates and returns a new DAO object for table data access.
func NewModuleDao() *ModuleDao {
	return &ModuleDao{
		group:   "default",
		table:   "module",
		columns: moduleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ModuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ModuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ModuleDao) Columns() ModuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ModuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ModuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ModuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
