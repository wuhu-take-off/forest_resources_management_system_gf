// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PrivateChat is the golang structure of table private_chat for DAO operations like Where/Data.
type PrivateChat struct {
	g.Meta                `orm:"table:private_chat, do:true"`
	PrivateChatId         interface{} // 私聊id
	PrivateChatSendId     interface{} // 发送方id
	PrivateChatReceiverId interface{} // 接收方id
	PrivateChatCreateTime interface{} // 发送时间
	PrivateChatMessage    interface{} //
}
