// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PolicyDao is the data access object for table policy.
type PolicyDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns PolicyColumns // columns contains all the column names of Table for convenient usage.
}

// PolicyColumns defines and stores column names for table policy.
type PolicyColumns struct {
	PolicyId         string // 政策Id
	PolicyHeadline   string // 政策标题
	PolicyDepartment string // 发布部门
	PolicyTime       string // 发布日期
	PolicyType       string // 政策类型
	PolicyContent    string // 政策内容
	PolicyAnnex      string // 附件
}

// policyColumns holds the columns for table policy.
var policyColumns = PolicyColumns{
	PolicyId:         "policy_id",
	PolicyHeadline:   "policy_headline",
	PolicyDepartment: "policy_department",
	PolicyTime:       "policy_time",
	PolicyType:       "policy_type",
	PolicyContent:    "policy_content",
	PolicyAnnex:      "policy_annex",
}

// NewPolicyDao creates and returns a new DAO object for table data access.
func NewPolicyDao() *PolicyDao {
	return &PolicyDao{
		group:   "default",
		table:   "policy",
		columns: policyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PolicyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PolicyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PolicyDao) Columns() PolicyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PolicyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PolicyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PolicyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
