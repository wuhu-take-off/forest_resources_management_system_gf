// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PeopleApplicationDao is the data access object for table people_application.
type PeopleApplicationDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns PeopleApplicationColumns // columns contains all the column names of Table for convenient usage.
}

// PeopleApplicationColumns defines and stores column names for table people_application.
type PeopleApplicationColumns struct {
	PeopleApplicationId      string // 民事申请ID
	PeopleApplicationTitle   string // 民事申请标题
	PeopleApplicant          string // 申请人
	PeopleApplicationPhone   string // 联系电话
	PeopleApplicationType    string // 申请类型
	PeopleApplicationContent string // 申请内容
	PeopleApplicationAnnex   string // 附件
	ModifyTime               string // 修改时间
	PeopleApplicationState   string // 申请状态(0:待审核 1:已处理)
	PeopleApplicationDesc    string // 处理描述
}

// peopleApplicationColumns holds the columns for table people_application.
var peopleApplicationColumns = PeopleApplicationColumns{
	PeopleApplicationId:      "people_application_id",
	PeopleApplicationTitle:   "people_application_title",
	PeopleApplicant:          "people_applicant",
	PeopleApplicationPhone:   "people_application_phone",
	PeopleApplicationType:    "people_application_type",
	PeopleApplicationContent: "people_application_content",
	PeopleApplicationAnnex:   "people_application_annex",
	ModifyTime:               "modify_time",
	PeopleApplicationState:   "people_application_state",
	PeopleApplicationDesc:    "people_application_desc",
}

// NewPeopleApplicationDao creates and returns a new DAO object for table data access.
func NewPeopleApplicationDao() *PeopleApplicationDao {
	return &PeopleApplicationDao{
		group:   "default",
		table:   "people_application",
		columns: peopleApplicationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PeopleApplicationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PeopleApplicationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PeopleApplicationDao) Columns() PeopleApplicationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PeopleApplicationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PeopleApplicationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PeopleApplicationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
