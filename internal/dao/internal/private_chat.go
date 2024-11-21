// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PrivateChatDao is the data access object for table private_chat.
type PrivateChatDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns PrivateChatColumns // columns contains all the column names of Table for convenient usage.
}

// PrivateChatColumns defines and stores column names for table private_chat.
type PrivateChatColumns struct {
	PrivateChatId         string // 私聊id
	PrivateChatSendId     string // 发送方id
	PrivateChatReceiverId string // 接收方id
	PrivateChatCreateTime string // 发送时间
	PrivateChatMessage    string //
}

// privateChatColumns holds the columns for table private_chat.
var privateChatColumns = PrivateChatColumns{
	PrivateChatId:         "private_chat_id",
	PrivateChatSendId:     "private_chat_send_id",
	PrivateChatReceiverId: "private_chat_receiver_id",
	PrivateChatCreateTime: "private_chat_create_time",
	PrivateChatMessage:    "private_chat_message",
}

// NewPrivateChatDao creates and returns a new DAO object for table data access.
func NewPrivateChatDao() *PrivateChatDao {
	return &PrivateChatDao{
		group:   "default",
		table:   "private_chat",
		columns: privateChatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PrivateChatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PrivateChatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PrivateChatDao) Columns() PrivateChatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PrivateChatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PrivateChatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PrivateChatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
