// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TreeSpeciesDao is the data access object for table tree_species.
type TreeSpeciesDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns TreeSpeciesColumns // columns contains all the column names of Table for convenient usage.
}

// TreeSpeciesColumns defines and stores column names for table tree_species.
type TreeSpeciesColumns struct {
	TreeSpeciesId   string // 树种ID
	TreeSpeciesName string // 树种名称
}

// treeSpeciesColumns holds the columns for table tree_species.
var treeSpeciesColumns = TreeSpeciesColumns{
	TreeSpeciesId:   "tree_species_id",
	TreeSpeciesName: "tree_species_name",
}

// NewTreeSpeciesDao creates and returns a new DAO object for table data access.
func NewTreeSpeciesDao() *TreeSpeciesDao {
	return &TreeSpeciesDao{
		group:   "default",
		table:   "tree_species",
		columns: treeSpeciesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TreeSpeciesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TreeSpeciesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TreeSpeciesDao) Columns() TreeSpeciesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TreeSpeciesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TreeSpeciesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TreeSpeciesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
