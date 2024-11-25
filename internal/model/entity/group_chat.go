// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GroupChat is the golang structure for table group_chat.
type GroupChat struct {
	GroupChatId         int    `json:"groupChatId"         orm:"group_chat_id"          ` // 部门聊天主键
	GroupChatSendId     int    `json:"groupChatSendId"     orm:"group_chat_send_id"     ` // 发送方Id
	GroupChatReceiverId int    `json:"groupChatReceiverId" orm:"group_chat_receiver_id" ` // 接收方Id
	GroupChatCreateTime int64  `json:"groupChatCreateTime" orm:"group_chat_create_time" ` // 创建时间
	GroupChatMessage    string `json:"groupChatMessage"    orm:"group_chat_message"     ` // 消息内容
}
