// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NoticeDao is the data access object for table notice.
type NoticeDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns NoticeColumns // columns contains all the column names of Table for convenient usage.
}

// NoticeColumns defines and stores column names for table notice.
type NoticeColumns struct {
	NoticeId        string // 公告id
	NoticePublisher string // 发布人
	NoticeTime      string // 发布日期
	NoticePriority  string // 优先级
	NoticeContent   string // 公告内容
	NoticeAnnex     string // 附件信息
}

// noticeColumns holds the columns for table notice.
var noticeColumns = NoticeColumns{
	NoticeId:        "notice_id",
	NoticePublisher: "notice_publisher",
	NoticeTime:      "notice_time",
	NoticePriority:  "notice_priority",
	NoticeContent:   "notice_content",
	NoticeAnnex:     "notice_annex",
}

// NewNoticeDao creates and returns a new DAO object for table data access.
func NewNoticeDao() *NoticeDao {
	return &NoticeDao{
		group:   "default",
		table:   "notice",
		columns: noticeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NoticeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NoticeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NoticeDao) Columns() NoticeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NoticeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NoticeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NoticeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
