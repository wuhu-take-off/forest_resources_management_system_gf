// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GroupChat is the golang structure of table group_chat for DAO operations like Where/Data.
type GroupChat struct {
	g.Meta              `orm:"table:group_chat, do:true"`
	GroupChatId         interface{} // 部门聊天主键
	GroupChatSendId     interface{} // 发送方Id
	GroupChatReceiverId interface{} // 接收方Id
	GroupChatCreateTime interface{} // 创建时间
	GroupChatMessage    interface{} // 消息内容
}
