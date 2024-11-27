// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// IdentityDao is the data access object for table identity.
type IdentityDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns IdentityColumns // columns contains all the column names of Table for convenient usage.
}

// IdentityColumns defines and stores column names for table identity.
type IdentityColumns struct {
	IdentityId      string // 身份ID
	IdentityName    string // 身份名称
	ModuleAuthority string // 模块访问权限
}

// identityColumns holds the columns for table identity.
var identityColumns = IdentityColumns{
	IdentityId:      "identity_id",
	IdentityName:    "identity_name",
	ModuleAuthority: "module_authority",
}

// NewIdentityDao creates and returns a new DAO object for table data access.
func NewIdentityDao() *IdentityDao {
	return &IdentityDao{
		group:   "default",
		table:   "identity",
		columns: identityColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *IdentityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *IdentityDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *IdentityDao) Columns() IdentityColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *IdentityDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *IdentityDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *IdentityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
