// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GroupChatDao is the data access object for table group_chat.
type GroupChatDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns GroupChatColumns // columns contains all the column names of Table for convenient usage.
}

// GroupChatColumns defines and stores column names for table group_chat.
type GroupChatColumns struct {
	GroupChatId         string // 部门聊天主键
	GroupChatSendId     string // 发送方Id
	GroupChatReceiverId string // 接收方Id
	GroupChatCreateTime string // 创建时间
	GroupChatMessage    string // 消息内容
}

// groupChatColumns holds the columns for table group_chat.
var groupChatColumns = GroupChatColumns{
	GroupChatId:         "group_chat_id",
	GroupChatSendId:     "group_chat_send_id",
	GroupChatReceiverId: "group_chat_receiver_id",
	GroupChatCreateTime: "group_chat_create_time",
	GroupChatMessage:    "group_chat_message",
}

// NewGroupChatDao creates and returns a new DAO object for table data access.
func NewGroupChatDao() *GroupChatDao {
	return &GroupChatDao{
		group:   "default",
		table:   "group_chat",
		columns: groupChatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GroupChatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GroupChatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GroupChatDao) Columns() GroupChatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GroupChatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GroupChatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GroupChatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
